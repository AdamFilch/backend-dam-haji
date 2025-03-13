package main

import (
	"fmt"
	"log"
	"net/http"
	"src/main/src/handlers"
)

func main() {

	http.HandleFunc("/leaderboard/{user}", handlers.GetLeaderboard)
	http.HandleFunc("/{gameid}/{user}", handlers.HandleGetGame)
	http.HandleFunc("/{gameid}/{user}/move/{start}/to/{end}", handlers.HandleGameMove)
	http.HandleFunc("/start-game/{user}", handlers.HandleInitGame)
	http.HandleFunc("/learn", handlers.LearnCheckers)
	http.HandleFunc("/how-to-use", handlers.HowToUse)
	http.HandleFunc("/history/{gameid}", handlers.GetCurrentGameHistory)

	// Initialize server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
