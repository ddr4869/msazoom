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

func (r Repository) GetFriendList(ctx context.Context, user_name string) ([]*ent.User, error) {
	u, err := r.entClient.User.
		Query().
		Where(user.UsernameEQ(user_name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	friends, err := u.QueryFollwer().All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	return friends, nil
}

func (r Repository) AddFriend(ctx context.Context, user_name, friend_name string) (*ent.User, error) {
	f, err := r.entClient.User.
		Query().
		Where(user.UsernameEQ(friend_name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	_, err = r.entClient.User.Update().Where(user.UsernameEQ(user_name)).AddFollwer(f).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	return f, nil
}

func (r Repository) RemoveFriend(ctx context.Context, user_name, friend_name string) (bool, error) {
	f, err := r.entClient.User.
		Query().
		Where(user.UsernameEQ(friend_name)).
		Only(ctx)
	if err != nil {
		return false, fmt.Errorf("failed creating board: %w", err)
	}
	_, err = r.entClient.User.Update().Where(user.UsernameEQ(user_name)).RemoveFollwer(f).Save(ctx)
	if err != nil {
		return false, fmt.Errorf("failed creating board: %w", err)
	}
	return true, nil
}

func (r Repository) UpdateRole(ctx context.Context, user_name string, role int) (*ent.User, error) {
	u, err := r.entClient.User.
		Query().
		Where(user.UsernameEQ(user_name)).
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	u, err = u.Update().SetRole(role).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	return u, nil
}

func (r Repository) CheckFriend(ctx context.Context, user_name, friend_name string) (bool, error) {
	is_friend, err := r.entClient.User.
		Query().
		Where(user.UsernameEQ(user_name)).
		Where(user.HasFollwerWith(user.Username(friend_name))).Exist(ctx)

	if err != nil {
		return false, fmt.Errorf("failed creating board: %w", err)
	}
	return is_friend, nil
}
