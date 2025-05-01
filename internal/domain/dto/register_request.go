package dto

import (
	"fmt"
)

type RegisterRequest struct {
	FirstName          string `json:"first_name"`
	Surname            string `json:"surname"`
	DOB                string `json:"dob"`
	Address            string `json:"address"`
	Department         string `json:"department"`
	Position           string `json:"position"`
	Patronymic         string `json:"patronymic"`
	Email              string `json:"email"`
	Phone              string `json:"phone"`
	Gender             string `json:"gender"`
	Citizenship        string `json:"citizenship"`
	Birthplace         string `json:"birthplace"`
	Nationality        string `json:"nationality"`
	PassportNumber     string `json:"passport_number"`
	PassportIssuedBy   string `json:"passport_issued_by"`
	PassportIssueDate  string `json:"passport_issue_date"`
	PassportExpiryDate string `json:"passport_expiry_date"`
}

func (r *RegisterRequest) IsValid() []error {
	errors := make([]error, 0, 1)
	if r.FirstName == "" {
		errors = append(errors, fmt.Errorf("first_name is required"))
	}
	if r.Surname == "" {
		errors = append(errors, fmt.Errorf("surname is required"))
	}

	return errors
}
