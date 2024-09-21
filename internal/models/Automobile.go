package models

import "time"

type Automobile struct {
	ID               int       `json:"id"`
	ReleaseDate      time.Time `json:"release_date"`
	Model            string    `json:"model"`
	LicensePlate     string    `json:"license_plate"`
	RegistrationDate time.Time `json:"registration_date"`
}
