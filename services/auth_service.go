package services

import (
	"cbt-backend/models"
	"cbt-backend/repositories"
	"cbt-backend/utils"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, username, password string) (string, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (s *authService) Register(ctx context.Context, user *models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user to database
	return s.userRepository.CreateUser(ctx, user)
}

func (s *authService) Login(ctx context.Context, username, password string) (string, error) {

	user, err := s.userRepository.GetUserByUsername(ctx, username)

	if err != nil {
		return "", errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", errors.New("invalid password")
	}

	token, err := utils.GenerateJWT(user.ID, user.Role)
	if err != nil {
		return "", err
	}

	return token, nil
}
