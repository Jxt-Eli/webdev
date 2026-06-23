package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

// TODO: switch to sqlx
func Connect() *sql.DB {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if err := godotenv.Load(); err != nil {
		fmt.Println("local env not found. Defaulting to prod env vars")
	}
	connstr := os.Getenv("CONN_STR")

	db, err := sql.Open("pgx", connstr)
	if err != nil {
		log.Fatal("error loading credentials")
	}

	if err = db.PingContext(ctx); err != nil {
		log.Fatal("failed to establish database connection")
	}
	// fmt.Printf("\ndb struct ptr: %v\n", db)

	return db
}
