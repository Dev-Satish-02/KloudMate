package main

import (
	"context"
	"fmt"
)

func checkAlert(log Log) {
	key := fmt.Sprintf("error_count:%s", log.Service)
	count, _ := rdb.Get(context.Background(), key).Int()

	if count > 20 {
		fmt.Println("🚨 ALERT: High errors in", log.Service)
		triggerAction(log.Service)
	}
}

func triggerAction(service string) {
	fmt.Println("⚡ Restarting service:", service)
}
