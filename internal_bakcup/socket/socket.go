package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var Upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type ChatSocketData struct {
	Data   map[string]interface{}
	ID     int
	Client *websocket.Conn
}

type MessageSocketData struct {
	Data   map[string]interface{}
	ID     string
	Client *websocket.Conn
}

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

var ChatSocketChannel = make(chan ChatSocketData)
var MessageSocketChannel = make(chan MessageSocketData)

type BroadCaster interface {
	Broadcast()
}

func SocketWrite(bc BroadCaster) {

}
