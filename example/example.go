package main

import (
	"fmt"
	"log"

	"github.com/timberslide/gotimberslide"
)

// Tweak these to be your own token and topic name
const (
	token = "TqPeDFG30Fz147NeCB59SUgT"
	topic = "kris/example"
)

func main() {
	// Create a Timberslide client
	client, err := ts.NewClient("gw.timberslide.com:443", token)

	// Connect to Timberslide
	err = client.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	defer client.Close()

	// Send some messages into a Timberslide topic
	ch, err := client.CreateChannel(topic)
	if err != nil {
		log.Fatalln(err)
	}
	ch.Send("Hello, Timberslide!")
	ch.Send("This is a very short collection of messages.")
	ch.Send("It's to quickly show some of the gotimberslide API.")
	ch.Send("Bye!")
	ch.Close()

	// Read those messages back
	for event := range client.Iter("kris/example", ts.PositionOldest) {
		fmt.Println("Received from Timberslide:", event.Message)
	}
}
