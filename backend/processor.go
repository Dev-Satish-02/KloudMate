package main

import (
	"encoding/json"
	"log"
)

type Log struct {
	Service   string  `json:"service"`
	Level     string  `json:"level"`
	Latency   float64 `json:"latency"`
	Timestamp int64   `json:"timestamp"`
}

func processLog(data []byte) {
	var logData Log
	json.Unmarshal(data, &logData)

	insertClickhouse(logData)
	updateRedis(logData)
	checkAlert(logData)

	log.Println("Processed:", logData)
}
