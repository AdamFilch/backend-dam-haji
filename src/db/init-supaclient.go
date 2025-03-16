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
	SupaClient = supabase.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_API_KEY"))
	if SupaClient == nil {
		log.Fatal("Failed to initialize Supabase client")
	}
}
