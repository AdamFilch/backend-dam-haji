package main

import (
	"fmt"
	"log"
	"net/http"
	"src/main/src/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r := mux.NewRouter()

	// API Routes
	r.HandleFunc("/leaderboard/{user}", handlers.GetLeaderboard)
	r.HandleFunc("/start-game/{user}", handlers.HandleInitGame)
	r.HandleFunc("/learn", handlers.LearnCheckers)
	r.HandleFunc("/how-to-use", handlers.HowToUse)
	r.HandleFunc("/history/{gameid}", handlers.GetCurrentGameHistory)
	r.HandleFunc("/{gameid}/{user}", handlers.HandleGetGame)
	r.HandleFunc("/{gameid}/{user}/move/{start}/to/{end}", handlers.HandleGameMove)

	// Initialize server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
