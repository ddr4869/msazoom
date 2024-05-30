package internal

import (
	"fmt"
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) CreateBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateBoardRequest)
	board, err := s.repository.CreateBoard(c, req.BoardName, req.BoardAdmin, req.BoardPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dto.NewSuccessResponse(c, board)
}

func (s *Server) GetBoardList(c *gin.Context) {
	boardList, err := s.repository.GetBoardList(c)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get board list")
		return
	}
	dto.NewSuccessResponse(c, boardList)
}

func (s *Server) GetBoardWithID(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetBoardWithIDRequest)
	fmt.Println(reqUri.BoardID)
	board, err := s.repository.GetBoardWithID(c, reqUri.BoardID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get board")
		return
	}
	dto.NewSuccessResponse(c, board)
}
