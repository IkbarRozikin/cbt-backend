package repositories

import (
	"cbt-backend/models"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{
		db,
	}
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) error {
	_, err := r.db.ExecContext(
		ctx,
		"INSERT INTO users (username, password, role_id) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Role,
	)
	return err
}
