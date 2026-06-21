package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "google.golang.org/genproto/googleapis/spanner/admin/database/v1"

	"github.com/Jxt-Eli/webdev/db"
	"github.com/Jxt-Eli/webdev/internal/handlers"
	"github.com/Jxt-Eli/webdev/internal/middleware"
)


func main() {
	database := db.Connect()
	defer database.Close()
	fmt.Println("postgres database connection successful")

	r := mux.NewRouter()
	r.Use(middleware.LoggingMiddleware, middleware.APIKeyMiddleware) // Attaches middleware to router to prevent boilerplate of wrapping each handler function in the middleware function

	srv := &handlers.Server{
		DB: database,
	}

	r.HandleFunc("/", handler).Methods("GET")
	r.HandleFunc("/users", srv.CreateUserHandler).Methods("POST")
	r.HandleFunc("/{email}", srv.ShowUsers).Methods("GET")
	r.HandleFunc("/{email}", srv.DelUserHandler).Methods("DELETE")

	port := ":8080"
	fmt.Printf("server running on port %s\n", port)
	http.ListenAndServe(port, r)
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n and this:\n %v", r.URL.Path, r)
}
