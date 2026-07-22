package tracking_ws

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

// Hub mantiene clientes WS y retransmite ubicaciones del conductor a oyentes.
type Hub struct {
	mu         sync.RWMutex
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte, 256),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			h.clients[client] = true
			h.mu.Unlock()
			log.Printf("tracking_ws: cliente conectado user_id=%d role_id=%d publisher=%v",
				client.UserID, client.RoleID, client.IsPublisher)

		case client := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
			wasPublisher := client.IsPublisher
			userID := client.UserID
			h.mu.Unlock()

			if wasPublisher {
				h.broadcastJSON(map[string]interface{}{
					"type":         "conductor_disconnected",
					"conductor_id": userID,
					"timestamp":    time.Now().UnixMilli(),
				})
			}
			log.Printf("tracking_ws: cliente desconectado user_id=%d", userID)

		case message := <-h.broadcast:
			h.mu.RLock()
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					// Cliente lento: se descarta el mensaje para no bloquear el hub.
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *Hub) broadcastJSON(payload map[string]interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("tracking_ws: error serializando broadcast: %v", err)
		return
	}
	select {
	case h.broadcast <- data:
	default:
		log.Printf("tracking_ws: buffer de broadcast lleno, mensaje descartado")
	}
}

func (h *Hub) broadcastExcept(sender *Client, payload map[string]interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("tracking_ws: error serializando mensaje: %v", err)
		return
	}

	h.mu.RLock()
	defer h.mu.RUnlock()
	for client := range h.clients {
		if client == sender {
			continue
		}
		select {
		case client.send <- data:
		default:
		}
	}
}
