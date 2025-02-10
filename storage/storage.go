package storage

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

type DataStore interface {
	Ping() error
	Close() error
}

type PostgresStore struct {
	DB *sql.DB
	l  *log.Logger
}

func (p PostgresStore) Ping() error {
	return p.DB.Ping()
}

func (p PostgresStore) Close() error {
	return p.DB.Close()
}

func NewPostgresStore(l *log.Logger) PostgresStore {
	connString, found := os.LookupEnv("CONN_STR")
	if !found {
		l.Fatal("DBCONN_STR not found in environment!")
	}

	db, err := sql.Open("postgres", connString)
	if err != nil {
		l.Fatalf("Unable to connect to the database server! Error: %v", err)
	}

	if err := db.Ping(); err != nil {
		l.Fatalf("Unable to communicate with the database server! Error: %v", err)
	}

	l.Print("Connected with postgres server succesfully!")
	return PostgresStore{
		DB: db,
		l:  l,
	}
}
