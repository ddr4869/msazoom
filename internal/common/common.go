package common

import (
	"fmt"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// Room 인터페이스 정의
type Room interface {
	GetID() string
	GetParticipants() map[string]Participant
}

// ChatRoom 구조체 정의
type ChatRoom struct {
	ID          string
	Participant map[string]Participant
	CreatedAt   time.Time
}

// GetID 메서드
func (c ChatRoom) GetID() string {
	return c.ID
}

// GetParticipants 메서드
func (c ChatRoom) GetParticipants() map[string]Participant {
	return c.Participant
}

// RoomManager 구조체 정의
type RoomManager struct {
	Mutex sync.RWMutex
	Map   map[string]Room
}

// Init 메서드
func (r *RoomManager) Init() {
	r.Map = make(map[string]Room)
}

// GetRoom 메서드
func (r *RoomManager) GetRoom(id string) Room {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	if room, ok := r.Map[id]; ok {
		return room
	}
	return nil
}

// Type switch를 사용하여 MessageRoom, MeetingRomm 분기 처리
func (r *RoomManager) InsertIntoRoom(chat_id string, room Room) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	fmt.Println("InsertIntoRoom, room -> ", room)
	switch room := room.(type) {
	case *MessageRoom:
		if _, exists := r.Map[chat_id]; !exists {
			r.Map[chat_id] = room
		}
		room.Participant[room.Sender] = Participant{Conn: room.Participant[room.Sender].Conn, Host: true}
		room.Participant[room.Receiver] = Participant{Conn: room.Participant[room.Receiver].Conn, Host: false}
	case *MeetingRoom:
		if _, exists := r.Map[chat_id]; !exists {
			r.Map[chat_id] = room
			room.Participant[room.Admin] = Participant{Conn: room.Participant[room.Admin].Conn, Host: true}
		} else {
			room.Participant[room.Admin] = Participant{Conn: room.Participant[room.Admin].Conn, Host: false}
		}

	}
}

func (r *RoomManager) QuitRoom(chat_id string, username string) bool {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	if room, ok := r.Map[chat_id]; ok {
		delete(room.GetParticipants(), username)
		if len(room.GetParticipants()) == 0 {
			delete(r.Map, chat_id)
			return true
		}
	}
	return false
}

func (r *RoomManager) Broadcast(socketChannel chan Message) {
	for {
		socketData := <-socketChannel
		r.Mutex.RLock()
		room, exists := r.Map[socketData.ID]
		r.Mutex.RUnlock()
		if !exists {
			continue
		}
		for _, client := range room.GetParticipants() {
			if client.Conn != socketData.Client {
				err := client.Conn.WriteJSON(socketData.Data)
				if err != nil {
					client.Conn.Close()
				}
			}
		}
	}
}

// Message 구조체 정의
type Message struct {
	ID     string
	Client *websocket.Conn
	Data   map[string]interface{}
}
