package main

import (
	"log"
	"net/http"

	"github.com/prabin/webrtc-server/api/routes"
	"github.com/prabin/webrtc-server/pkg/webrtc"
)

func main() {
	webrtc.Init()

	router := routes.SetupRoutes()

	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/", fs)

	log.Println("Server started on :8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
