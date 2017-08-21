package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/rubenv/sql-migrate"
	"log"
	"os"
)

func main() {
	var err error
	dbStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", dbStr)
	if err != nil {
		log.Fatalf("Unable to connect to db: %s\n", err.Error())
	}

	migrations := &migrate.FileMigrationSource{
		Dir: "migrations/postgres",
	}

	n, err := migrate.Exec(db, "postgres", migrations, migrate.Up)
	if err != nil {
		log.Fatalf("Unable to apply migrations: %s\n", err.Error())
	}
	log.Printf("Applied %d migrations.\n", n)
}
