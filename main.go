package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

// Message object to store details
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

// keep track of connected clients
var clients = make(map[*websocket.Conn]bool)

// broadcast channel using the Message structure
var broadcast = make(chan Message)

// server app calls the upgrader.upgrade from an http req handler to get a *Conn
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	defer ws.Close()

	// register our new client
	clients[ws] = true

	// loop run read and broadcast message
	for {
		var msg Message
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the recent received message to the broadcast channel
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		new_msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(new_msg)
			if err != nil {
				log.Printf("Error handling message: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	// create a simple file server from folder named public
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	// Configure websocket route
	http.HandleFunc("/ws", handleConnections)

	// start listenting on incoming chats
	go handleMessages()

	// start server on port 8000
	log.Println("http started on port 8000")
	err := http.ListenAndServe("0.0.0.0:8000", nil)
	if err != nil {
		log.Fatal("ListenAndServer error: ", err)
	}
}
