package main

import (
	"github.com/gorilla/websocket"
)

// clientはチャットを行っている1人のユーザを示す
type client struct {
	socket *websocket.Conn
	// sendはメッセージが送られるチャンネル。バッファーつき。
	send chan []byte
	room *room
}

func (c *client) read() {
	for {
		// メッセージの読み込み
		if _, msg, err := c.socket.ReadMessage(); err == nil {
			c.room.forword <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteMessage(websocket.TextMessage, msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
