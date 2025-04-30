package repository

import (
	"context"
	"foveo/internal/domain/model"
	"github.com/jmoiron/sqlx"
)

type IUserRepo interface {
	Create(ctx context.Context, user *model.User) error
}

type UserRepo struct {
	db *sqlx.DB
}

func InitUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (first_name, surname, ...) VALUES ($1, $2, ...) RETURNING id`
	return r.db.QueryRowContext(ctx, query, user.FirstName, user.Surname /* ... */).Scan(user.ID)
}
