package main

import (
	"fmt"
	"time"
	types "msg-broker/types"
)

func main() {
    broker := types.NewBroker()

	subscriber := broker.Subscribe("example_topic")
	go func() {
		for {
			select {
			case msg, ok := <-subscriber.Channel:
				if !ok {
				fmt.Println("Subscriber channel closed.")
				return
				}
				fmt.Printf("Received: %v\n", msg)
			case <-subscriber.Unsubscribe:
				fmt.Println("Unsubscribed.")
				return
			}
		}
	}()

	broker.Publish("example_topic", "Hello, World!")
	broker.Publish("example_topic", "This is a test message.")

	time.Sleep(2 * time.Second)
	broker.Unsubscribe("example_topic", subscriber)

	broker.Publish("example_topic", "This message won't be received.")

	time.Sleep(time.Second)
}