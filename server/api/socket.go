package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type User struct {
	SocketID string `json:"socketId"`
	Username string `json:"username"`
}

var clients = make(map[*websocket.Conn]User)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origin
	},
}

func (server *Server) handleSocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Fatal("Upgrade error: ", err)
		return
	}
	defer conn.Close()

	var user User
	messageType, p, err := conn.ReadMessage()
	if err != nil {
		fmt.Println("Read error1: ", err)
		return
	}
	user.SocketID = conn.RemoteAddr().String()
	user.Username = string(p)
	clients[conn] = user
	conn.WriteMessage(messageType, []byte("Hello"))

	count := 1
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Read error2: ", err)
			delete(clients, conn)
			break
		}
		handleMessage(conn, messageType, p)

		conn.WriteJSON(count)
		count++
	}
}

func sendFriendsList(conn *websocket.Conn) {
	var users []User
	for _, user := range clients {
		users = append(users, user)
	}
	conn.WriteJSON(users)
}

func handleMessage(conn *websocket.Conn, messageType int, data []byte) {
	for clientConn, user := range clients {
		if conn != clientConn {
			clientConn.WriteMessage(messageType, append([]byte(user.Username+": "), data...))
		}
	}
}
