package internal

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/ddr4869/msazoom/internal/domain/rtc"
	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

var broadcast = rtc.Broadcast

func (s *Server) CreateChat(c *gin.Context) {

	roomID := rtc.AllRooms.CreateRoom()

	type resp struct {
		RoomID string `json:"room_id"`
	}

	log.Println("AllRooms.Map -> ", rtc.AllRooms.Map)
	json.NewEncoder(c.Writer).Encode(resp{RoomID: roomID})
}

func (s *Server) RoomListTest(c *gin.Context) {
	dto.NewSuccessResponse(c, rtc.AllRooms.Map)
}

func (s *Server) JoinRoomTest(c *gin.Context) {
	// get roomID from query
	req := c.MustGet("req").(RoomID)
	fmt.Println("req->", req)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}
	// get first key of AllRooms.Map
	var roomId string
	for k := range rtc.AllRooms.Map {
		roomId = k
		rtc.AllRooms.InsertIntoRoom(roomId, false, ws)
		break
	}
	//AllRooms.InsertIntoRoom("test", false, ws)

	go rtc.Broadcaster()
	for {
		var msg rtc.BroadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Quit")
			break
		}
		fmt.Println("roomId -> ", roomId)
		msg.Client = ws
		msg.RoomID = roomId
		log.Println("message -> ", msg.Message)
		broadcast <- msg

	}
}
