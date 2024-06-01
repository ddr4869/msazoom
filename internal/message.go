package internal

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // 모든 요청을 허용합니다. 보안상의 이유로 변경해야 할 수 있습니다.
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (s *Server) WriteBoardMessage(c *gin.Context) {
	req := c.MustGet("req").(dto.WriteBoardMessageRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)

	message, err := s.repository.WriteBoardMessage(c, req.BoardID, claims.Name, req.Message)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to write board message")
		return
	}
	dto.NewSuccessResponse(c, dto.MessageEntToResponse(message))
}

func (s *Server) GetBoardMessage(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetBoardMessageRequest)

	messages, err := s.repository.GetBoardMessage(c, reqUri.BoardID)
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

func (s *Server) SocketWriteBoardMessage(c *gin.Context) {
	// Upgrade upgrades the HTTP server connection to the WebSocket protocol.
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer conn.Close()

	//conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
	for {
		messageType, p, err := conn.ReadMessage()

		fmt.Println(string(p))
		if err != nil {
			log.Printf("conn.ReadMessage: %v", err)
			return
		}
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("conn.WriteMessage: %v", err)
			return
		}
	}
}
