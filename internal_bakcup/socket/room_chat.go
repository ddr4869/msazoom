package socket

import (
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[int]RoomChat
}

// define map[string]Participant type
type RoomChat struct {
	Participant map[string]Participant
	Title       string
	Private     bool
	Admin       string
	Created_at  time.Time
}

// #### Signaling ####
var AllChatRooms RoomMap

func (r *RoomMap) Init() {
	r.Map = make(map[int]RoomChat)
}

func (r *RoomMap) GetRoom(chat_id int) RoomChat {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	// check if the room exists
	if _, ok := r.Map[chat_id]; !ok {
		return RoomChat{}
	}
	return r.Map[chat_id]
}

func (r *RoomMap) GetRandomRoomKey() int {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	if len(r.Map) == 0 {
		return -1
	}
	keys := make([]int, 0, len(r.Map))
	for k, roomChat := range r.Map {
		if !roomChat.Private && len(roomChat.Participant) == 1 {
			keys = append(keys, k)
		}
	}
	return keys[rand.Intn(len(keys))]
}

// If the room does not exist, create a new room
// If the room already exists, insert the user into the room
func (r *RoomMap) InsertIntoRoom(chat_id int, chat_title, username string, private, host bool, ws *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	now := time.Now()

	p := Participant{
		Host: host,
		Conn: ws,
	}
	if host {
		r.Map[chat_id] = RoomChat{
			Participant: map[string]Participant{username: p},
			Title:       chat_title,
			Private:     private,
			Admin:       username,
			Created_at:  now,
		}
	} else {
		r.Map[chat_id].Participant[username] = p
	}
}

// If the user is the host, delete the room
func (r *RoomMap) QuitRoom(chat_id int, username string) bool {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	p := r.Map[chat_id].Participant[username]
	if p.Host {
		delete(r.Map, chat_id)
		return true
	} else {
		delete(r.Map[chat_id].Participant, username)
		return false
	}
}

func (r *RoomMap) Broadcast() {
	for {
		socketChannel := <-ChatSocketChannel
		for _, client := range r.Map[socketChannel.ID].Participant {
			if client.Conn != socketChannel.Client {
				AllChatRooms.Mutex.Lock()
				err := client.Conn.WriteJSON(socketChannel.Data)
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
