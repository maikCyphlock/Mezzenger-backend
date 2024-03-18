package db

import (
	"database/sql"
	_ "db/models"
	"log"

	"golang.org/x/crypto/bcrypt"

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

// GenerateHashedPassword generates a hashed password from a plaintext password.
// The returned string is the hashed password.
// The error is not nil if there was an error hashing the password.
func GenerateHashedPassword(password string) (string, error) {
	const minCost = 10

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), minCost)

	return string(hashedPassword), err
}

func CreateUser(user *User) (u *User, err error) {
	hashedPassword, err := GenerateHashedPassword(user.Password)
	res, err := DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, hashedPassword)
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

func GetUserByCredentials(username, password string) (*User, error) {
	var user User
	var hashedPassword string
	var err error

	err = DB.QueryRow("SELECT id, username, password FROM users WHERE username = ?", username).Scan(&user.ID, &user.Username, &hashedPassword)

	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
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
