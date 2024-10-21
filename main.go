package main

import (
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, req *http.Request) {
	log.Println("Hello go http server.")
	msg := "Hello go http server."
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Fatalf("Failed to start go http server.")
	}
}

func main() {
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8080", nil)
}
