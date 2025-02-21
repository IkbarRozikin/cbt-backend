package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	Username  string     `json:"username" validate:"required,usernameRegexp,min=3,max=25"`
	Name      string     `json:"name" validate:"required"`
	Email     string     `json:"email"`
	Password  string     `json:"password" validate:"required,min=3,max=100"`
	Address   string     `json:"address"`
	Grade     int        `json:"grade"`
	Photo     string     `json:"photo"`
	Gender    string     `json:"gender"`
	RoleID    uuid.UUID  `json:"role_id" validate:"required"`
	SchoolID  uuid.UUID  `json:"school_id" validate:"required"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
