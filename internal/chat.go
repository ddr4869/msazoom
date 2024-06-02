package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ddr4869/msazoom/internal/domain/rtc"
	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

var broadcast = rtc.Broadcast

func (s *Server) CreateMochChat(c *gin.Context) {
	_, _ = s.repository.CreateChat(c, "test1", "tom")
	_, _ = s.repository.CreateChat(c, "test2", "tom")
	_, _ = s.repository.CreateChat(c, "test3", "tom")
	chat, _ := s.repository.CreateChat(c, "test4", "tom")
	dto.NewSuccessResponse(c, chat)
}

func (s *Server) RoomConditionTest(c *gin.Context) {
	dto.NewSuccessResponse(c, &rtc.AllRooms)
}

func (s *Server) CreateChat(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateChatRequest)
	chat, err := s.repository.CreateChat(c, req.Title, req.Username)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to create chat")
	}
	fmt.Println("Chat Created: ", chat.ID)
	dto.NewSuccessResponse(c, chat)
}

func (s *Server) JoinChat(c *gin.Context) {
	// get roomID from query
	req := c.MustGet("req").(dto.JoinChatRequest)
	fmt.Println("req.ChatID: ", req.ChatID)
	fmt.Println("req.Username: ", req.Username)
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}

	chat, err := s.repository.GetChat(c, req.ChatID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get chat")
		return
	}

	var isHost bool
	if chat.ChatUser == req.Username {
		isHost = true
	} else {
		isHost = false
	}

	rtc.AllRooms.InsertIntoRoom(req.ChatID, req.Username, isHost, ws)

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
