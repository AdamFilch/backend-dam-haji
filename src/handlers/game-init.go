package handlers

import (
	"log"
	"net/http"
	"os"
	"main/src/utils"
	"time"

	"github.com/gorilla/mux"
	"github.com/nedpals/supabase-go"
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
	GameID          string              `json:"game_id_pk"`
	Player1Username string              `json:"player1_username"`
	BoardState      map[string][]string `json:"board_state"`
	Status          string              `json:"status"`
	CreatedAt       time.Time           `json:"created_at"`
	UpdatedAt       time.Time           `json:"updated_at"`
}

func HandleInitGame(w http.ResponseWriter, r *http.Request) {

	supaClient := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"))
	if supaClient == nil {
		log.Fatal("Failed to initialize Supabase client")
	}

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
	err = supaClient.DB.From("users").Insert(newUser).Execute(&res)
	if err != nil {
		log.Fatal("An error has been encountered trying to insert to Users_T: ", err)
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

	err = supaClient.DB.From("games").Insert(newGame).Execute(&res)
	if err != nil {
		log.Fatal("An error has been encountered trying to insert to Games_t: ", err)
	}

	var insertedGames []any
	err = supaClient.DB.From("games").Select("*").Eq("game_id_pk", generatedGameID).Execute(&insertedGames)
	if err != nil {
		log.Println("An error has been encountered trying to fetching from games table: ", err)
	}

	log.Println("Fetched Inserted Games", insertedGames)

	p := initGamePayload{
		GameID:          generatedGameID,
		Username:        user,
		Player2GameLink: "test",
	}
	utils.Serve(w, p)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {
	log.Println("Log Handle Get Game")
}
