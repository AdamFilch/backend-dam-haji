package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	fmt.Println("Vars:", vars)
	log.Println("Received user:", p)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(p)
}

func HandleGetGame(w http.ResponseWriter, r *http.Request) {

}
