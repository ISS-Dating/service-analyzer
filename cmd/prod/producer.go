package main

import (
	"github.com/segmentio/nsq-go"
)

func main() {
	// Starts a new producer that publishes to the TCP endpoint of a nsqd node.
	// The producer automatically handles connections in the background.
	producer, _ := nsq.StartProducer(nsq.ProducerConfig{
		Topic:   "hello",
		Address: "localhost:4150",
	})

	// Publishes a message to the topic that this producer is configured for,
	// the method returns when the operation completes, potentially returning an
	// error if something went wrong.
	producer.Publish([]byte("Hello World!"))

	// Stops the producer, all in-flight requests will be canceled and no more
	// messages can be published through this producer.
	producer.Stop()
}
