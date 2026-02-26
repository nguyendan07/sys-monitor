package hub

import (
	"encoding/json"
	"log"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/nguyendan07/sys-monitor/internal/models"
)

// The hub manages a collection of active clients.
type Hub struct {
	// The clients are connecting. Use the map as the set.
	clients	map[*websocket.Conn]bool

	Broadcast chan models.SystemMetrics
	Register chan *websocket.Conn
	Unregister chan *websocket.Conn

	mu sync.Mutex
}

func NewHub() *Hub {
	return &Hub{
		Broadcast: make(chan models.SystemMetrics),
		Register: make(chan *websocket.Conn),
		Unregister: make(chan *websocket.Conn),
		clients: make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <- h.Register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Println("New client connected")
		case client := <- h.Unregister:
			h.mu.Lock()
			_, ok := h.clients[client]
			if ok {
				delete(h.clients, client)
				client.Close()
			}
			h.mu.Unlock()
			log.Println("Client disconnected")
		case metrics := <- h.Broadcast:
			// Convert struct to  JSON
			msg, _ := json.Marshal(metrics)

			h.mu.Lock()
			for client := range h.clients {
				err := client.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					log.Printf("Error sending message: %v", err)
					client.Close()
					delete(h.clients, client)
				}
			}
			h.mu.Unlock()
		}
	}
}
