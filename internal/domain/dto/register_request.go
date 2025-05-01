package dto

import (
	"fmt"
)

type RegisterRequest struct {
	FirstName  string `json:"first_name"`
	Surname    string `json:"surname"`
	DOB        string `json:"dob"`
	Address    string `json:"address"`
	Department string `json:"department"`
	Position   string `json:"position"`
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
