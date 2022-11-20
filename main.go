package main

import (
	"log"
	"net/http"
	"shev-chat/handlers"
)

func init() {
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/ws", handlers.WebsocketHandler)
}

func main() {
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
