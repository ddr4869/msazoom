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
	user, err := s.repository.CreateUser(c, req.Username, hash, req.Email)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to create user")
		return
	}
	resp := dto.UserNormalResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
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
	resp := make([]dto.UserNormalResponse, 0)
	for _, friend := range friends {
		resp = append(resp, dto.UserEntToResponse(friend))
	}
	dto.NewSuccessResponse(c, resp)
}

func (s *Server) GetFriendRequestList(c *gin.Context) {
	claims := c.MustGet("claims").(*utils.UserClaims)
	friends, err := s.repository.GetFriendRequestList(c, claims.Name)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to get friend request list")
		return
	}
	resp := make([]dto.UserNormalResponse, 0)
	for _, friend := range friends {
		resp = append(resp, dto.UserEntToResponse(friend))
	}
	dto.NewSuccessResponse(c, resp)
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
	dto.NewSuccessResponse(c, dto.UserEntToResponse(user))
}

func (s *Server) CheckFriend(c *gin.Context) {
	req := c.MustGet("req").(dto.CheckFriendRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)
	if req.Friend == claims.Name {
		dto.NewErrorResponse(c, http.StatusBadRequest, nil, "cannot add yourself as friend")
		return
	}
	is_friend, err := s.repository.CheckFriend(c, claims.Name, req.Friend)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to check friend")
		return
	}
	dto.NewSuccessResponse(c, is_friend)
}

func (s *Server) RemoveFriend(c *gin.Context) {
	req := c.MustGet("req").(dto.RemoveFriendRequest)
	claims := c.MustGet("claims").(*utils.UserClaims)
	_, err := s.repository.RemoveFriend(c, claims.Name, req.Friend)
	if err != nil {
		dto.NewErrorResponse(c, http.StatusBadRequest, err, "failed to remove friend")
		return
	}
	dto.NewSuccessResponse(c, "success")
}
