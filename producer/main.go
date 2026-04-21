package main

import (
	"context"
	"encoding/json"
	"math/rand"
	"time"

	"github.com/segmentio/kafka-go"
)

type Log struct {
	Service   string  `json:"service"`
	Level     string  `json:"level"`
	Latency   float64 `json:"latency"`
	Timestamp int64   `json:"timestamp"`
}

func main() {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "logs",
	})

	defer writer.Close() // good practice

	for {
		log := Log{
			Service:   "auth-service",
			Level:     []string{"INFO", "ERROR"}[rand.Intn(2)],
			Latency:   rand.Float64() * 500,
			Timestamp: time.Now().Unix(),
		}

		data, _ := json.Marshal(log)

		err := writer.WriteMessages(context.Background(), kafka.Message{
			Value: data,
		})

		if err != nil {
			panic(err) // or log it
		}

		time.Sleep(500 * time.Millisecond)
	}
}
