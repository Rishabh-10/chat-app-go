package main

import (
	"chat-app/server"
	"net/http"
)

func main() {
	app := server.New()

	http.HandleFunc("/ws", app.Echo)

	http.ListenAndServe(":8001", nil)
}
