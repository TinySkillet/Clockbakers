package main

import (
	"database/sql"
	"log"
	"os"

	"embed"

	h "github.com/TinySkillet/ClockBakers/handlers"
	"github.com/TinySkillet/ClockBakers/storage"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
)

//go:embed sql/schema/*.sql
var embedMigrations embed.FS

func main() {
	godotenv.Load()
	logger := log.New(os.Stdout, "cbaker-api: ", log.LstdFlags)

	portString, found := os.LookupEnv("PORT")
	if !found {
		logger.Fatal("No PORT variable found in the environment!")
	}

	connStr := os.Getenv("CONN_STR")
	if connStr == "" {
		logger.Fatal("No database connection string found in the environment!")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		logger.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// set embedded files for goose
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("postgres"); err != nil {
		logger.Fatalf("Error setting dialect: %v", err)
	}

	// apply migrations
	if err := goose.Up(db, "sql/schema"); err != nil {
		logger.Fatalf("Error applying migrations: %v", err)
	}

	store := storage.NewPostgresStore(logger)
	server := h.NewAPIServer(portString, logger, store)

	server.Run()
}
