package service

import (
	"context"
	"foveo/internal/domain/dto"
	"foveo/internal/repository"
)

type IUserService interface {
	Register(ctx context.Context, requestDto dto.RegisterRequest) error
}

type UserService struct {
	userRepo repository.IUserRepo
}

func InitUserService(userRepo repository.IUserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(ctx context.Context, requestDto dto.RegisterRequest) error {
	//return s.userRepo.Create(ctx, user)
	return nil
}
