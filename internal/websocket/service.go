package websocket

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
	ws "github.com/gorilla/websocket"
)

type Service struct {
	clients    map[string]*Client
	upgrader   *ws.Upgrader
	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

func NewService(upgrader *ws.Upgrader) *Service {
	return &Service{upgrader: upgrader, register: make(chan *Client, 256), unregister: make(chan *Client, 256), broadcast: make(chan *Message, 256)}
}

func (s *Service) Upgrade(w http.ResponseWriter, r *http.Request) (*ws.Conn, error) {
	return s.upgrader.Upgrade(w, r, nil)
}

func (s *Service) StartHub() {
	s.clients = make(map[string]*Client)
	for {
		select {
		case client := <-s.register:
			s.clients[client.conn.RemoteAddr().String()] = client
			fmt.Printf(color.GreenString("Client connected:%s\n"), client.conn.RemoteAddr())
		case client := <-s.unregister:
			delete(s.clients, client.conn.RemoteAddr().String())
			client.conn.Close()
			fmt.Printf(color.RedString("Client disconnected:%s\n"), client.conn.RemoteAddr())
		case msg := <-s.broadcast:
			for _, client := range s.clients {
				select {
				case client.send <- msg:
				default:
					close(client.send)
					delete(s.clients, client.conn.RemoteAddr().String())
				}
			}
		}
	}
}
