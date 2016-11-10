package ts

import (
	"bufio"
	"context"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"os/user"

	"github.com/BurntSushi/toml"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

// Client contains configuration for talking to Timberslide
type Client struct {
	Host  string
	Token string
}

// NewClient creates a client from the config file
func NewClient(file string) (Client, error) {
	var client Client

	usr, err := user.Current()
	if err != nil {
		return client, err
	}

	rcLoc := fmt.Sprintf("%s/%s", usr.HomeDir, file)
	b, err := ioutil.ReadFile(rcLoc)
	if err != nil {
		return client, fmt.Errorf("Could not read configuration at %s", rcLoc)
	}
	if _, err = toml.Decode(string(b), &client); err != nil {
		return client, err
	}

	return client, nil
}

// Send takes stdin and sends it to the topic
func (c *Client) Send(topic string) error {
	fmt.Fprintf(os.Stderr, "ts: connecting...\n")
	host, _, err := net.SplitHostPort(c.Host)
	if err != nil {
		return err
	}
	creds := credentials.NewClientTLSFromCert(nil, host)
	conn, err := grpc.Dial(c.Host, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()

	client := NewIngestClient(conn)
	md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", c.Token))
	ctx := metadata.NewContext(context.Background(), md)
	stream, err := client.StreamEvents(ctx)
	if err != nil {
		return err
	}

	msg := &Event{Topic: topic}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg.Message = scanner.Text()
		// TODO make printing to stdout optional
		fmt.Println(msg.Message)
		stream.Send(msg)
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	// We finished normally, so...
	// set the done flag and block until we hear something from the server
	stream.Send(&Event{Topic: topic, Done: true})
	fmt.Fprintf(os.Stderr, "ts: done, waiting for server...\n")
	stream.Recv()
	stream.CloseSend()

	return nil
}

// Get receives the stream from the topic and writes it to stdout
func (c *Client) Get(topic string) error {
	host, _, err := net.SplitHostPort(c.Host)
	if err != nil {
		return err
	}
	creds := credentials.NewClientTLSFromCert(nil, host)
	conn, err := grpc.Dial(c.Host, grpc.WithTransportCredentials(creds))
	if err != nil {
		return err
	}
	client := NewStreamerClient(conn)
	md := metadata.Pairs("authorization", fmt.Sprintf("Bearer %s", c.Token))
	ctx := metadata.NewContext(context.Background(), md)
	stream, err := client.GetStream(ctx, &Topic{Topic: topic})
	if err != nil {
		return err
	}
	for {
		event, err := stream.Recv()
		if err != nil {
			return err
		}
		fmt.Println(event.Message)
	}
}
