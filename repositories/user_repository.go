package repositories

import (
	"cbt-backend/models"
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) error
	GetByUsername(email string) (*models.User, error)
	GetUserById(id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, updates map[string]any) error
	DeletUser(ctx context.Context, userID uuid.UUID) error
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

	row := r.db.QueryRowContext(ctx, "SELECT username FROM users WHERE username = $1", user.Username)

	err := row.Scan(&user.Username)

	if err == nil {
		return fmt.Errorf("user dengan username %s sudah terdaftar", user.Username)
	} else if err != sql.ErrNoRows {
		return err
	}

	_, err = r.db.ExecContext(
		ctx,
		"INSERT INTO users (username, name, email, password, address, grade, photo, gender, role_id, school_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		user.Username, user.Name, user.Email, user.Password, user.Address, user.Grade, user.Photo, user.Gender, user.RoleID, user.SchoolID,
	)

	return err
}

func (r *userRepository) GetByUsername(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, password FROM users WHERE username = $1"

	err := r.db.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetUserById(id uuid.UUID) (*models.User, error) {
	var user models.User

	query := `
		SELECT id, username, name, email, password, address, grade, photo, gender, 
				role_id, school_id, created_at, updated_at, deleted_at
		FROM users WHERE id = $1
	`

	err := r.db.QueryRow(query, id).Scan(
		&user.ID, &user.Username, &user.Name, &user.Email, &user.Password,
		&user.Address, &user.Grade, &user.Photo, &user.Gender,
		&user.RoleID, &user.SchoolID, &user.CreatedAt, &user.UpdatedAt, &user.DeletedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(ctx context.Context, userID uuid.UUID, updates map[string]any) error {
	query := "UPDATE users SET "
	args := []any{}
	i := 1

	for key, value := range updates {
		if i > 1 {
			query += ", "
		}
		query += key + " = $" + fmt.Sprintf("%d", i)
		args = append(args, value)
		i++
	}

	query += ", updated_at = NOW() WHERE id = $" + fmt.Sprintf("%d", i)
	args = append(args, userID)

	_, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (r *userRepository) DeletUser(ctx context.Context, userID uuid.UUID) error {

	query := `UPDATE users SET deleted_at = NOW() WHERE id = $1`

	_, err := r.db.ExecContext(ctx, query, userID)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil

}
