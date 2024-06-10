package dto

type GetRoomRequest struct {
	ID int `uri:"chat_id" binding:"required"`
}

type CreateChatRequest struct {
	Title    string `form:"title" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password"`
}

type JoinChatRequest struct {
	ChatID   int    `form:"chat_id" binding:"required"`
	Username string `form:"username" binding:"required"`
	Password string `form:"password"`
}

type ChatResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Admin      string `json:"admin"`
	Private    bool   `json:"private"`
	Created_at string `json:"created_at"`
}

type CheckPasswordRequest struct {
	ChatID   int    `json:"chat_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}
