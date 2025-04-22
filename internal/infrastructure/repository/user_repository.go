package repository

import (
	"context"
	"geoservise-jwt/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, user model.User) (int64, error)
	GetByUsername(ctx context.Context, username string) (model.User, error)
}
