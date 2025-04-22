package service

import (
	"context"
	"geoservise-jwt/internal/infrastructure/repository"
	"geoservise-jwt/internal/model"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Create(ctx context.Context, user model.User) (int64, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) Create(ctx context.Context, user model.User) (int64, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return 0, err
	}
	user.Password = string(hashedPassword)
	return s.repo.Create(ctx, user)
}

func (s *userService) GetByUsername(ctx context.Context, username string) (model.User, error) {
	return s.repo.GetByUsername(ctx, username)
}
