package dto

import "github.com/ddr4869/msazoom/ent"

type GetFriendMessageRequest struct {
	FriendName string `form:"friend_name" binding:"required"`
}

type ConnectMessageRequest struct {
	UserName   string `form:"user_name" binding:"required"`
	FriendName string `form:"friend_name" binding:"required"`
}

type MessageResponse struct {
	MessageID int    `json:"id"`
	Message   string `json:"message"`
	Writer    string `json:"writer"`
	Sender    string `json:"sender"`
}

func MessageEntToResponse(message *ent.Message) MessageResponse {
	return MessageResponse{
		MessageID: message.ID,
		Message:   message.Message,
		Writer:    message.Sender,
		Sender:    message.Receiver,
	}
}
