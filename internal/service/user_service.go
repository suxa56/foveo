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
