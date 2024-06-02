package rtc

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[int]Participants
}

// define map[string]Participant type
type Participants map[string]Participant

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

// #### Signaling ####
var AllRooms RoomMap

type BroadcastMsg struct {
	Message map[string]interface{}
	BoardId int
	Client  *websocket.Conn
}

var Broadcast = make(chan BroadcastMsg)

func (r *RoomMap) Init() {
	r.Map = make(map[int]Participants)
}

func (r *RoomMap) Get(board_id int) Participants {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[board_id]
}

// If the room does not exist, create a new room
// If the room already exists, insert the user into the room
func (r *RoomMap) InsertIntoRoom(chat_id int, username string, host bool, ws *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	p := Participant{
		Host: host,
		Conn: ws,
	}
	log.Println("Inserting into Room with RoomID: ", chat_id)
	if host {
		r.Map[chat_id] = Participants{username: p}
	} else {
		r.Map[chat_id][username] = p
	}
}

// If the user is the host, delete the room
func (r *RoomMap) QuitRoom(board_id int, username string) bool {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	p := r.Map[board_id][username]
	if p.Host {
		delete(r.Map, board_id)
		return true
	} else {
		delete(r.Map[board_id], username)
		return false
	}
}

func Broadcaster() {
	for {
		msg := <-Broadcast
		fmt.Println("broadcast msg -> ", msg)
		for _, client := range AllRooms.Map[msg.BoardId] {
			if client.Conn != msg.Client {
				AllRooms.Mutex.Lock()
				err := client.Conn.WriteJSON(msg.Message)
				if err != nil {
					fmt.Println("broadcast close !!")
					// print error
					log.Println("Error broadcasting message: ", err)
					client.Conn.Close()
					AllRooms.Mutex.Unlock()
					return
				}
				AllRooms.Mutex.Unlock()
			}
		}
	}
}
