package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Connection *websocket.Conn
}

func New() *Server {
	return &Server{}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		fmt.Println(origin)
		return origin == "ws://localhost"
	},
}

func (s *Server) Echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		fmt.Printf("Received message: %s\n", message)

		err = conn.WriteMessage(websocket.TextMessage, []byte("Message received"))
		if err != nil {
			log.Println("Error writing message:", err)
			break
		}
	}
}
