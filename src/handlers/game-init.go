package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"src/main/src/utils"
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

func HandleInitGame(w http.ResponseWriter, r *http.Request) {

	supaClient := supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"))

	vars := mux.Vars(r)
	user := vars["user"]

	newUser := newUserStruct{
		Username:    user,
		TotalPoints: 0,
		GamesWon:    0,
		CreatedAt:   time.Now(),
	}

	var res any
	err := supaClient.DB.From("users_t").Insert(newUser).Execute(&res)
	if err != nil {
		log.Fatal("An error has been encountered trying to insert to Users_T")
	}

	fmt.Println("InitResult", res)

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
