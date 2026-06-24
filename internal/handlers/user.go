package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"database/sql"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"

	"github.com/Jxt-Eli/webdev/internal/models"
)

type Server struct {
	DB *sql.DB
}

func (srv *Server) DelUserHandler(w http.ResponseWriter, r *http.Request) {

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

func (svr *Server) CreateUserHandler(w http.ResponseWriter, r *http.Request) {

	var newUser models.User
	ctx := r.Context()

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		return
	}
	plainTextPassword := []byte(newUser.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(plainTextPassword, bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("error: %s", err)
		http.Error(w, "Bad Request", http.StatusInternalServerError)
		return
	}
	// TODO: Remove manual dereference
	(&newUser).Password = string(hashedPassword)

	query :=
		`
			INSERT INTO users (username, email, password)
			VALUES ($1,$2, $3)
			RETURNING id, created_at
		`
// TODO: switch to sqlx
	err = svr.DB.QueryRowContext(ctx, query, newUser.Username, newUser.Email, newUser.Password).Scan(&newUser.ID, &newUser.CreatedAt)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		fmt.Printf("error: %s", err)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&newUser)

	fmt.Fprintf(w, "account successfully created. thanks for working with us, %s\n", newUser.Username)

}

func (srv *Server) ShowUsers(w http.ResponseWriter, r *http.Request) {

	var user models.User
	resp := mux.Vars(r)
	userName := resp["email"]

	query :=
		`
		SELECT * FROM users
		WHERE email=$1
	`
// TODO: switch to sqlx
	err := srv.DB.QueryRow(query, userName).Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		http.Error(w, "400 Bad Request", http.StatusBadRequest)
		fmt.Printf("Database Error: %s", err)
		return
	}
	w.Header().Set("Content-type", "application-json")
	json.NewEncoder(w).Encode(&user)

}
