package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var (
	pongWait     = 10 * time.Second
	pingInterval = (pongWait * 9) / 10
)

type ClientList map[*Client]bool

type Client struct {
	connection *websocket.Conn
	manager    *Manager

	// egress is used to avoid concurrent writes on the websocket connection
	egress chan []Event
}

func NewClient(conn *websocket.Conn, manager *Manager) *Client {
	return &Client{
		connection: conn,
		manager:    manager,
		egress:     make(chan []Event),
	}
}

func (c *Client) readMessages() {
	defer func() {
		// cleanup connection
		c.manager.removeClient(c)
	}()

	// ping-pong 데드라인 설정
	// writeMessage 에서 ping 을 보냈을때 pongWait 시간내 pong 이 readMessage 에 도착하지 않는 경우
	// websocket 연결을 종료한다.
	if err := c.connection.SetReadDeadline(time.Now().Add(pongWait)); err != nil {
		log.Println()
		return
	}

	// 메시지 크기 제한 설정
	c.connection.SetReadLimit(512)

	// ping-pong 이 정상적으로 이루어질 경우
	c.connection.SetPongHandler(c.pongHandler)

	for {
		_, payload, err := c.connection.ReadMessage()

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

		var request Event

		if err := json.Unmarshal(payload, &request); err != nil {
			log.Printf("error marshalling event : %v", err)
			break
		}

		if err := c.manager.routeEvent(request, c); err != nil {
			log.Printf("error handling message : %v", err)
		}
	}
}

func (c *Client) writeMessages() {
	defer func() {
		c.manager.removeClient(c)
	}()

	ticker := time.NewTicker(pingInterval)

	for {
		select {
		case message, ok := <-c.egress:
			// egress channel is closed
			if !ok {
				// send close message to client
				if err := c.connection.WriteMessage(websocket.CloseMessage, nil); err != nil {
					log.Println("connection closed: ", err)
				}
				return
			}

			data, err := json.Marshal(message)
			if err != nil {
				log.Println(err)
				return
			}

			// send text message to client
			if err := c.connection.WriteMessage(websocket.TextMessage, data); err != nil {
				// if error occured print error message
				log.Printf("failed to send message: %v", err)
			}

			log.Println("message sent")

		// heartbeat
		case <-ticker.C:
			log.Println("ping")

			// send a ping to the client
			if err := c.connection.WriteMessage(websocket.PingMessage, []byte(``)); err != nil {
				log.Println("write message error: ", err)
				return
			}
		}
	}
}

// ping-pong 이 정상적으로 연결 되었을 경우 ping-pong 타이머를 초기화
func (c *Client) pongHandler(pongMsg string) error {
	log.Println("pong")
	return c.connection.SetReadDeadline(time.Now().Add(pongWait))
}
