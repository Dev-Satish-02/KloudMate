package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func startConsumer() {
	fmt.Println("👀 Consumer loop starting...")

	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "logs",
		GroupID: "group-1",
	})

	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			fmt.Println("❌ Kafka error:", err)
			continue
		}

		fmt.Println("📥 Received message:", string(msg.Value))
		processLog(msg.Value)
	}
}
