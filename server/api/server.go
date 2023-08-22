package api

import (
	"github.com/gin-gonic/gin"
	"github.com/quocgiahcmut/vue-go-chat/utils"
)

type Server struct {
	config utils.Config
	router *gin.Engine
}

func NewServer(config utils.Config) *Server {
	server := &Server{
		config: config,
	}

	router := gin.Default()

	socket := NewSocket()

	router.GET("/ping", server.Pong)
	router.GET("/socket.io/*any", gin.WrapH(socket.server))
	router.POST("/socket.io/*any", gin.WrapH(socket.server))

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) Pong(ctx *gin.Context) {
	ctx.JSON(200, "pong")
}
