package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type CustomClaimsExample struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

func (s *Server) CreateBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.CreateBoardRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)

	board, err := s.repository.CreateBoard(c, req.BoardName, claims.Name, req.BoardPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	dto.NewSuccessResponse(c, dto.BoardEntToResponse(board))
}

func (s *Server) GetBoardList(c *gin.Context) {
	boardList, err := s.repository.GetBoardList(c)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to get board list")
		return
	}
	boardResponse := make([]dto.BoardResponse, 0)
	for _, board := range boardList {
		boardResponse = append(boardResponse, dto.BoardEntToResponse(board))
	}
	dto.NewSuccessResponse(c, boardResponse)
}

func (s *Server) GetBoardWithID(c *gin.Context) {
	reqUri := c.MustGet("reqUri").(dto.GetBoardWithIDUriRequest)
	board, err := s.repository.GetBoardWithID(c, reqUri.BoardID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get board")
		return
	}
	dto.NewSuccessResponse(c, dto.BoardEntToResponse(board))
}

func (s *Server) RecommendBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.RecommendBoardRequest)

	board, err := s.repository.RecommendBoard(c, req.BoardID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to recommend board")
		return
	}
	dto.NewSuccessResponse(c, dto.BoardEntToResponse(board))
}

func (s *Server) DeleteBoard(c *gin.Context) {
	req := c.MustGet("req").(dto.DeleteBoardWithIDUriRequest)
	board, err := s.repository.GetBoardWithID(c, req.BoardID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get board")
		return
	}
	if board.BoardAdmin != c.MustGet("claims").(*utils.UserClaims).Name {
		dto.NewErrorResponse(c, http.StatusUnauthorized, nil, "you are not the admin of this board")
		return
	}
	if board.BoardPassword != req.BoardPassword {
		dto.NewErrorResponse(c, http.StatusUnauthorized, nil, "invalid password")
		return
	}
	err = s.repository.DeleteBoard(c, req.BoardID)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to delete board")
		return
	}
	dto.NewSuccessResponse(c, "success")
}
