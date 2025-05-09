package main

import (
	"fmt"
	"log"
	"main/src/db"
	"main/src/handlers"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/nedpals/supabase-go"
)

var supaClient *supabase.Client

func init() {
	supaClient = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"))
	if supaClient == nil {
		log.Fatal("Failed to initialize Supabase client")
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db.InitSupabase()
	r := mux.NewRouter()

	// API Routes
	r.HandleFunc("/", handlers.GetLeaderboard)
	r.HandleFunc("/leaderboard/{user}", handlers.GetLeaderboard)
	r.HandleFunc("/start-game/{user}", handlers.HandleInitGame)
	r.HandleFunc("/learn-checkers", handlers.LearnCheckers)
	r.HandleFunc("/how-to-use", handlers.HowToUse)
	r.HandleFunc("/history/{gameid}", handlers.GetCurrentGameHistory)
	r.HandleFunc("/set-game/{gameid}/{templatenum}", handlers.SetGameBoard)
	r.HandleFunc("/{gameid}/{user}/move/{start}/to/{end}", handlers.HandleGameMove)
	r.HandleFunc("/{gameid}/{user}", handlers.HandleGetPlayer2)
	r.HandleFunc("/{gameid}", handlers.HandleGetGame)

	// Initialize server
	port := ":8080"
	fmt.Println("Server running on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r))
}
