package api

import (
	"fmt"
	"log"
	"net/http"

	wsstore "github.com/SajjadManafi/simple-uber/internal/wsStore"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var clients = make(map[WebSocketConnection]string)

var wsChan = make(chan WsPayLoad)

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

// WsPayLoad represents a payload sent over the WebSocket connection.
type WsPayLoad struct {
	ID      int32               `json:"id"`
	Type    string              `json:"type"`
	Action  string              `json:"action"`
	Message string              `json:"message"`
	Conn    WebSocketConnection `json:"-"`
}

// WsJSONResponse represents a response sent over the WebSocket connection.
type WsJSONResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
}

// createTrip handles the create trip request.
func (server *Server) createTrip(ctx *gin.Context) {
	server.WsEndPoint(ctx.Writer, ctx.Request)
}

// WsEndPoint handles the WebSocket connection.
func (server *Server) WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("Client Connected To EndPoint!")

	conn := WebSocketConnection{ws}
	clients[conn] = ""

	var response WsJSONResponse
	response.Action = "connected"
	response.Message = `<em><small>Connected to server!</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
		return
	}

	go server.ListenToWs(&conn)
}

func (server *Server) ListenToWs(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayLoad

	for {
		err := conn.ReadJSON(&payload)
		if err == nil {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func (server *Server) ListenToWsChannel() {
	var response WsJSONResponse

	for {
		e := <-wsChan
		switch e.Action {
		case "setLocation":
			cordiante := wsstore.Cordinate{
				X: 44.04,
				Y: 25.0780,
			}

			server.driverWSS.Insert(wsstore.Driver{
				ID:         e.ID,
				Cordinate:  cordiante,
				Connection: e.Conn.Conn,
			})
			fmt.Println(server.driverWSS.Get(e.ID))
		}
		response.Action = "Got here"
		response.Message = fmt.Sprintf("Some message, and action was %s", e.Action)
		server.broadcast(&e.Conn, response)
	}
}

func (server *Server) broadcast(conn *WebSocketConnection, response WsJSONResponse) {
	err := conn.WriteJSON(response)
	if err != nil {
		log.Println(err)
		_ = conn.Close()
		delete(clients, *conn)
	}
}
