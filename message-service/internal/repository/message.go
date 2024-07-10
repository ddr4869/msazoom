package repository

import (
	"context"

	"github.com/ddr4869/msazoom/message-service/ent"
	"github.com/ddr4869/msazoom/message-service/ent/message"
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
		Where(
			message.Or(
				message.And(
					message.Sender(sender),
					message.Receiver(receiver),
				),
				message.And(
					message.Sender(receiver),
					message.Receiver(sender),
				),
			),
		).
		Order(ent.Asc(message.FieldCreatedAt)).
		All(ctx)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
