package models

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"string"`
	Email    string    `json:"email"`
	Password string    `json:"-"`
}
