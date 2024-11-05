package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var (
	websocketUpgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type Manager struct {
	clients ClientList
	sync.RWMutex

	// switch-case 로 분기 처리를 하는 경우 depth 가 깊고 길이가 길어지므로 key-value 식으로 key 에 event 를 매칭 시켜 사용하는 방식
	handlers map[string]EventHandler
}

func NewManager() *Manager {
	m := &Manager{
		clients:  make(ClientList),
		handlers: make(map[string]EventHandler),
	}

	m.setupEventHandlers()

	return m
}

func (m *Manager) setupEventHandlers() {
	// 'EventSendMessage' 이벤트가 들어오면, 맵핑되어 있는 'SendMessage' 실행하기 위해
	// 이벤트-실행 함수를 맵핑
	m.handlers[EventSendMessage] = SendMessage
}

func SendMessage(event Event, c *Client) error {
	fmt.Print(event)
	return nil
}

// 이벤트가 들어오는 경우.
// 예외처리 및 라우팅
func (m *Manager) routeEvent(event Event, c *Client) error {
	// 핸들러에 해당 이벤트 타입이 정의되어 있는지 체크
	if handler, ok := m.handlers[event.Type]; ok {
		// 해당 이벤트 타입이 정의되어 있는 경우 handler 에서 해당 이벤트를 실행
		if err := handler(event, c); err != nil {
			return err
		}
		return nil
	} else {
		return errors.New("there is no such event type")
	}
}

func (m *Manager) serveWS(w http.ResponseWriter, r *http.Request) {
	log.Println("New Connection")

	// upgrade regular http connection into websocket connection.
	// This function is called whenever websocket upgrade connection request is requested
	conn, err := websocketUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	client := NewClient(conn, m)
	m.addClient(client)

	go client.readMessages()
	go client.writeMessages()
}

func (m *Manager) addClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	// client connection status turn true
	m.clients[client] = true
}

func (m *Manager) removeClient(client *Client) {
	m.Lock()
	defer m.Unlock()

	if _, ok := m.clients[client]; ok {
		client.connection.Close()
		delete(m.clients, client)
		log.Println("Client disconnected")
	}
}
