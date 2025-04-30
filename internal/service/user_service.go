package service

import (
	"context"
	"foveo/internal/domain/model"
	"foveo/internal/repository"
)

type IUserService interface {
	Register(ctx context.Context, user *model.User) error
}

type UserService struct {
	userRepo repository.IUserRepo
}

func InitUserService(userRepo repository.IUserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (u *UserService) Register(ctx context.Context, user *model.User) error {
	return u.Register(ctx, user)
}
