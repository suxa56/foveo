package auth

import (
	"context"
	"foveo/internal/domain/dto"
	"foveo/internal/domain/model"
	"foveo/internal/repository"
	"log/slog"
	"time"
)

type Auth struct {
	log          *slog.Logger
	userRepo     repository.IUserRepo
	userSaver    IUserSaver
	userProvider IUserProvider
	tokenTTL     time.Duration
}

type IUserSaver interface {
	SaveUser(ctx context.Context, request dto.RegisterRequest) (id int64, err error)
}

type IUserProvider interface {
	User(ctx context.Context, id int64) (user model.User, err error)
	IsAdmin(ctx context.Context, id int64) (isAdmin bool, err error)
}

func New(
	log *slog.Logger,
	userRepo repository.IUserRepo,
	userSaver IUserSaver,
	userProvider IUserProvider,
	tokenTTL time.Duration) *Auth {
	return &Auth{
		log:          log,
		userRepo:     userRepo,
		userSaver:    userSaver,
		userProvider: userProvider,
		tokenTTL:     tokenTTL,
	}
}
