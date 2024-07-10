package internal

import (
	"log"
	"net/http"
	"slices"
	"strings"

	"github.com/ddr4869/msazoom/message-service/internal/dto"
	"github.com/ddr4869/msazoom/message-service/internal/socket"
	"github.com/ddr4869/msazoom/message-service/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) MessageConditionCheck(c *gin.Context) {
	dto.NewSuccessResponse(c, &socket.AllMessageRooms)
}

func (s *Server) GetFriendMessage(c *gin.Context) {
	req := c.MustGet("req").(dto.GetFriendMessageRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)

	messages, err := s.repository.GetFriendMessage(c, claims.Name, req.FriendName)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get board message")
		return
	}
	messageResponse := make([]dto.MessageResponse, 0)
	for _, message := range messages {
		messageResponse = append(messageResponse, dto.MessageEntToResponse(message))
	}
	dto.NewSuccessResponse(c, messageResponse)
}

func (s *Server) GetNumberOfUnreadMessage(c *gin.Context) {
	req := c.MustGet("req").(dto.GetFriendMessageRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)

	count, err := s.repository.GetNumberOfUnreadMessage(c, req.FriendName, claims.Name)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get number of unread message")
		return
	}
	dto.NewSuccessResponse(c, count)
}

func (s *Server) ConnectMessage(c *gin.Context) {
	req := c.MustGet("req").(dto.ConnectMessageRequest)
	ws, err := socket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to upgrade connection")
		return
	}
	defer ws.Close()

	key := GenerateSocketKey([]string{req.UserName, req.FriendName})
	socket.AllMessageRooms.InsertIntoRoom(key, req.UserName, req.FriendName, ws)

	go socket.AllMessageRooms.Broadcast(c, s.repository)
	var socketData socket.MessageSocketData
	socketData.Client = ws
	socketData.ID = key

	for {
		err := ws.ReadJSON(&socketData.Data)
		if err != nil {
			log.Println("Quit or Delete room")
			ws.Close()
			_ = socket.AllMessageRooms.QuitRoom(key, req.UserName)
			break
		}
		socket.MessageSocketChannel <- socketData
	}
}

func GenerateSocketKey(keys []string) string {
	slices.Sort(keys)
	return strings.Join(keys, "")
}
