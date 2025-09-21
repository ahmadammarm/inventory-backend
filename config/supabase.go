package config

import (
	"log/slog"
	"os"
	"sync"

	storage_go "github.com/supabase-community/storage-go"
)

var (
	SupabaseClient *storage_go.Client
	Once   sync.Once
)

func SupabaseStorageConnect() {
	Once.Do(func() {

		url := os.Getenv("SUPABASE_URL")
		apiKey := os.Getenv("SUPABASE_API_KEY")

		if url == "" || apiKey == "" {
			slog.Error("SUPABASE_URL or SUPABASE_API_KEY environment variable is not set")
		}

		client := storage_go.NewClient(url, apiKey, nil)

		SupabaseClient = client
	})
}
