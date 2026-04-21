package main

import (
	"context"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
)

var conn, _ = clickhouse.Open(&clickhouse.Options{
	Addr: []string{"localhost:9000"},
})

func insertClickhouse(log Log) {
	conn.Exec(context.Background(),
		"INSERT INTO logs VALUES (?, ?, ?, ?)",
		time.Unix(log.Timestamp, 0),
		log.Service,
		log.Level,
		log.Latency,
	)
}
