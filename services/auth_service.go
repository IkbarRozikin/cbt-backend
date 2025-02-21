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
	RegisterService(ctx context.Context, user *models.User) error
	LoginService(ctx context.Context, user *models.Login) (string, error)
}

type authService struct {
	userRepository repositories.UserRepository
}

func NewAuthService(userRepo repositories.UserRepository) AuthService {
	return &authService{
		userRepository: userRepo,
	}
}

func (s *authService) RegisterService(ctx context.Context, user *models.User) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Save user to database
	return s.userRepository.CreateUser(ctx, user)
}

func (s *authService) LoginService(ctx context.Context, user *models.Login) (string, error) {

	storedUser, err := s.userRepository.GetByUsername(user.Username)

	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(storedUser.Password), []byte(user.Password)); err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(storedUser.ID, storedUser.RoleID, storedUser.SchoolID)
	if err != nil {
		return "", errors.New("could not generate token")
	}

	return token, nil
}
