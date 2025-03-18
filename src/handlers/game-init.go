package handlers

import (
	"log"
	"main/src/db"
	"main/src/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type initGamePayload struct {
	GameID          string              `json:"gameId"`
	Username        string              `json:"username"`
	Player2GameLink string              `json:"player2_gameLink"`
	BoardState      map[string][]string `json:"board_state"`
	Data            map[string]string   `json:"data"`
}

type newUserStruct struct {
	Username    string    `json:"username"`
	TotalPoints int       `json:"total_points"`
	GamesWon    int       `json:"games_won"`
	CreatedAt   time.Time `json:"created_at"`
}

type newGameStruct struct {
	GameID          string              `json:"game_id_pk"`
	Player1Username string              `json:"player1_username"`
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
	var err error
	

	var existingUser []newUserStruct
	err = db.SupaClient.DB.From("users").Select("*").Eq("username", user).Execute(&existingUser)
	if err != nil {
		log.Println("Error fetching existing user from users_t: ", err)
	} else {
		err = db.SupaClient.DB.From("users").Insert(newUser).Execute(&res)
		if err != nil {
			log.Println("An error has been encountered trying to insert to Users_T: ", err)
		}
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
	generatedGameID := utils.CreateNanoID()

	newGame := newGameStruct{
		GameID:          generatedGameID,
		Player1Username: user,
		BoardState:      initBoardState,
		Status:          "ongoing",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = db.SupaClient.DB.From("games").Insert(newGame).Execute(&res)
	if err != nil {
		log.Println("An error has been encountered trying to insert to Games_t: ", err)
	}

	var insertedGames []any
	err = db.SupaClient.DB.From("games").Select("*").Eq("game_id_pk", generatedGameID).Execute(&insertedGames)
	if err != nil {
		log.Println("An error has been encountered trying to fetching from games table: ", err)
	}

	additionalData := map[string]string{
		"instructions": "You have started a game! now send the Player 2 Game Link to the person you want to play with, be sure to tell them to substitute {your-username} with their player username!",
		"how-to-play":  r.Host + `/how-to-play`,
		"Your-piece":   "Black",
		"Welcome":      `Welcome, ` + user + ` to Backend Dam Haji AKA Backend Checkers!`,
	}

	if len(existingUser) > 0 {
		additionalData["Welcome"] = "Welcome Back " + user + ", lets get a W!"
	}

	p := initGamePayload{
		GameID:          generatedGameID,
		Username:        user,
		Player2GameLink: r.Host + `/` + generatedGameID + `/{your-username}`,
		BoardState:      initBoardState,
		Data:            additionalData,
	}
	utils.Serve(w, p)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {
	




	
}
