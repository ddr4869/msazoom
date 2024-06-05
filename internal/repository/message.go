package repository

import (
	"context"

	"github.com/ddr4869/msazoom/ent"
	"github.com/ddr4869/msazoom/ent/message"
)

func (r *Repository) WriteFriendMessage(ctx context.Context, sender, receiver, message string) (*ent.Message, error) {
	msg, err := r.entClient.Message.Create().
		SetSender(sender).
		SetReceiver(receiver).
		SetMessage(message).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (r *Repository) GetFriendMessage(ctx context.Context, sender, receiver string) ([]*ent.Message, error) {
	msg, err := r.entClient.Message.Query().
		Where(message.Sender(sender)).
		Where(message.Receiver(receiver)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return msg, nil
}
