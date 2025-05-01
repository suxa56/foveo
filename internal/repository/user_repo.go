package repository

import (
	"context"
	"fmt"
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
	query := `INSERT INTO users (
		first_name, surname, patronymic, dob, email, phone, address, gender, 
		department_id, position_id, is_employee, status, citizenship, birthplace, 
		nationality, passport_number, passport_issued_by, passport_issue_date, passport_expiry_date
	) VALUES (
		$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19
	) RETURNING id`

	err := r.db.QueryRowContext(ctx, query,
		user.FirstName,
		user.Surname,
		user.Patronymic,
		user.DOB,
		user.Email,
		user.Phone,
		user.Address,
		user.Gender,
		user.DepartmentID,
		user.PositionID,
		user.IsEmployee,
		user.Status,
		user.Citizenship,
		user.Birthplace,
		user.Nationality,
		user.PassportNumber,
		user.PassportIssuedBy,
		user.PassportIssueDate,
		user.PassportExpiryDate,
	).Scan(&user.ID)

	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}
