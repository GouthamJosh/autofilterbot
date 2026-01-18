package core

import (
	"log"
	"net/http"
)

type RunAppOptions struct {
	MongodbURI         string
	BotToken           string
	LogLevel           string
	DisableConsoleLogs bool
	Port               string
}

func Run(opts RunAppOptions) {
	// --- init db, logger, bot, etc ---
	log.Println("Starting application")
	log.Println("Log level:", opts.LogLevel)

	// Health endpoint (recommended for Render)
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	addr := "0.0.0.0:" + opts.Port
	log.Println("HTTP server listening on", addr)

	// Start HTTP server (REQUIRED for Render)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
