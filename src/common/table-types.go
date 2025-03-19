package common

import "time"

type TableGameStruct struct {
	GameID               string              `json:"game_id_pk"`
	BlackPlayer1Username string              `json:"black_player1_username"`
	WhitePlayer2Username string              `json:"white_player2_username"`
	BoardState           map[string][]string `json:"board_state"`
	WinnerUsername       string              `json:"winner_username"`
	Status               string              `json:"status"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
}


type TableMovesStruct struct {
	MoveID int `json:"move_id_pk"`
	GameID string `json:"game_id"`
	Username string `json:"username"`
	StartPosition string `json:"move_from"`
	EndPosition string `json:"move_to"`
	CreatedAt time.Time `json:"created_at"`
	PieceColor string `json:"piece_color"`
}