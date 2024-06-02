package dto

type CreateChatRequest struct {
	Title    string `form:"title" binding:"required"`
	Username string `form:"username" binding:"required"`
}

type JoinChatRequest struct {
	ChatID   int    `form:"chat_id" binding:"required"`
	Username string `form:"username" binding:"required"`
}
