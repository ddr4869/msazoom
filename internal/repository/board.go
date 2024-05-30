package repository

import (
	"context"
	"fmt"

	"github.com/ddr4869/msazoom/ent"
	"github.com/ddr4869/msazoom/ent/board"
)

func (r Repository) CreateBoard(ctx context.Context, board_name, board_admin, board_password string) (*ent.Board, error) {
	u, err := r.entClient.Board.
		Create().
		SetBoardName(board_name).
		SetBoardAdmin(board_admin).
		SetBoardPassword(board_password).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating board: %w", err)
	}
	return u, nil
}

func (r Repository) GetBoardList(ctx context.Context) ([]*ent.Board, error) {
	boardList, err := r.entClient.Board.Query().Order(ent.Desc(board.FieldCreatedAt)).All(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get board list: %w", err)
	}
	return boardList, nil
}

func (r Repository) GetBoardWithID(ctx context.Context, id int) (*ent.Board, error) {
	board, err := r.entClient.Board.Query().Where(
		board.IDEQ(id),
	).First(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get board: %w", err)
	}
	return board, nil
}

func (r Repository) RecommendBoard(ctx context.Context, id int) (*ent.Board, error) {
	board, err := r.entClient.Board.UpdateOneID(id).AddBoardStar(1).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to recommend board: %w", err)
	}
	return board, nil
}
