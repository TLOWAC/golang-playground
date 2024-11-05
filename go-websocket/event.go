package main

import "encoding/json"

type Event struct {
	Type    string          `json:"type"`
	Payload json.RawMessage `json:"payload"`
}

type EventHandler func(event Event, c *Client) error

// websocket event list
const (
	EventSendMessage = "send_message"
)

// when 'EventSendMessage' event is comming then use this dto
type SendMessageEvent struct {
	Message string `json:"message"`
	From    string `json:"From"`
}
