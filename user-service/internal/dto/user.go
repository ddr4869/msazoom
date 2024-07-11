package dto

import "github.com/ddr4869/msazoom/user-service/ent"

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type UserLoginResponse struct {
	Username    string `json:"username"`
	Role        int    `json:"role"`
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}

type AddFriendRequest struct {
	Friend string `json:"friend"`
}

type CheckFriendRequest struct {
	Friend string `form:"friend" binding:"required"`
}

type RemoveFriendRequest struct {
	Friend string `form:"friend" binding:"required"`
}

type UserNormalResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
}

type UserWithMessageResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role     int    `json:"role"`
	Email    string `json:"email"`
	Message  UnreadMessage
}

type UnreadMessage struct {
	UnreadMessageCount int32 `json:"unread_message_count"`
}

func UserEntToResponse(user *ent.User) UserNormalResponse {
	return UserNormalResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
	}
}

func UserEntToResponseWithMessage(user *ent.User, msg UnreadMessage) UserWithMessageResponse {
	return UserWithMessageResponse{
		ID:       user.ID,
		Username: user.Username,
		Role:     user.Role,
		Email:    user.Email,
		Message:  msg,
	}
}
