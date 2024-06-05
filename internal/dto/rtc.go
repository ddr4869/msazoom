package dto

type GetRoomRequest struct {
	ID int `uri:"chat_id" binding:"required"`
}

type CreateChatRequest struct {
	Title    string `form:"title" binding:"required"`
	Username string `form:"username" binding:"required"`
}

type JoinChatRequest struct {
	ChatID   int    `form:"chat_id" binding:"required"`
	Username string `form:"username" binding:"required"`
}

type ChatResponse struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Admin      string `json:"admin"`
	Created_at string `json:"created_at"`
}
