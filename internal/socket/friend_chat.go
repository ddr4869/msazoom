package socket

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type MessageMap struct {
	Mutex sync.RWMutex
	Map   map[int]MessageChat
}

type MessageChat struct {
	Sender    string
	Receiver  string
	Message   string
	CreatedAt time.Time
}

func (m *MessageMap) SocketRead(ws *websocket.Conn, req struct {
	ChatID   int
	Username string
}) {
	for {
		var message MessageChat
		err := ws.ReadJSON(&message)
		if err != nil {
			log.Println("Error reading message:", err)
			ws.Close()
			return
		}
		m.Mutex.Lock()
		m.Map[req.ChatID] = message
		m.Mutex.Unlock()

		socketData := SocketData{
			Data:   map[string]interface{}{"message": message.Message, "sender": message.Sender, "receiver": message.Receiver, "createdAt": message.CreatedAt},
			ID:     req.ChatID,
			Client: ws,
		}
		SocketChannel <- socketData
	}
}

func (m *MessageMap) Broadcast() {
	for {
		socketData := <-SocketChannel
		fmt.Println("broadcast msg -> ", socketData)
		m.Mutex.RLock()
		message, ok := m.Map[socketData.ID]
		m.Mutex.RUnlock()
		if ok {
			socketData.Data["message"] = message.Message
			socketData.Data["sender"] = message.Sender
			socketData.Data["receiver"] = message.Receiver
			socketData.Data["createdAt"] = message.CreatedAt
			m.Mutex.Lock()
			err := socketData.Client.WriteJSON(socketData.Data)
			if err != nil {
				fmt.Println("broadcast close ->", err)
				socketData.Client.Close()
				m.Mutex.Unlock()
				return
			}
			m.Mutex.Unlock()
		}
	}
}
