package cmd

import (
	"flag"
	"os"

	"github.com/Jisin0/autofilterbot/internal/core"
)

func Execute() {
	mongodbUri := flag.String("mongodb-uri", "", "mongodb uri for database")
	botToken := flag.String("bot-token", "", "bot token obtained from @botfather")
	logLevel := flag.String("log-level", "info", "level of logs to be shown")
	noOutput := flag.Bool("no-output", false, "disable console logs")
	port := flag.String("port", "", "port to bind HTTP server")

	flag.Parse()

	// Render injects PORT
	if *port == "" {
		*port = os.Getenv("PORT")
	}
	if *port == "" {
		*port = "8080" // local default
	}

	core.Run(core.RunAppOptions{
		MongodbURI:         *mongodbUri,
		BotToken:           *botToken,
		LogLevel:           *logLevel,
		DisableConsoleLogs: *noOutput,
		Port:               *port,
	})
}
