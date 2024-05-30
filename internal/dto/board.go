package dto

type CreateBoardRequest struct {
	BoardName     string `json:"board_name"`
	BoardAdmin    string `json:"board_admin"`
	BoardPassword string `json:"board_password"`
}

type GetBoardWithIDRequest struct {
	BoardID int `uri:"board_id" binding:"required"`
}
