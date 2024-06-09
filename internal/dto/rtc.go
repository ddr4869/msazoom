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
	Password   string `json:"password"`
	Created_at string `json:"created_at"`
}
