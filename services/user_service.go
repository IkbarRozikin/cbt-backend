package services

import (
	"cbt-backend/models"
	"cbt-backend/repositories"
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
)

type UserService interface {
	GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error)
	UpdateUser(ctx context.Context, userID uuid.UUID, input map[string]any) error
	DeleteUser(ctx context.Context, userID uuid.UUID) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{
		userRepository: userRepo,
	}
}

func (s *userService) GetUserById(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.userRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userService) UpdateUser(ctx context.Context, userID uuid.UUID, input map[string]any) error {
	// Validasi jika perlu, misalnya pastikan input tidak kosong
	if len(input) == 0 {
		return errors.New("no fields to update")
	}

	// Panggil repository untuk update data
	err := s.userRepository.UpdateUser(ctx, userID, input)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}

func (s *userService) DeleteUser(ctx context.Context, userID uuid.UUID) error {

	err := s.userRepository.DeletUser(ctx, userID)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
