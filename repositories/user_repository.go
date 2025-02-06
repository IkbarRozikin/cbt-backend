package repositories

import (
	"cbt-backend/models"
	"context"
	"database/sql"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)
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
		"INSERT INTO users (username, password, role) VALUES ($1, $2, $3)",
		user.Username, user.Password, user.Role,
	)
	return err
}

func (r *userRepository) GetUserByUsername(ctx context.Context, username string) (*models.User, error) {
	var user models.User
	err := r.db.QueryRowContext(
		ctx,
		"SELECT id, username, password, role FROM users WHERE username=$1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password, &user.Role)

	if err != nil {
		return nil, err
	}
	return &user, nil
}
