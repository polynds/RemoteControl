package ws

import (
	"encoding/json"
	"errors"
	"fmt"
)

type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}
type Message struct {
	CMD  string      `json:"cmd"`
	Data interface{} `json:"data"`
}

func NewMessage() *Message {
	return &Message{}
}

func (m *Message) ParseJson(data string) (*Message, error) {
	if !json.Valid([]byte(data)) {
		return nil, errors.New("invalid json")
	}
	var msg Message
	err := json.Unmarshal([]byte(data), &msg)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (m *Message) ParseMessage(data string) (*Message, error) {
	message, err := m.ParseJson(data)
	if message.CMD == "point" {
		point, ok := message.Data.(map[string]interface{})
		if ok {
			x := point["x"].(float64)
			y := point["y"].(float64)
			fmt.Println("x:", x)
			fmt.Println("y:", y)
		} else {
			fmt.Println("Data字段不是预期的类型")
		}
	}
	return message, err
}
