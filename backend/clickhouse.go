package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var conn clickhouse.Conn

func initClickHouse() {
	var err error

	conn, err = clickhouse.Open(&clickhouse.Options{
		Addr:     []string{"localhost:8123"},
		Protocol: clickhouse.HTTP,
	})

	if err != nil {
		panic(err)
	}

	if err := conn.Ping(context.Background()); err != nil {
		panic("ClickHouse connection failed: " + err.Error())
	}

	fmt.Println("✅ Connected to ClickHouse")
}

func insertClickhouse(log Log) {
	err := conn.Exec(context.Background(),
		"INSERT INTO logs VALUES (?, ?, ?, ?)",
		time.Unix(log.Timestamp, 0),
		log.Service,
		log.Level,
		log.Latency,
	)

	if err != nil {
		fmt.Println("❌ ClickHouse insert error:", err)
	} else {
		fmt.Println("✅ Inserted into ClickHouse")
	}
}