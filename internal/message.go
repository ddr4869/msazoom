package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/socket"
	"github.com/ddr4869/msazoom/internal/utils"
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

func (s *Server) ConnectMessage(c *gin.Context) {
	req := c.MustGet("req").(dto.ConnectMessageRequest)
	ws, err := socket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to upgrade connection")
		return
	}
	defer ws.Close()

	key := GenerateSocketKey(req.UserName, req.FriendName)
	socket.AllMessageRooms.InsertIntoRoom(key, req.UserName, req.FriendName, ws)

	go socket.AllMessageRooms.Broadcast()
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
		fmt.Println("socketData.Client set, req.username", req.UserName)

		socket.MessageSocketChannel <- socketData
	}
}

func GenerateSocketKey(key1, key2 string) string {
	if key1 > key2 {
		return key1 + key2
	} else {
		return key2 + key1
	}
}
