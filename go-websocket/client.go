package main

import (
	"log"

	"github.com/gorilla/websocket"
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	// egress is used to avoid concurrent writes on the websocket connection
	egress chan []byte
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan []byte),
	}
}

func (c *Client) readMessages() {
	defer func() {
		// cleanup connection
		c.manager.removeClient(c)
	}()

	for {
		messageType, payload, err := c.connection.ReadMessage()

		if err != nil {
			// when the web socket close upexpected
			// websocket.CloseGoingAway - 웹 소켓이 닫히고 있는 중인지
			// websocket.CloseAbnormalClosure - 웹 소켓이 비정상적으로 닫혔는지
			// 즉, errCode 가 CloseGoingAway(1001), CloseAbnormalClosure(1006) 인지 체크
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error reading message %v", err)
			}
			break
		}

		log.Println(messageType)
		log.Println(string(payload))

	}
}
