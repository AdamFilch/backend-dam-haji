package handlers

import (
	"log"
	"net/http"
	"src/main/src/utils"

	"github.com/gorilla/mux"
)

type initGamePayload struct {
	GameID          string `json:"gameId"`
	Username        string `json:"username"`
	Player2GameLink string `json:"player2_gameLink"`
}

func HandleInitGame(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	user := vars["user"]

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
