package repository

import (
	"context"
	"fmt"

	"github.com/ddr4869/msazoom/ent"
	"github.com/ddr4869/msazoom/ent/chat"
)

func (r Repository) CreateChat(ctx context.Context, title, username, hash string) (*ent.Chat, error) {
	c, err := r.entClient.Chat.
		Create().
		SetChatName(title).
		SetChatPassword(hash).
		SetChatUser(username).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	return c, nil
}

func (r Repository) GetChat(ctx context.Context, chat_id int) (*ent.Chat, error) {
	chat, err := r.entClient.Chat.Query().Where(chat.ID(chat_id)).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat list: %w", err)
	}
	return chat, nil
}

func (r Repository) GetChatList(ctx context.Context) ([]*ent.Chat, error) {
	chatList, err := r.entClient.Chat.Query().Order(ent.Desc(chat.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat list: %w", err)
	}
	return chatList, nil
}

func (r Repository) DeleteChat(ctx context.Context, id int) error {
	err := r.entClient.Chat.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return fmt.Errorf("failed to delete chat: %w", err)
	}
	return nil
}
