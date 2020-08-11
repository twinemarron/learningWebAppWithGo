package main

import (
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	// socket はこのクライアントのための WebSocket です
	socket *websocket.Conn
	// send はメッセージが送られるチャネルです
	send chan *message
	// room はこのクライアントが参加しているチャットルームです。
	room *room
	// userData はユーザーに関する情報を保持します
	userData map[string]interface{}
}

func (c *client) read() {
	for {
		var msg *message
		if err := c.socket.ReadJSON(&msg); err == nil {
			msg.When = time.Now()
			msg.Name = c.userData["name"].(string)
			if avatarURL, ok := c.userData["avatar_url"]; ok {
				msg.AvatarURL = avatarURL.(string)
			}
			c.room.forward <- msg
		} else {
			break
		}
	}
	c.socket.Close()
}

func (c *client) write() {
	for msg := range c.send {
		if err := c.socket.WriteJSON(msg); err != nil {
			break
		}
	}
	c.socket.Close()
}
