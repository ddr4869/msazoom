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

func (s *Server) RoomConditionCheck(c *gin.Context) {
	dto.NewSuccessResponse(c, &rtc.AllRooms)
}

func (s *Server) GetChat(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetRoomRequest)
	p := rtc.AllRooms.GetRoom(reqUri.ID)
	if p.Participant == nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "No room available")
		return
	}
	dto.NewSuccessResponse(c, p)
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

	rtc.AllRooms.InsertIntoRoom(req.ChatID, chat.ChatName, req.Username, isHost, ws)

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
	ChatResponse := make([]dto.ChatResponse, 0)
	for id, chat := range rtc.AllRooms.Map {
		ChatResponse = append(ChatResponse, dto.ChatResponse{
			ID:         id,
			Title:      chat.Title,
			Admin:      chat.Admin,
			Created_at: chat.Created_at.Format("2006-01-02 15:04:05"),
		})
	}
	dto.NewSuccessResponse(c, ChatResponse)
}

func (s *Server) RandomChating(c *gin.Context) {
	// get roomID from query
	n := rtc.AllRooms.GetRandomRoomKey()
	if n == -1 {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "No room available")
		return
	}
	dto.NewSuccessResponse(c, n)
}
