package internal

import (
	"fmt"
	"log"

	"github.com/ddr4869/msazoom/internal/domain/rtc"
	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

var broadcast = rtc.Broadcast

func (s *Server) CreateChat(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateChatRequest)
	rtc.AllRooms.CreateRoom(req.Board_id)

	dto.NewSuccessResponse(c, req.Board_id)
}

func (s *Server) RoomListTest(c *gin.Context) {
	dto.NewSuccessResponse(c, rtc.AllRooms.Map)
}

func (s *Server) JoinRoomTest(c *gin.Context) {
	// get roomID from query
	req := c.MustGet("req").(dto.CreateChatRequest)
	fmt.Println("req->", req)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}
	rtc.AllRooms.InsertIntoRoom(req.Board_id, false, ws)
	go rtc.Broadcaster()
	for {
		var msg rtc.BroadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Quit")
			break
		}
		msg.Client = ws
		msg.BoardId = req.Board_id
		log.Println("message -> ", msg.Message)
		broadcast <- msg
	}
}
