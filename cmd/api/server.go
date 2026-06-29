package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"

	// "google.golang.org/genproto/googleapis/spanner/admin/database/v1"

	"github.com/Jxt-Eli/webdev/internal/db"
	"github.com/Jxt-Eli/webdev/internal/handlers"
	"github.com/Jxt-Eli/webdev/internal/middleware"
)

func main() {
	database := db.Connect()
	defer database.Close()
	fmt.Println("postgres database connection successful")

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware, middleware.APIKeyMiddleware) // Does the dirty function nesting work.

	srv := &handlers.Server{
		DB: database,
	}

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/users", srv.CreateUserHandler).Methods("POST")
	r.HandleFunc("/{email}", srv.ShowUsers).Methods("GET")
	r.HandleFunc("/{email}", srv.DelUserHandler).Methods("DELETE")

	if err := godotenv.Load(); err != nil {
		fmt.Printf("%s\nDefaulting to prod env vars...", err)
	}
	serverPort := os.Getenv("PORT")
	fmt.Printf("server running on port %s\n", serverPort)	//TODO: Remove
	http.ListenAndServe(serverPort, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to my first Go backend</h1><h3>I just realized that you can literally put html in here lol</h3>\nYou requested: %s\n This is your request struct btw:\n %v", r.URL.Path, r)
}
