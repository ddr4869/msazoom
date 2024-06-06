package common

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

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

var ChatSocketChannel = make(chan Message)
var MessageSocketChannel = make(chan Message)

type BroadCaster interface {
	Broadcast()
}

func SocketWrite(bc BroadCaster) {

}
