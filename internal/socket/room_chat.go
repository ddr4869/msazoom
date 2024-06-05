package socket

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[int]Participants
}

// define map[string]Participant type
type Participants struct {
	Participant map[string]Participant
	Title       string
	Admin       string
	Created_at  time.Time
}

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// #### Signaling ####
var AllRooms RoomMap

func (r *RoomMap) Init() {
	r.Map = make(map[int]Participants)
}

func (r *RoomMap) GetRoom(chat_id int) Participants {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()
	// check if the room exists
	if _, ok := r.Map[chat_id]; !ok {
		return Participants{}
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
	for key := range r.Map {
		keys = append(keys, key)
	}
	return keys[rand.Intn(len(r.Map))]
}

// If the room does not exist, create a new room
// If the room already exists, insert the user into the room
func (r *RoomMap) InsertIntoRoom(chat_id int, chat_title, username string, host bool, ws *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	now := time.Now()

	p := Participant{
		Host: host,
		Conn: ws,
	}
	if host {
		r.Map[chat_id] = Participants{
			Participant: map[string]Participant{username: p},
			Title:       chat_title,
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
		socketChannel := <-SocketChannel
		fmt.Println("broadcast msg -> ", socketChannel)
		for _, client := range r.Map[socketChannel.ID].Participant {
			if client.Conn != socketChannel.Client {
				AllRooms.Mutex.Lock()
				err := client.Conn.WriteJSON(socketChannel.Data)
				if err != nil {
					fmt.Println("broadcast close ->", err)
					client.Conn.Close()
					AllRooms.Mutex.Unlock()
					return
				}
				AllRooms.Mutex.Unlock()
			}
		}
	}
}
