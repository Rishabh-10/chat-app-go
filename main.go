package main

import (
	"chat-app/server"
	"net/http"
)

func main() {
	app := server.New()

	go app.DisplayMessages()

	http.HandleFunc("/ws", app.Start)

	http.ListenAndServe(":8001", nil)
}
