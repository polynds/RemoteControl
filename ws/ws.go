package ws

import (
	"bytes"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/polynds/RemoteControl/cache"
	"github.com/polynds/RemoteControl/device"
	"log"
	"net/http"
	"time"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newLine = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	hub  *Hub
	conn *websocket.Conn
	send chan []byte
}

func (c *Client) readDump() {
	defer func() {
		c.hub.unregister <- c
		_ = c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error {
		_ = c.conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newLine, space, -1))
		fmt.Println("received:", string(message))

		msg := NewMessage()
		data, err := msg.ParseMessage(string(message))
		if err != nil {
			log.Println(err)
			c.hub.broadcast <- []byte("error CMD")
			continue
		}

		screenSize := device.NewScreenSize()
		fmt.Println("received:", data)
		switch data.CMD {
		case "init":
			point := data.Data.(map[string]interface{})
			w := point["width"].(float64)
			h := point["height"].(float64)
			clientScreenSize := device.NewClientSize(int(w), int(h))
			cache.Set("clientScreenSize", clientScreenSize)
			fmt.Println("init client screen size:", clientScreenSize)
			fmt.Println("init server screen size:", screenSize)
		case "move":
			point := data.Data.(map[string]interface{})
			fmt.Println("received:", point["x"])
			x := point["x"].(float64)
			y := point["y"].(float64)
			fmt.Println("move:", x, y)
			clientScreenSize := cache.Get("clientScreenSize").(*device.Region)
			fmt.Println("move client screen size:", clientScreenSize)
			fmt.Println("move client screen size from cache:", cache.Get("clientScreenSize"))
			mapPoint := device.MapCoordinates(
				device.NewClientSize(clientScreenSize.Width, clientScreenSize.Height),
				screenSize,
				device.Coordinates{X: int(x), Y: int(y)},
			)
			fmt.Println("move mapPoint:", x, y)

			go device.TouchMove(mapPoint.X, mapPoint.Y)
		case "screen":
			go device.Capture()
		case "close":
			go device.CloseCapture()
		case "up":
			go device.PressKeyUp()
		case "down":
			go device.PressKeyDown()
		case "left":
			go device.PressKeyLeft()
		case "right":
			go device.PressKeyRight()
		}
		c.hub.broadcast <- []byte("ok")
	}
}

func (c *Client) writeDump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = w.Write(message)

			n := len(c.send)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newLine)
				_, _ = w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}

		case <-ticker.C:
			_ = c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}

		}
	}
}

func WsClient(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{
		hub:  hub,
		conn: conn,
		send: make(chan []byte, 256),
	}
	client.hub.register <- client

	go client.writeDump()
	go client.readDump()
}
