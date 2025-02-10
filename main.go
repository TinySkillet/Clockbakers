package main

import (
	"log"
	"os"

	h "github.com/TinySkillet/ClockBakers/handlers"
	"github.com/TinySkillet/ClockBakers/storage"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	logger := log.New(os.Stdout, "cbaker-api: ", log.LstdFlags)

	portString, found := os.LookupEnv("PORT")
	if !found {
		logger.Fatal("No PORT variable found in the environment!")
	}

	store := storage.NewPostgresStore(logger)

	server := h.NewAPIServer(portString, logger, store)
	server.Run()
}
