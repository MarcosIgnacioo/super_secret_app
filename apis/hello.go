package apis

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func GenericHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"esto": "functiona",
	})
}

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Birthday string `json:"birthday"`
}

func SQL(c *gin.Context) {
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

	var user User
	for rows.Next() {
		rows.Scan(&user.ID, &user.Name, &user.Password, &user.Birthday)
	}

	c.JSON(200, gin.H{
		"esto": user.Name,
	})
}
