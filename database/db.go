package database

import (
	"database/sql"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

func NewUser(name string, password string, birthday string) *User {
	return &User{Name: name, Password: password, Birthday: birthday}
}

var db *sql.DB
var err error

func Init() {
	db, err = sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Println(err)
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
}
func End() {
	db.Close()
}

func Insert(user *User) error {
	statement, err := db.Prepare("INSERT INTO users (name, password, birthday) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}
	// dereferencia solito lol
	statement.Exec(user.Name, user.Password, user.Birthday)
	log.Println("Inserted the person into database!")
	return nil
}

func SelectLastUser() (*User, error) {
	rows, err := db.Query("SELECT id, name, password, birthday FROM users")
	if err != nil {
		return nil, err
	}

	var user User

	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Password, &user.Birthday)
	}

	return &user, nil
}
