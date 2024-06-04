package internal

import (
	"fmt"
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

func (s *Server) GetFriendList(c *gin.Context) {
	claims := c.MustGet("claims").(*utils.UserClaims)
	friends, err := s.repository.GetFriendList(c, claims.Name)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get friend list")
		return
	}
	fmt.Println("claims.Name: ", claims.Name)
	fmt.Println("friends: ", friends)
	dto.NewSuccessResponse(c, friends)
}

func (s *Server) AddFriend(c *gin.Context) {
	req := c.MustGet("req").(dto.AddFriendRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)
	if req.Friend == claims.Name {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "cannot add yourself as friend")
		return
	}
	user, err := s.repository.AddFriend(c, claims.Name, req.Friend)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to add friend")
		return
	}
	dto.NewSuccessResponse(c, user)
}

func (s *Server) CheckFriend(c *gin.Context) {
	req := c.MustGet("req").(dto.CheckFriendRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)
	if req.Friend == claims.Name {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "cannot add yourself as friend")
		return
	}
	user, err := s.repository.CheckFriend(c, claims.Name, req.Friend)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to check friend")
		return
	}
	dto.NewSuccessResponse(c, user)
}
