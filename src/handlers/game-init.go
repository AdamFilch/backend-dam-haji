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
	GameID               string              `json:"game_id_pk"`
	BlackPlayer1Username string              `json:"black_player1_username"`
	BoardState           map[string][]string `json:"board_state"`
	Status               string              `json:"status"`
	CreatedAt            time.Time           `json:"created_at"`
	UpdatedAt            time.Time           `json:"updated_at"`
}

func HandleInitGame(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["user"]

	newUser := newUserStruct{
		Username:    user,
		TotalPoints: 0,
		GamesWon:    0,
		CreatedAt:   time.Now().UTC(),
	}

	var res any
	var err error

	var existingUser []newUserStruct
	err = db.SupaClient.DB.From("users").Select("*").Eq("username", user).Execute(&existingUser)
	if err != nil {
		log.Println("Error: HandleInitGame - Fetching from users_t: ", err)
	}
	if len(existingUser) == 0 {
		err = db.SupaClient.DB.From("users").Insert(newUser).Execute(&res)
		if err != nil {
			log.Println("Error: HandleInitGame - Inserting to users_t: ", err)
		}
	}


	generatedGameID := utils.CreateNanoID()

	newGame := newGameStruct{
		GameID:               generatedGameID,
		BlackPlayer1Username: user,
		BoardState:           common.InitBoardState,
		Status:               "ongoing",
		CreatedAt:            time.Now().UTC(),
		UpdatedAt:            time.Now().UTC(),
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


type startGamePayload struct {
	GameID               string              `json:"gameId"`
	BlackPlayer1Username string              `json:"black_player1_username"`
	WhitePlayer2Username string              `json:"white_player2_username"`
	BoardState           map[string][]string `json:"board_state"`
	UpdatedAt            string              `json:"last_updated_at"`
	Data                 map[string]string   `json:"data"`
}

type updateGameStruct struct {
	WhitePlayer2Username string    `json:"white_player2_username"`
	UpdatedAt            time.Time `json:"updated_at"`
}

type alreadyMatchedPayload struct {
	GameID               string            `json:"gameId"`
	WhitePlayer2Username string            `json:"white_player2_username"`
	BlackPlayer2Username string            `json:"black_player2_username"`
	Data                 map[string]string `json:"data"`
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["user"]
	gameID := vars["gameid"]

	var res any
	var err error
	
	var fetchedGame []common.TableGameStruct
	err = db.SupaClient.DB.From("games").Select("*").Eq("game_id_pk", gameID).Execute(&fetchedGame)
	if err != nil {
		log.Println("Error: HandleGetGame - Fetching from Games_T: ", err)
	}
	if len(fetchedGame) == 0 {
		additionalData := map[string]string {
			"error": "Oops!",
			"message": "It seems like this game does not exists within our system!",
			"start-game": "Please create a new game by using: " + r.Host + `/start-game/` + user,
		}

		utils.Serve(w, additionalData)
		return
	}

	// If game has 2 players already
	if fetchedGame[0].BlackPlayer1Username != "" && fetchedGame[0].WhitePlayer2Username != "" {
		if fetchedGame[0].BlackPlayer1Username != user && fetchedGame[0].WhitePlayer2Username != user {
			additionalData := map[string]string{
				"start-game": "Start your own game by using: " + r.Host + `/start-game/` + user,
				"error":      "Unfortunately this game already has 2 players playing, " + fetchedGame[0].BlackPlayer1Username + " and " + fetchedGame[0].WhitePlayer2Username,
			}

			p := alreadyMatchedPayload{
				GameID:               gameID,
				BlackPlayer2Username: fetchedGame[0].BlackPlayer1Username,
				WhitePlayer2Username: fetchedGame[0].WhitePlayer2Username,
				Data:                 additionalData,
			}
			utils.Serve(w, p)
			return
		}
	}


	newUser := newUserStruct{
		Username:    user,
		TotalPoints: 0,
		GamesWon:    0,
		CreatedAt:   time.Now().UTC(),
	}

	var existingUser []newUserStruct
	err = db.SupaClient.DB.From("users").Select("*").Eq("username", user).Execute(&existingUser)
	if err != nil {
		log.Println("Error: HandleInitGame - Fetching from users_t: ", err)
	}
	if len(existingUser) == 0 {
		err = db.SupaClient.DB.From("users").Insert(newUser).Execute(&res)
		if err != nil {
			log.Println("Error: HandleInitGame - Inserting to users_t: ", err)
		}
	}


	updatedGame := updateGameStruct{
		WhitePlayer2Username: user,
		UpdatedAt:            time.Now().UTC(),
	}

	err = db.SupaClient.DB.From("games").Update(updatedGame).Eq("game_id_pk", gameID).Execute(&res)
	if err != nil {
		log.Println("Error: HandleGameGetGame - Updating from games_t", err)
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
