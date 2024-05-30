package repository

import (
	"context"
	"fmt"

	"github.com/ddr4869/msazoom/ent"
	"github.com/ddr4869/msazoom/ent/user"
)

func (r Repository) CreateUser(ctx context.Context, user_name, password_hash string) (*ent.User, error) {
	user, err := r.entClient.User.
		Create().
		SetUsername(user_name).
		SetPassword(password_hash).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("User already exists: %w", err)
	}
	return user, nil
}

func (r Repository) GetUser(ctx context.Context, user_name string) (*ent.User, error) {
	user, err := r.entClient.User.
		Query().
		Where(user.UsernameEQ(user_name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	return user, nil
}
