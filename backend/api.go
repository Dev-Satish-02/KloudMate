package main

import (
	"fmt"
	"net/http"
)

func startAPI() {
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "OK")
	})

	http.ListenAndServe(":8080", nil)
}
