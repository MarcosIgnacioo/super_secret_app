package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

func User(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Println(err)
	}

	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)

	if err != nil {
		log.Fatal(err)
	}

	statement, err := db.Prepare(`CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		name VARCHAR(64) NULL,
		password VARCHAR(64) NULL,
		birthday DATE NULL)`)

	if err != nil {
		log.Println("Error in creating table")
	} else {
		log.Println("Successfully created table persona 5!")
	}

	statement.Exec()

	// Create

	statement, _ = db.Prepare("INSERT INTO users (name, password, birthday) VALUES (?, ?, ?)")
	statement.Exec("Persona", "contra", "28-08-2024")
	log.Println("Inserted the person into database!")
	rows, _ := db.Query("SELECT id, name, password, birthday FROM users")

	var user Person
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Password, &user.Birthday)
	}

	fmt.Fprintf(w, "%s", user.Name)
}
