package repository

import (
	"context"

	"github.com/ddr4869/msazoom/ent"
	"github.com/ddr4869/msazoom/ent/message"
)

func (r *Repository) WriteBoardMessage(ctx context.Context, boardID int, username string, message string) (*ent.Message, error) {
	msg, err := r.entClient.Message.Create().
		SetBoardID(boardID).
		SetWriter(username).
		SetMessage(message).
		Save(context.Background())
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (r *Repository) GetBoardMessage(ctx context.Context, boardID int) ([]*ent.Message, error) {
	messages, err := r.entClient.Message.Query().
		Where(message.BoardID(boardID)).
		All(context.Background())
	if err != nil {
		return nil, err
	}
	return messages, nil
}
