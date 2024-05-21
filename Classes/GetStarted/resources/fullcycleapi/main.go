package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"` // or `json:"-"`
}

func main() {
	fmt.Println("FullCycle API v5.0")

	mux := http.NewServeMux()

	mux.HandleFunc("/users", listUsersHandler)
	mux.HandleFunc("POST /users", createUserHandler)

	http.ListenAndServe(":3000", mux)
}

func listUsersHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Get users...", r.Method)

	// It does a connection to DB
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	// Select data from DB
	rows, err := db.Query("Select * from users")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	// Create a slice of users
	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		users = append(users, u)
	}
	// Convert to JSON
	if err := json.NewEncoder(w).Encode(users); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Creating users...", r.Method)

	// It does a connection to DB
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, err := db.Exec("INSERT INTO users (id, name, email) VALUES (?, ?, ?)", u.ID, u.Name, u.Email); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
