package internal

import (
	"net/http"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/gin-gonic/gin"
)

func (s *Server) UserLoginValid(c *gin.Context) {
	var req dto.UserLoginRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "ID or PW is empty")
		return
	}
	c.Set("req", req)
	c.Next()
}

func (s *Server) UserCreateValid(c *gin.Context) {
	var req dto.UserCreateRequest
	if err := c.ShouldBindBodyWithJSON(&req); err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "ID or PW is empty")
		return
	}
	c.Set("req", req)
	c.Next()
}
