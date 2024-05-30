package dto

import "github.com/ddr4869/msazoom/ent"

type CreateBoardRequest struct {
	BoardName     string `json:"board_name"`
	BoardAdmin    string `json:"board_admin"`
	BoardPassword string `json:"board_password"`
}

type GetBoardWithIDUriRequest struct {
	BoardID       int    `json:"board_id" binding:"required"`
	BoardPassword string `json:"board_password" binding:"required"`
}

type RecommendBoardRequest struct {
	BoardID int `json:"board_id"`
}

type BoardResponse struct {
	BoardID       int    `json:"id"`
	BoardName     string `json:"board_name"`
	BoardAdmin    string `json:"board_admin"`
	BoardPassword string `json:"board_password"`
	BoardStar     int    `json:"board_star"`
}

// convert ent.Board to dto.BoardResponse
func BoardEntToResponse(board *ent.Board) BoardResponse {
	return BoardResponse{
		BoardID:       board.ID,
		BoardName:     board.BoardName,
		BoardAdmin:    board.BoardAdmin,
		BoardPassword: board.BoardPassword,
		BoardStar:     board.BoardStar,
	}
}
