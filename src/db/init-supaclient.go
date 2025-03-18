package db

import (
	"log"
	"os"

	"github.com/nedpals/supabase-go"
)

// Global Supabase client
var SupaClient *supabase.Client

// Initialize Supabase client
func InitSupabase() {

	supabaseURL := os.Getenv("SUPABASE_URL")
	supabaseKey := os.Getenv("SUPABASE_API_KEY")

	if supabaseURL == "" || supabaseKey == "" {
		log.Fatal("SUPABASE_URL or SUPABASE_API_KEY is not set")
	}

	SupaClient = supabase.CreateClient(supabaseURL, supabaseKey)
	if SupaClient == nil {
		log.Fatal("Failed to initialize Supabase client")
	}
	
}
