package websocket

import (
	"net/http"

	ws "github.com/gorilla/websocket"
)

type Websocket struct {
	Upgrader *ws.Upgrader
}

type Message struct {
	messageType int
	payload     []byte
}

var (
	ReadBufferSize   = 1024
	WriterBufferSize = 1024
)

func NewWebSocket() *Websocket {
	upgrader := &ws.Upgrader{
		ReadBufferSize:  ReadBufferSize,
		WriteBufferSize: WriterBufferSize,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	return &Websocket{Upgrader: upgrader}
}
