package models

import "time"

type Automobile struct {
	ID               int       `json:"id" db:"id"`
	ReleaseDate      time.Time `json:"release_date" db:"release_date"`
	Model            string    `json:"model" db:"model"`
	LicensePlate     string    `json:"license_plate" db:"license_plate"`
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
}

type Inspection struct {
	ID             int       `json:"id" db:"id"`
	AutomobileID   int       `json:"automobile_id" db:"automobile_id"`
	CardNumber     string    `json:"card_number" db:"card_number"`
	InspectionDate time.Time `json:"inspection_date" db:"inspection_date"`
	Notes          string    `json:"notes" db:"notes"`
}
