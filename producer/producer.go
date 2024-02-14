package main

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

func main() {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "my-topic",
	})
	err := w.WriteMessages(context.Background(),
		kafka.Message{
			Value: []byte("Hello, Kafka!"),
		},
	)
	if err != nil {
		log.Fatalf("failed to write messages: %v", err)
	}
	if err := w.Close(); err != nil {
		log.Fatalf("failed to close writer: %v", err)
	}
	log.Println("Message written")
}
