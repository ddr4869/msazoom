package dto

type CreateChatRequest struct {
	Board_id int `form:"board_id" binding:"required"`
}
