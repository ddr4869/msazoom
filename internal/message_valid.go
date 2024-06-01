package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) WriteBoardMessageValid(c *gin.Context) {
	var req dto.WriteBoardMessageRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) GetBoardMessageValid(c *gin.Context) {
	var reqUri dto.GetBoardMessageRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("reqUri", reqUri)
	c.Next()
}

type RoomID struct {
	RoomID string `form:"room_id" json:"room_id"`
}

func (s *Server) JoinRoomTestValid(c *gin.Context) {
	var req RoomID
	if err := c.ShouldBind(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}
