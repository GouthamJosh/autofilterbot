package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Jisin0/autofilterbot/cmd"
	"github.com/joho/godotenv"  // Optional: For loading .env files
)

func main() {
	// Load environment variables from .env file (optional)
	godotenv.Load()

	// Get port from environment variable, default to 8080
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Define a simple webhook handler (customize for your bot's needs, e.g., Telegram webhook processing)
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		// Example: Log the request and respond OK
		// In a real implementation, parse r.Body for Telegram updates and process them
		log.Println("Webhook received on port", port)
		w.WriteHeader(http.StatusOK)
	})

	// Start the HTTP server in the background (non-blocking)
	go func() {
		log.Printf("Starting HTTP server on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Server error:", err)
		}
	}()

	// Execute the bot's command logic (from the cmd package)
	cmd.Execute()
}
