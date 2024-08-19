package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Server struct {
	Connections map[*websocket.Conn]bool
	Broadcast   chan string
}

func New() *Server {
	conns := make(map[*websocket.Conn]bool)
	broadcaster := make(chan string)

	return &Server{Connections: conns, Broadcast: broadcaster}
}

// used for upgarding the protocol from HTTP to websocket connection
// also for customizing the connection eg buffer size, allowing origins, handling errors
// sub protocols accepted etc
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,

	// currently allowing all origins
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// creating a connection with the clients
func (s *Server) Start(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer conn.Close()

	s.Connections[conn] = true

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Error reading message:", err)
			break
		}

		s.Broadcast <- string(message)
	}
}

// used for writing the messages back to the clients
func (s *Server) DisplayMessages() {
	for {
		msg := <-s.Broadcast

		for conn := range s.Connections {
			err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				fmt.Println("Error: ", err.Error())
			}
		}
	}
}
