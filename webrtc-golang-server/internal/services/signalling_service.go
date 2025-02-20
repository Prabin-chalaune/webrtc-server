package services

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/prabin/webrtc-server/api/dtos"
	"github.com/prabin/webrtc-server/internal/repo"
)

func HandleSignalingMessage(conn *websocket.Conn, msg dtos.SignalingMessage) {
	switch msg.Type {
	case "join":
		room := repo.GetOrCreateRoom(msg.Room)
		room.AddClient(conn)
		log.Printf("Client joined room: %s", msg.Room)
	case "offer", "answer", "candidate":
		broadcastToRoom(msg.Room, msg)
	default:
		log.Println("Unknown message type:", msg.Type)
	}
}

func broadcastToRoom(roomID string, msg dtos.SignalingMessage) {
	room := repo.GetRoom(roomID)
	if room == nil {
		log.Println("Room not found:", roomID)
		return
	}

	for client := range room.Clients {
		if err := client.WriteJSON(msg); err != nil {
			log.Println("WebSocket write error:", err)
		}
	}
}
