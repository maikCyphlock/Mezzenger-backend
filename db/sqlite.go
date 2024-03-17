package db

import (
	"database/sql"
	"log"

	_ "db/models"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		log.Fatal(err)
	}
	DB = db

	createTable()
}

type User struct {
	ID       int
	Username string
	Password string
}

// Error implements error.

func CreateUser(user *User) (u *User, err error) {
	res, err := DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, user.Password)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	user.ID = int(id)
	return user, nil

}

func GetUserByUsernameAndPassword(username, password string) (*User, error) {
	var user User
	err := DB.QueryRow("SELECT id, username, password FROM users WHERE username = ? AND password = ?", username, password).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func createTable() {
	_, err := DB.Exec(`
        CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT UNIQUE NOT NULL,
            password TEXT NOT NULL
        );
    `)
	if err != nil {
		log.Fatal(err)
	}
}
