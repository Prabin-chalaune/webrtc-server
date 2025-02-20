package routes

import (
	"github.com/gorilla/mux"
	"github.com/prabin/webrtc-server/api/controller"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ws", controller.HandleWebSocket)

	return router
}
