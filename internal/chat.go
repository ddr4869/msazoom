package internal

import (
	"log"
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/socket"
	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
)

//var broadcast = socket.ChatSocketChannel

func (s *Server) RoomConditionCheck(c *gin.Context) {
	dto.NewSuccessResponse(c, &socket.AllChatRooms)
}

func (s *Server) GetChat(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetRoomRequest)
	p := socket.AllChatRooms.GetRoom(reqUri.ID)
	if p.Participant == nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "No room available")
		return
	}
	dto.NewSuccessResponse(c, p)
}

func (s *Server) CreateChat(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateChatRequest)
	hash := ""
	var err error
	if req.Password != "" {
		hash, err = utils.HashPassword(req.Password)
		if err != nil {
			dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to hash password")
			return
		}
	}
	chat, err := s.repository.CreateChat(c, req.Title, req.Username, hash)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to create chat")
	}
	dto.NewSuccessResponse(c, chat)
}

func (s *Server) JoinChat(c *gin.Context) {
	req := c.MustGet("req").(dto.JoinChatRequest)
	ws, err := socket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Fatal("Web Socket Upgrade Error", err)
	}
	chat, err := s.repository.GetChat(c, req.ChatID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get chat")
		return
	}
	if req.Password != "" && !utils.CheckPasswordHash(req.Password, chat.ChatPassword) {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "Password is incorrect")
		return
	}

	if chat.ChatUser == req.Username {
		if chat.ChatPassword != "" {
			socket.AllChatRooms.InsertIntoRoom(req.ChatID, chat.ChatName, req.Username, true, true, ws)
		} else {
			socket.AllChatRooms.InsertIntoRoom(req.ChatID, chat.ChatName, req.Username, false, true, ws)
		}
	} else {
		if chat.ChatPassword != "" {
			socket.AllChatRooms.InsertIntoRoom(req.ChatID, chat.ChatName, req.Username, true, false, ws)
		} else {
			socket.AllChatRooms.InsertIntoRoom(req.ChatID, chat.ChatName, req.Username, false, false, ws)
		}
	}

	go socket.AllChatRooms.Broadcast()

	for {
		var socketData socket.ChatSocketData
		err := ws.ReadJSON(&socketData.Data)
		if err != nil {
			ws.Close()
			isDelete := socket.AllChatRooms.QuitRoom(req.ChatID, req.Username)
			if isDelete {
				err = s.repository.DeleteChat(c, req.ChatID)
				if err != nil {
					dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to delete chat on DB")
				}
			}
			break
		}
		socketData.Client = ws
		socketData.ID = req.ChatID
		socket.ChatSocketChannel <- socketData
	}
}

func (s *Server) GetChatList(c *gin.Context) {
	ChatResponse := make([]dto.ChatResponse, 0)
	for id, chat := range socket.AllChatRooms.Map {
		if len(chat.Participant) != 1 {
			continue
		}
		ChatResponse = append(ChatResponse, dto.ChatResponse{
			ID:         id,
			Title:      chat.Title,
			Admin:      chat.Admin,
			Private:    chat.Private,
			Created_at: chat.Created_at.Format("2006-01-02 15:04:05"),
		})
	}
	dto.NewSuccessResponse(c, ChatResponse)
}

func (s *Server) CheckPassword(c *gin.Context) {
	req := c.MustGet("req").(dto.CheckPasswordRequest)
	chat, err := s.repository.GetChat(c, req.ChatID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get chat")
		return
	}
	if !utils.CheckPasswordHash(req.Password, chat.ChatPassword) {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "Password is incorrect")
		return
	}
	dto.NewSuccessResponse(c, true)
}

func (s *Server) RandomChating(c *gin.Context) {
	// get roomID from query
	n := socket.AllChatRooms.GetRandomRoomKey()
	if n == -1 {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "No room available")
		return
	}
	dto.NewSuccessResponse(c, n)
}
