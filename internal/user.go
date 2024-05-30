package internal

import (
	"net/http"
	"strconv"

	"github.com/ddr4869/msazoom/internal/dto"
	"github.com/ddr4869/msazoom/internal/utils"
	"github.com/gin-gonic/gin"
)

func (s *Server) UserLogin(c *gin.Context) {
	req := c.MustGet("req").(dto.UserLoginRequest)
	user, err := s.repository.GetUser(c, req.Username)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get user")
		return
	}
	if !utils.CheckPasswordHash(req.Password, user.Password) {
		dto.NewErrorResponse(c, http.StatusUnauthorized, nil, "invalid password")
		return
	}
	token, err := utils.GenerateJWT(user.Username, strconv.Itoa(user.Role))
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to generate token")
		return
	}
	resp := dto.UserLoginResponse{
		Username:    user.Username,
		Role:        user.Role,
		AccessToken: token,
	}
	dto.NewSuccessResponse(c, resp)
}

func (s *Server) UserCreate(c *gin.Context) {
	req := c.MustGet("req").(dto.UserCreateRequest)

	hash, err := utils.HashPassword(req.Password)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusInternalServerError, err, "failed to hash password")
		return
	}
	user, err := s.repository.CreateUser(c, req.Username, hash)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to create user")
		return
	}
	resp := dto.UserLoginResponse{
		Username: user.Username,
		Role:     user.Role,
	}
	dto.NewSuccessResponse(c, resp)
}
