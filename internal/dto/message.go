package dto

import "github.com/ddr4869/msazoom/ent"

type WriteBoardMessageRequest struct {
	BoardID int    `json:"board_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

type GetBoardMessageRequest struct {
	BoardID int `uri:"board_id"`
}

type MessageResponse struct {
	MessageID int    `json:"id"`
	BoardID   int    `json:"board_id"`
	Message   string `json:"message"`
	Writer    string `json:"writer"`
}

func MessageEntToResponse(message *ent.Message) MessageResponse {
	return MessageResponse{
		MessageID: message.ID,
		BoardID:   message.BoardID,
		Message:   message.Message,
		Writer:    message.Writer,
	}
}
