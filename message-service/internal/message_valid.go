package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/message-service/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) ConnectMessageValid(c *gin.Context) {
	var req dto.ConnectMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) GetFriendMessageValid(c *gin.Context) {
	var req dto.GetFriendMessageRequest
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}
