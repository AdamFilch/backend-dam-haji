package handlers

import (
	"log"
	"net/http"
	"src/main/src/db"
	"src/main/src/utils"
	"time"

	"github.com/gorilla/mux"
)

type initGamePayload struct {
	GameID          string `json:"gameId"`
	Username        string `json:"username"`
	Player2GameLink string `json:"player2_gameLink"`
}

type newUserStruct struct {
	Username    string    `json:"username"`
	TotalPoints int       `json:"total_points"`
	GamesWon    int       `json:"games_won"`
	CreatedAt   time.Time `json:"created_at"`
}

type newGameStruct struct {
	Player1Username string              `json:"player1_username"`
	Player2Username string              `json:"player2_username"`
	BoardState      map[string][]string `json:"board_state"`
	Status          string              `json:"status"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

func HandleInitGame(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["user"]

	newUser := newUserStruct{
		Username:    user,
		TotalPoints: 0,
		GamesWon:    0,
		CreatedAt:   time.Now(),
	}

	var res any
	err := db.SupaClient.DB.From("users_t").Insert(newUser).Execute(&res)
	if err != nil {
		log.Fatal("An error has been encountered trying to insert to Users_T")
	}

	initBoardState := map[string][]string{
		"0": {"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
		"A": {" ", "X", " ", "X", " ", "X", " ", "X", " ", "X"},
		"B": {"X", " ", "X", " ", "X", " ", "X", " ", "X", " "},
		"C": {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		"D": {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		"E": {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		"F": {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		"G": {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		"H": {" ", " ", " ", " ", " ", " ", " ", " ", " ", " "},
		"I": {" ", "0", " ", "0", " ", "0", " ", "0", " ", "0"},
		"J": {"0", " ", "0", " ", "0", " ", "0", " ", "0", " "},
		"Z": {"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"},
	}

	newGame := newGameStruct{
		Player1Username: user,
		Player2Username: "null",
		BoardState:      initBoardState,
		Status:          "ongoing",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = db.SupaClient.DB.From("users_t").Insert(newGame).Execute(&res)
	if err != nil {
		log.Fatal("An error has been encountered trying to insert to Games_t")
	}

	p := initGamePayload{
		GameID:          "3",
		Username:        user,
		Player2GameLink: "test",
	}
	utils.Serve(w, p)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Log Handle Get Game")
}
