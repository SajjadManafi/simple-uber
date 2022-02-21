package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// upgradeConnection upgrades the HTTP connection to a WebSocket connection.
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebSocketConnection struct {
	*websocket.Conn
}

func (server *Server) createTrip(ctx *gin.Context) {
	server.WsEndPoint(ctx.Writer, ctx.Request)
}

func (server *Server) WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected To EndPoint!")

	err = ws.WriteMessage(websocket.TextMessage, []byte("Hello Client!"))
	if err != nil {
		log.Println(err)
		return
	}
}
