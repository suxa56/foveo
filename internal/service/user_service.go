package service

import (
	"context"
	"foveo/internal/domain/dto"
	"foveo/internal/domain/model"
	"foveo/internal/repository"
	"time"
)

type IUserService interface {
	Register(ctx context.Context, requestDto dto.RegisterRequest) error
	GetList(ctx context.Context) ([]model.User, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
}

type UserService struct {
	userRepo repository.IUserRepo
}

func InitUserService(userRepo repository.IUserRepo) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) Register(ctx context.Context, requestDto dto.RegisterRequest) error {
	dob, _ := time.Parse(time.DateOnly, requestDto.DOB)
	passportIssue, _ := time.Parse(time.DateOnly, requestDto.PassportIssueDate)
	passportExp, _ := time.Parse(time.DateOnly, requestDto.PassportExpiryDate)
	user := &model.User{
		ID:                 0,
		FirstName:          requestDto.FirstName,
		Surname:            requestDto.Surname,
		Patronymic:         requestDto.Patronymic,
		DOB:                dob,
		Email:              requestDto.Email,
		Phone:              requestDto.Phone,
		Address:            requestDto.Address,
		Gender:             requestDto.Gender,
		DepartmentID:       nil,
		PositionID:         nil,
		IsEmployee:         false,
		Status:             "pending",
		Citizenship:        requestDto.Citizenship,
		Birthplace:         requestDto.Birthplace,
		Nationality:        requestDto.Nationality,
		PassportNumber:     requestDto.PassportNumber,
		PassportIssuedBy:   requestDto.PassportIssuedBy,
		PassportIssueDate:  passportIssue,
		PassportExpiryDate: passportExp,
		Roles:              nil,
	}
	return s.userRepo.Create(ctx, user)
}

func (s *UserService) GetList(ctx context.Context) ([]model.User, error) {
	return s.userRepo.GetList(ctx)
}

func (s *UserService) GetByID(ctx context.Context, id int) (*model.User, error) {
	return s.userRepo.GetByID(ctx, id)
}
