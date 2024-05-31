package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
)

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

	var messageResponse []dto.MessageResponse
	for _, message := range messages {
		messageResponse = append(messageResponse, dto.MessageEntToResponse(message))
	}
	dto.NewSuccessResponse(c, messageResponse)
}
