package websocket

import (
	"fmt"
	"os"
	"time"

	ws "github.com/gorilla/websocket"
)

type Client struct {
	conn *ws.Conn
	svc  *Service
	send chan *Message
}

func NewClient(conn *ws.Conn, svc *Service) *Client {
	return &Client{conn: conn, svc: svc, send: make(chan *Message, 256)}
}

func (c *Client) reader() {
	c.conn.SetReadLimit(int64(ReadBufferSize * ReadBufferSize))
	c.conn.SetReadDeadline(time.Now().Add(time.Minute * 3))
	defer func() {
		c.svc.unregister <- c
	}()
	for {
		messageType, p, err := c.conn.ReadMessage()
		if err != nil {
			fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
			return
		}
		c.svc.broadcast <- &Message{messageType: messageType, payload: p}
	}
}

func (c *Client) writer() {
	c.conn.SetWriteDeadline(time.Now().Add(time.Minute * 3))
	defer c.conn.Close()
	for {
		for msg := range c.send {
			err := c.conn.WriteMessage(ws.TextMessage, msg.payload)
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR:%v\n", err)
				return
			}

		}
	}
}
