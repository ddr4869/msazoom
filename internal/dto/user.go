package dto

type UserLoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserCreateRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	Username    string `json:"username"`
	Role        int    `json:"role"`
	AccessToken string `json:"access_token"`
}

type AddFriendRequest struct {
	Friend string `json:"friend"`
}

type CheckFriendRequest struct {
	Friend string `form:"friend" binding:"required"`
}
