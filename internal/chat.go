package internal

import (
	"log"
	"net/http"

	"github.com/ddr4869/msazoom/internal/domain/rtc"
	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

var broadcast = rtc.Broadcast

func (s *Server) CreateChat(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateChatRequest)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	chat, err := s.repository.CreateChat(c, req.Title, req.Username)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to create chat")
	}
	rtc.AllRooms.InsertIntoRoom(chat.ID, req.Username, true, ws)

	go rtc.Broadcaster()

	for {
		var msg rtc.BroadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Quit and Delete room")
			ws.Close()
			isDelete := rtc.AllRooms.QuitRoom(chat.ID, req.Username)
			if isDelete {
				err = s.repository.DeleteChat(c, chat.ID)
				if err != nil {
					dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to delete chat")
				}
			}
			break
		}
		msg.Client = ws
		msg.BoardId = chat.ID
		log.Println("message -> ", msg.Message)
		broadcast <- msg
	}
	dto.NewSuccessResponse(c, chat)

}

func (s *Server) JoinChat(c *gin.Context) {
	// get roomID from query
	req := c.MustGet("req").(dto.JoinChatRequest)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}
	rtc.AllRooms.InsertIntoRoom(req.ChatID, req.Username, false, ws)
	go rtc.Broadcaster()
	for {
		var msg rtc.BroadcastMsg
		err := ws.ReadJSON(&msg.Message)
		if err != nil {
			log.Println("Quit and Delete room")
			ws.Close()
			isDelete := rtc.AllRooms.QuitRoom(req.ChatID, req.Username)
			if isDelete {
				err = s.repository.DeleteChat(c, req.ChatID)
				if err != nil {
					dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to delete chat")
				}
			}
			break
		}
		msg.Client = ws
		msg.BoardId = req.ChatID
		log.Println("message -> ", msg.Message)
		broadcast <- msg
	}
}

func (s *Server) GetChatList(c *gin.Context) {
	// show chat list from db
	chats, err := s.repository.GetChatList(c)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get chat list")
	}
	dto.NewSuccessResponse(c, chats)
}
