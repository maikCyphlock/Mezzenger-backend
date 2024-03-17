package models

import (
	"fmt"
)

type User struct {
	ID       int
	Username string
	Password string
}

func (u *User) String() string {
	return fmt.Sprintf("ID: %d, Username: %s, Password: %s", u.ID, u.Username, u.Password)
}
