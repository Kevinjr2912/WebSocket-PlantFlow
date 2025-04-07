package adapters

import (
	"consumersData/domain/entities"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

type Gorilla struct {
	clients  map[*websocket.Conn]bool
	mu       sync.Mutex
	upgrader websocket.Upgrader
}

func NewGorilla() *Gorilla {
	return &Gorilla{
		clients: make(map[*websocket.Conn]bool),
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}
}

func (g *Gorilla) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := g.upgrader.Upgrade(w, r, nil)

	if err != nil {
		http.Error(w, "Could not upgrade to WebSocket", http.StatusBadRequest)
		return
	}

	g.mu.Lock()
	g.clients[conn] = true
	g.mu.Unlock()

	defer func() {
		conn.Close()
		
		// Eliminamos a un cliente desconectado
		g.mu.Lock()
		delete(g.clients, conn)
		g.mu.Unlock()
		
	}()

	for {

		_, _, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("Unexpected close error: %v\n", err)
			}
			break
		}

	}

}

func (g *Gorilla) BroadCast(messageType string, message []byte) {

	g.mu.Lock()
	defer g.mu.Unlock()

	for client := range g.clients {
		if err := client.WriteJSON(entities.Message{Type: messageType, Data: string(message)}); err != nil {
			client.Close()
			delete(g.clients, client)
		}
	}
	
}
