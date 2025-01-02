package domain

import "github.com/google/uuid"

// User telegram user
type User struct {
	ID       uuid.UUID
	UserName string
}

func NewUser(userName string) *User {
	return &User{
		ID:       uuid.New(),
		UserName: userName,
	}
}
