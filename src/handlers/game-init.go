package handlers

import (
	"log"
	"main/src/common"
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
		log.Println("Error: HandleInitGame - Fetching from users_t: ", err)
	} else {
		err = db.SupaClient.DB.From("users").Insert(newUser).Execute(&res)
		if err != nil {
			log.Println("Error: HandleInitGame - Inserting to Users_T: ", err)
		}
	}

	generatedGameID := utils.CreateNanoID()

	newGame := newGameStruct{
		GameID:          generatedGameID,
		Player1Username: user,
		BoardState:      common.InitBoardState,
		Status:          "ongoing",
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}

	err = db.SupaClient.DB.From("games").Insert(newGame).Execute(&res)
	if err != nil {
		log.Println("Error: HandleInitGame - Inserting to Games_t: ", err)
	}

	var insertedGames []any
	err = db.SupaClient.DB.From("games").Select("*").Eq("game_id_pk", generatedGameID).Execute(&insertedGames)
	if err != nil {
		log.Println("Error: HandleInitGame - Fetching from Games_T: ", err)
	}

	additionalData := map[string]string{
		"instructions":         "You have started a game! now send the Player 2 Game Link to the person you want to play with, be sure to tell them to substitute {your-username} with their player username!",
		"how_to_play":          r.Host + `/how-to-play`,
		"your_piece":           "Black",
		"turn":                 user,
		"welcome":              `Welcome, ` + user + ` to Backend Dam Haji AKA Backend Checkers!`,
		"make_your_first_move": r.Host + `/` + generatedGameID + `/` + user + `/move/{origin}/to/{final}`,
	}

	if len(existingUser) > 0 {
		additionalData["Welcome"] = "Welcome Back " + user + ", lets get a W!"
	}

	p := initGamePayload{
		GameID:          generatedGameID,
		Username:        user,
		Player2GameLink: r.Host + `/` + generatedGameID + `/{your-username}`,
		BoardState:      common.InitBoardState,
		Data:            additionalData,
	}
	utils.Serve(w, p)
}

type selectGameStruct struct {
	GameID               string              `json:"game_id_pk"`
	BlackPlayer1Username string              `json:"black_player1_username"`
	WhitePlayer2Username string              `json:"white_player2_username"`
	BoardState           map[string][]string `json:"board_state"`
	WinnerUsername       string              `json:"winner_username"`
	Status               string              `json:"status"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
}

type startGamePayload struct {
	GameID               string              `json:"gameId"`
	BlackPlayer1Username string              `json:"black_player1_username"`
	WhitePlayer2Username string              `json:"white_player2_username"`
	BoardState           map[string][]string `json:"board_state"`
	UpdatedAt            string              `json:"last_updated_at"`
	Data                 map[string]string   `json:"data"`
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["user"]
	gameID := vars["gameid"]

	var res any
	var err error

	newUser := newUserStruct{
		Username:    user,
		TotalPoints: 0,
		GamesWon:    0,
		CreatedAt:   time.Now(),
	}

	var existingUser []newUserStruct
	err = db.SupaClient.DB.From("users").Select("*").Eq("username", user).Execute(&existingUser)
	if err != nil {
		log.Println("Error: HandleInitGame - Fetching from users_t: ", err)
	} else {
		err = db.SupaClient.DB.From("users").Insert(newUser).Execute(&res)
		if err != nil {
			log.Println("Error: HandleInitGame - Inserting to Users_T: ", err)
		}
	}

	var fetchedGame []selectGameStruct
	err = db.SupaClient.DB.From("games").Select("*").Eq("game_id_pk", gameID).Execute(&fetchedGame)
	if err != nil {
		log.Println("Error: HandleGetGame - Fetching from Games_T: ", err)
	}

	additionalData := map[string]string{
		"instructions":         "Whenever a user has made a move, refresh the page to get the move they made!",
		"how_to_play":          r.Host + `/how-to-play`,
		"Your_piece":           "White",
		"Welcome":              `Welcome, ` + user + ` to Backend Dam Haji AKA Backend Checkers!`,
		"make_your_first_move": r.Host + `/` + gameID + `/` + user + `/move/{origin}/to/{final}`,
	}

	if len(existingUser) > 0 {
		additionalData["Welcome"] = "Welcome Back " + user + ", lets get a W!"
	}

	p := startGamePayload{
		GameID:               gameID,
		BlackPlayer1Username: fetchedGame[0].BlackPlayer1Username,
		WhitePlayer2Username: user,
		BoardState:           common.InitBoardState,
		Data:                 additionalData,
	}

	utils.Serve(w, p)

}
