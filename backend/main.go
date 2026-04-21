package main

import "fmt"

func main() {
	fmt.Println("🚀 Backend starting...")
	
	initClickHouse()

	go startConsumer()

	fmt.Println("📡 Consumer started")

	startAPI()
}
