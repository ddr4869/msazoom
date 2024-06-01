package rtc

import (
	"fmt"
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Participant struct {
	Host bool
	Conn *websocket.Conn
}

type RoomMap struct {
	Mutex sync.RWMutex
	Map   map[int][]Participant
}

var Broadcast = make(chan BroadcastMsg)

func (r *RoomMap) Init() {
	r.Map = make(map[int][]Participant)
}

func (r *RoomMap) Get(board_id int) []Participant {
	r.Mutex.RLock()
	defer r.Mutex.RUnlock()

	return r.Map[board_id]
}

// #### Signaling ####
var AllRooms RoomMap

type BroadcastMsg struct {
	Message map[string]interface{}
	BoardId int
	Client  *websocket.Conn
}

func (r *RoomMap) CreateRoom(board_id int) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Map[board_id] = []Participant{}
}

func (r *RoomMap) InsertIntoRoom(board_id int, host bool, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	p := Participant{host, conn}

	log.Println("Inserting into Room with RoomID: ", board_id)
	r.Map[board_id] = append(r.Map[board_id], p)
}

func (r *RoomMap) DeleteRoom(board_id int) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	delete(r.Map, board_id)
}

func (r *RoomMap) DeleteRoomClient(board_id int, conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	var newParticipants []Participant
	for _, p := range r.Map[board_id] {
		if p.Conn != conn {
			newParticipants = append(newParticipants, p)
		}
	}
	r.Map[board_id] = newParticipants
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
					client.Conn.Close()
					return
				}
				AllRooms.Mutex.Unlock()
			}
		}
	}
}
