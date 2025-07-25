package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
	msgChan := make(chan string)

	//goroutine to simulate some background work
	go func() {
		time.Sleep(1 * time.Second)
		msgChan <- "Hello from the goroutine!"
	}()

	//wait for message from goroutine
	msg := <-msgChan
	fmt.Fprintf(w, msg)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}