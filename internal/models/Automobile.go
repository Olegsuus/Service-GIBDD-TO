package models

import "time"

type Automobile struct {
	ID               int       `json:"id" db:"id"`
	ReleaseDate      time.Time `json:"release_date" db:"release_date"`
	Model            string    `json:"model" db:"model"`
	LicensePlate     string    `json:"license_plate" db:"license_plate"` //Номерной знак
	RegistrationDate time.Time `json:"registration_date" db:"registration_date"`
}
