package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
)

// func (s *Server) WriteFriendMessage(c *gin.Context) {
// 	req := c.MustGet("req").(dto.JoinMessageRequest)

// 	message, err := s.repository.WriteFriendMessage(c, req.UserName, req.FriendName, req.Message)
// 	if err != nil {
// 		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to write board message")
// 		return
// 	}
// 	dto.NewSuccessResponse(c, dto.MessageEntToResponse(message))
// }

func (s *Server) GetFriendMessage(c *gin.Context) {
	req := c.MustGet("req").(dto.GetFriendMessageRequest)
	claim := c.MustGet("claim").(*utils.UserClaims)

	messages, err := s.repository.GetFriendMessage(c, claim.Name, req.FriendName)
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
	// ws, err := socket.Upgrader.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to upgrade connection")
	// 	return
	// }
	// defer ws.Close()

	// for {
	// 	var msg dto.MessageResponse
	// 	messageType, p, err := ws.ReadJSON()

	// 	fmt.Println("new message -> ", string(p))
	// 	if err != nil {
	// 		log.Printf("conn.ReadMessage: %v", err)
	// 		return
	// 	}
	// 	if err := conn.WriteMessage(messageType, p); err != nil {
	// 		log.Printf("conn.WriteMessage: %v", err)
	// 		return
	// 	}
	// }
}
