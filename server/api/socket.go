package api

import (
	socketio "github.com/googollee/go-socket.io"
)

type Socket struct {
	server *socketio.Server
}

func NewSocket() *Socket {
	socket := &Socket{}
	server := socketio.NewServer(nil)
	socket.server = server
	return socket
}

func (socket *Socket) RegisterEvent() {
	socket.server.OnConnect("/", socket.HandleOnConnect)
}

func (socket *Socket) HandleOnConnect(conn socketio.Conn) error {
	conn.SetContext("")
	return nil
}
