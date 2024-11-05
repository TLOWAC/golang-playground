package main

import (
	"context"
	"net/http"
)

func main() {

	ctx := context.Background()

	manager := NewManager(ctx)
	http.HandleFunc("/ws", manager.serveWS)
	http.HandleFunc("/login", manager.loginHandler)
}
