package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {
	topic := "test"
	groupID := "example-group"

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:   []string{"localhost:9092"},
		GroupID:   groupID,
		Topic:     topic,
		Partition: 0,
		MinBytes:  10e3, // 10KB
		MaxBytes:  10e6, // 10MB
	})
	defer reader.Close()

	fmt.Println("Consumer started...")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Printf("Consumed: %s\n", string(msg.Value))
	}
}
