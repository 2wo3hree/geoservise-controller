package postgres

import (
	"context"
	"geoservise-jwt/internal/infrastructure/repository"
	"geoservise-jwt/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type userRepo struct {
	db *pgxpool.Pool
}

func NewUserRepo(db *pgxpool.Pool) repository.UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) Create(ctx context.Context, user model.User) (int64, error) {
	query := `INSERT INTO users (username, password)
			  VALUES ($1,$2) RETURNING id`
	var id int64
	err := r.db.QueryRow(ctx, query, user.Username, user.Password).Scan(&id)
	return id, err
}

func (r *userRepo) GetByUsername(ctx context.Context, username string) (model.User, error) {
	query := `SELECT id, username,password FROM users WHERE username=$1`
	row := r.db.QueryRow(ctx, query, username)

	var user model.User
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	return user, err
}
