package models

import "github.com/gorilla/websocket"

type Room struct {
	ID      string
	Clients map[*websocket.Conn]bool
}

type Client struct {
	Conn *websocket.Conn
}

func CreateRoom(id string) *Room {
	return &Room{
		ID:      id,
		Clients: make(map[*websocket.Conn]bool),
	}
}

func (r *Room) AddClient(conn *websocket.Conn) {
	r.Clients[conn] = true
}

func (r *Room) RemoveClient(conn *websocket.Conn) {
	delete(r.Clients, conn)
}
