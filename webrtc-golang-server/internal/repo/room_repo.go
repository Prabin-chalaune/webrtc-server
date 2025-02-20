package repo

import (
	"sync"

	"github.com/prabin/webrtc-server/internal/models"
)

var (
	rooms = make(map[string]*models.Room)
	mutex sync.Mutex
)

func GetOrCreateRoom(roomID string) *models.Room {
	mutex.Lock()
	defer mutex.Unlock()

	if room, exists := rooms[roomID]; exists {
		return room
	}

	room := models.CreateRoom(roomID)
	rooms[roomID] = room
	return room
}

func GetRoom(roomID string) *models.Room {
	mutex.Lock()
	defer mutex.Unlock()

	return rooms[roomID]
}
