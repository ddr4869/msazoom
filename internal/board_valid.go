package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateBoardValid(c *gin.Context) {
	var req dto.CreateBoardRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) GetBoardWithIDValid(c *gin.Context) {
	var reqUri dto.GetBoardWithIDUriRequest
	if err := c.ShouldBindUri(&reqUri); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("reqUri", reqUri)
	c.Next()
}

func (s *Server) DeleteBoardValid(c *gin.Context) {
	var req dto.GetBoardWithIDUriRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) RecommendBoardValid(c *gin.Context) {
	var req dto.RecommendBoardRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get block number")
		return
	}
	c.Set("req", req)
	c.Next()
}
