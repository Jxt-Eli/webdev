package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// "google.golang.org/genproto/googleapis/spanner/admin/database/v1"

	"github.com/Jxt-Eli/webdev/db"
)

// INFO: public struct
type User struct{
	ID        string      `json:"id"`
	Username  string      `json:"username"`
	Email     string      `json:"email"`
	CreatedAt time.Time   `json:"created_at"`

	Password	string `json:"-"`
}

type Server struct{
	DB *sql.DB
}


func main() {
	database := db.Connect()
	defer database.Close()
	fmt.Println("postgres database connection successful")

	r := mux.NewRouter()
	r.Use(middleware) // Attaches middleware to router to prevent boilerplate of wrapping each handler function in the middleware function 

	srv := &Server{
		DB: database,
	}


	
	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/users", srv.createUserHandler).Methods("POST")
	r.HandleFunc("/{email}", srv.showUsers).Methods("GET")
	r.HandleFunc("/{email}", srv.delUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{user_id}/repos/{repo}", userHandler).Methods("GET")


	port := ":8080"
	fmt.Printf("server running on port:%s\n", port)
	http.ListenAndServe(port, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n and this:\n %v", r.URL.Path, r)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	info := mux.Vars(r)
	user_id := info["user_id"]
	repo := info["repo"]
	fmt.Fprintf(w, "your user id is: %s and you requested the repo titled: %s\n", user_id, repo)
}

