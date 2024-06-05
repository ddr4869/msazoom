package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetChatValid(c *gin.Context) {
	var reqUri dto.GetRoomRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("reqUri", reqUri)
	c.Next()
}

func (s *Server) CreateChatValid(c *gin.Context) {
	var req dto.CreateChatRequest
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) JoinChatValid(c *gin.Context) {
	var req dto.JoinChatRequest
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}
