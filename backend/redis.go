package main

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var rdb = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func updateRedis(log Log) {
	if log.Level == "ERROR" {
		key := fmt.Sprintf("error_count:%s", log.Service)
		rdb.Incr(context.Background(), key)
		rdb.Expire(context.Background(), key, 60*time.Second)
	}
}
