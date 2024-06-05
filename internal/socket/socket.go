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

type SocketData struct {
	Data   map[string]interface{}
	ID     int
	Client *websocket.Conn
}

var SocketChannel = make(chan SocketData)

type BroadCaster interface {
	Broadcast()
}

func SocketWrite(bc BroadCaster) {

}
