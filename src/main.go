package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"src/main/src/handlers"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	supabase "github.com/lengzuo/supa"
)

var supaClient *supabase.Client

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	conf := supabase.Config{
		ApiKey:     os.Getenv("SUPABASE_API_KEY"),
		ProjectRef: os.Getenv("SUPABASE_URL"),
		Debug:      true,
	}
	supaClient, err = supabase.New(conf)
	if err != nil {
		fmt.Println("Failed in init of supa client: ", err)
		return
	}

	r := mux.NewRouter()

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
