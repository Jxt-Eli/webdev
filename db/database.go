package db

import(
	// "fmt"
	"log"
	"database/sql"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Connect() *sql.DB{
	connstr := "host=/run/postgresql dbname=webdev port=5432 user=jxt-eli password=password1937 sslmode=disable"
	
	db, err := sql.Open("pgx", connstr)
	if err != nil{
		log.Fatal("error loading credentials")
	}

	err = db.Ping()
	if err != nil{
		log.Fatal("failed to establish database connection")
	}
	// fmt.Printf("\ndb struct ptr: %v\n", db)

	return db
}
