package main

import "net/http"

func main() {
	manager := NewManager()
	http.HandleFunc("/ws", manager.serveWS)

}
