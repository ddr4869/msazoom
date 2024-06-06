package socket

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/ddr4869/msazoom/internal/repository"
	"github.com/gorilla/websocket"
)

type MessageMap struct {
	Mutex sync.RWMutex
	Map   map[string]MessageChat
}

type MessageChat struct {
	Participant map[string]Participant
	Sender      string
	Receiver    string
	CreatedAt   time.Time
}

var AllMessageRooms MessageMap

func (m *MessageMap) Init() {
	m.Map = make(map[string]MessageChat)
}

func (m *MessageMap) GetRoom(chat_key string) MessageChat {
	m.Mutex.RLock()
	defer m.Mutex.RUnlock()
	// check if the room exists
	if _, ok := m.Map[chat_key]; !ok {
		return MessageChat{}
	}
	return m.Map[chat_key]
}

func (m *MessageMap) InsertIntoRoom(chat_key string, sender, receiver string, ws *websocket.Conn) {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()

	now := time.Now()

	p := Participant{
		Conn: ws,
	}
	// if the room does not exist, create a new room
	if _, ok := m.Map[chat_key]; !ok {
		m.Map[chat_key] = MessageChat{
			Participant: map[string]Participant{sender: p},
			Sender:      sender,
			Receiver:    receiver,
			CreatedAt:   now,
		}
	} else {
		m.Map[chat_key].Participant[sender] = p
	}
}

// If the user is the host, delete the room
func (m *MessageMap) QuitRoom(chat_key string, username string) bool {
	m.Mutex.Lock()
	defer m.Mutex.Unlock()
	p := m.Map[chat_key].Participant[username]
	if p.Host {
		delete(m.Map, chat_key)
		return true
	} else {
		delete(m.Map[chat_key].Participant, username)
		return false
	}
}

func (m *MessageMap) Broadcast(c context.Context, r repository.Repository) {
	for {
		socketData := <-MessageSocketChannel
		// type assertion
		sender, ok := socketData.Data["sender"].(string)
		if !ok {
			log.Println("Invalid type for sender")
			continue
		}

		receiver, ok := socketData.Data["receiver"].(string)
		if !ok {
			log.Println("Invalid type for receiver")
			continue
		}

		message, ok := socketData.Data["message"].(string)
		if !ok {
			log.Println("Invalid type for message")
			continue
		}
		fmt.Println("socketData.Client set, sender, receiver, message", sender, receiver, message)
		_, err := r.WriteFriendMessage(c, sender, receiver, message)
		if err != nil {
			log.Println("Failed to write friend message, err ->  ", err)
			continue
		}

		for _, client := range m.Map[socketData.ID].Participant {
			if client.Conn != socketData.Client {
				AllChatRooms.Mutex.Lock()
				err := client.Conn.WriteJSON(socketData.Data)
				if err != nil {
					client.Conn.Close()
					AllChatRooms.Mutex.Unlock()
					return
				}
				AllChatRooms.Mutex.Unlock()
			}
		}
	}
}
