package dto

import "github.com/ddr4869/msazoom/message-service/ent"

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
	Sender    string `json:"sender"`
	Receiver  string `json:"receiver"`
}

func MessageEntToResponse(message *ent.Message) MessageResponse {
	return MessageResponse{
		MessageID: message.ID,
		Message:   message.Message,
		Sender:    message.Sender,
		Receiver:  message.Receiver,
	}
}
