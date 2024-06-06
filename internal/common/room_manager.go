package common

import (
	"time"

	"github.com/gorilla/websocket"
)

var AllMessageRooms = RoomManager{}
var AllChatRooms = RoomManager{}

type MessageRoom struct {
	ChatRoom
	Sender   string
	Receiver string
}

type MeetingRoom struct {
	ChatRoom
	Title string
	Admin string
}

func Init() {
	AllMessageRooms.Init()
	AllChatRooms.Init()
}

func InsertIntoMessageRoom(id, sender, receiver string, ws *websocket.Conn) {
	now := time.Now()
	p := Participant{Conn: ws}
	room := MessageRoom{
		ChatRoom: ChatRoom{
			ID:          id,
			Participant: map[string]Participant{sender: p},
			CreatedAt:   now,
		},
		Sender:   sender,
		Receiver: receiver,
	}
	AllMessageRooms.InsertIntoRoom(id, room)
}

func InsertIntoChatRoom(id, title, admin string, ws *websocket.Conn) {
	now := time.Now()
	p := Participant{Host: true, Conn: ws}
	room := MeetingRoom{
		ChatRoom: ChatRoom{
			ID:          id,
			Participant: map[string]Participant{admin: p},
			CreatedAt:   now,
		},
		Title: title,
		Admin: admin,
	}
	AllChatRooms.InsertIntoRoom(id, room)
}

func QuitMessageRoom(id, username string) {
	AllMessageRooms.QuitRoom(id, username)
}

func QuitChatRoom(id, username string) {
	AllChatRooms.QuitRoom(id, username)
}

func BroadcastMessages() {
	go AllMessageRooms.Broadcast(MessageSocketChannel)
}

func BroadcastChats() {
	go AllChatRooms.Broadcast(ChatSocketChannel)
}
