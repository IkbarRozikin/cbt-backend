package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username" validate:"required,usernameRegexp,min=3,max=25"`
	Password  string    `json:"password" validate:"required,min=3,max=100"`
	Role      string    `json:"role_id" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
