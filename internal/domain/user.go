package domain

import "time"

type User struct {
	ID                 int
	FirstName          string
	Surname            string
	Patronymic         string
	DOB                time.Time
	Email              string
	Phone              string
	Address            string
	Gender             string
	DepartmentID       *int
	PositionID         *int
	IsEmployee         bool
	Status             string
	Citizenship        string
	Birthplace         string
	Nationality        string
	PassportNumber     string
	PassportIssuedBy   string
	PassportIssueDate  time.Time
	PassportExpiryDate time.Time
	Roles              []Role
}
