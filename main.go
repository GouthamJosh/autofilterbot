package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Jisin0/autofilterbot/cmd"
	"github.com/joho/godotenv"  // Optional: For .env support
)

func main() {
	// Load environment variables from .env file (optional)
	godotenv.Load()

	// Get port from environment variable, default to 8080 (Render sets this automatically)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Get webhook URL from env (set in Render for production)
	webhookURL := os.Getenv("WEBHOOK_URL")

	// Define the webhook handler (processes updates from Telegram)
	http.HandleFunc("/webhook", func(w http.ResponseWriter, r *http.Request) {
		var update map[string]interface{}  // Or use tgbotapi.Update if imported
		if err := json.NewDecoder(r.Body).Decode(&update); err != nil {
			log.Println("Webhook decode error:", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		// Process the update (pass to your bot logic in cmd package if needed)
		// For now, just log it; integrate with cmd.Execute() or bot handlers
		log.Println("Processed webhook update:", update["update_id"])
		w.WriteHeader(http.StatusOK)
	})

	// Start the HTTP server in the background (non-blocking, uses Render's PORT)
	go func() {
		log.Printf("Starting HTTP server on port %s", port)
		if err := http.ListenAndServe(":"+port, nil); err != nil {
			log.Fatal("Server error:", err)
		}
	}()

	// If webhook URL is set, assume webhooks are enabled (polling disabled in cmd)
	if webhookURL != "" {
		log.Println("Webhooks enabled for URL:", webhookURL)
	}

	// Execute the bot's command logic (from the cmd package)
	// Note: Ensure cmd.Execute() sets up the bot with webhooks if WEBHOOK_URL is provided
	cmd.Execute()
}
