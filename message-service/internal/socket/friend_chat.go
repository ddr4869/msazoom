package socket

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/ddr4869/msazoom/message-service/internal/repository"
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
		_, err := r.WriteFriendMessage(c, sender, receiver, message)
		if err != nil {
			log.Println("Failed to write friend message, err ->  ", err)
			continue
		}

		for _, client := range m.Map[socketData.ID].Participant {
			// kafka test
			// kafkaMessage := &sarama.ProducerMessage{
			// 	Topic:     GenerateSocketKey(sender, receiver),
			// 	Key:       sarama.StringEncoder("TEST"),
			// 	Value:     sarama.ByteEncoder(message),
			// 	Partition: int32(0),
			// }
			// _, _, err = Producer.SendMessage(kafkaMessage)
			// if err != nil {
			// 	log.Printf("Failed to send message to Kafka: %v", err)
			// }

			if client.Conn != socketData.Client {
				AllMessageRooms.Mutex.Lock()
				err := client.Conn.WriteJSON(socketData.Data)
				if err != nil {
					client.Conn.Close()
					AllMessageRooms.Mutex.Unlock()
					return
				}
				AllMessageRooms.Mutex.Unlock()
			}
		}
	}
}
