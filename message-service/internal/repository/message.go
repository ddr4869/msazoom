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
		SetIsRead(false).
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
	for _, m := range msg {
		if m.Sender == sender {
			continue
		}
		_, err := r.entClient.Message.UpdateOne(m).
			SetIsRead(true).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	return msg, nil
}

func (r *Repository) GetNumberOfUnreadMessage(ctx context.Context, sender, receiver string) (int, error) {
	count, err := r.entClient.Message.Query().
		Where(
			message.And(
				message.Sender(sender),
				message.Receiver(receiver),
				message.IsRead(false),
			),
		).
		Count(ctx)
	if err != nil {
		return 0, err
	}
	return count, nil
}
