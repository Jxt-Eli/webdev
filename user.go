package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	// "github.com/pelletier/go-toml/query"
	// "github.com/gorilla/mux"
)

var arr = make([]User, 0)

func (srv *Server) delUserHandler(w http.ResponseWriter, r *http.Request) {

	info := mux.Vars(r)
	email := info["email"]

	query := 
	`
		DELETE FROM users
		WHERE  email = $1
	`
	_, err := srv.DB.Exec(query, email)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Account deleted successfully!")
}

func (svr *Server) createUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	query :=
		`
			INSERT INTO users (username, email)
			VALUES ($1,$2)
			RETURNING id, created_at
		`
	err = svr.DB.QueryRow(query, newUser.Username, newUser.Email).Scan(&newUser.ID, &newUser.CreatedAt)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newUser)

	fmt.Fprintf(w, "account successfully created. thanks for working with us, %s\n", newUser.Username)

	arr = append(arr, newUser)
}

func (srv *Server) showUsers(w http.ResponseWriter, r *http.Request) {

	var user User
	resp := mux.Vars(r)
	userName := resp["email"]

	
	query := 
	`
		SELECT * FROM users
		WHERE email=$1
	`
	err := srv.DB.QueryRow(query, userName).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		fmt.Printf("Database Error: %s", err)
		return
	}
	w.Header().Set("Content-type", "application-json")
	json.NewEncoder(w).Encode(&user)
	
}
