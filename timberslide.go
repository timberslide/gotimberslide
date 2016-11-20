package ts

import (
	"fmt"
	"io/ioutil"
	"net"
	"os/user"

	"github.com/BurntSushi/toml"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const (
	// PositionNewest will start reading a topic at the newest message
	PositionNewest int64 = -1
	// PositionOldest will start reading a topic from the oldest message
	PositionOldest int64 = -2
)

// Client contains configuration for talking to Timberslide
type Client struct {
	Host  string
	Token string
	conn  *grpc.ClientConn
}

// NewClient creates a new client
func NewClient(host string, token string) (Client, error) {
	var client Client
	client.Host = host
	client.Token = token
	return client, nil
}

// NewClientFromFile creates a client from the config file
func NewClientFromFile(file string) (Client, error) {
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

// addTokenToContext adds a bearer token to the given context
func addTokenToContext(ctx context.Context, token string) context.Context {
	md := metadata.Pairs("ts-access-token", token)
	ctx = metadata.NewContext(ctx, md)
	return ctx
}

// Connect connects to Timberslide
func (c *Client) Connect() error {
	host, _, err := net.SplitHostPort(c.Host)
	if err != nil {
		return err
	}
	creds := credentials.NewClientTLSFromCert(nil, host)
	c.conn, err = grpc.Dial(c.Host, grpc.WithTransportCredentials(creds), grpc.WithBlock())
	if err != nil {
		return err
	}
	return nil
}

// Close closes the client's connection to Timberslide
func (c *Client) Close() error {
	return c.conn.Close()
}

// Channel is used to send messages over
type Channel struct {
	Topic  string
	client IngestClient
	stream Ingest_StreamEventsClient
}

// CreateChannel returns a channel on which the client can send events
func (c *Client) CreateChannel(topic string) (Channel, error) {
	var err error
	channel := Channel{
		Topic: topic,
	}
	channel.client = NewIngestClient(c.conn)
	ctx := addTokenToContext(context.Background(), c.Token)
	channel.stream, err = channel.client.StreamEvents(ctx)
	if err != nil {
		return channel, err
	}
	return channel, nil
}

// Send sends a message for the channel's topic
func (c *Channel) Send(message string) error {
	if err := c.stream.Send(&Event{Topic: c.Topic, Message: message}); err != nil {
		return err
	}
	return nil
}

// Close closes out the channel with the server
func (c *Channel) Close() error {
	var err error
	if err = c.stream.Send(&Event{Topic: c.Topic, Done: true}); err != nil {
		return err
	}
	if _, err = c.stream.Recv(); err != nil {
		return err
	}
	err = c.stream.CloseSend()
	if err != nil {
		return err
	}
	return nil
}

// Iter providers an iterator to range over event in a topic
func (c *Client) Iter(topic string, position int64) <-chan Event {
	ch := make(chan Event)
	go func() {
		client := NewStreamerClient(c.conn)
		ctx := addTokenToContext(context.Background(), c.Token)
		stream, err := client.GetStream(ctx, &Topic{Name: topic, Position: position})
		if err != nil {
			return
		}
		for {
			event, err := stream.Recv()
			if err != nil {
				break
			}
			ch <- *event
		}
	}()
	return ch
}

// GetTopics returns a list of all topics visible to this user
func (c *Client) GetTopics() ([]*Topic, error) {
	client := NewTopicsClient(c.conn)
	ctx := addTokenToContext(context.Background(), c.Token)
	topicsReply, err := client.GetTopics(ctx, &TopicsReq{})
	return topicsReply.Topics, err
}
