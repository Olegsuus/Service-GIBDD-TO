package models

type Automobile struct {
	ID               int        `json:"id"`
	ReleaseDate      CustomDate `json:"release_date" validate:"required"`
	Model            string     `json:"model" validate:"required,min=2,max=100"`
	LicensePlate     string     `json:"license_plate" validate:"required,license_plate,min=5,max=10"`
	RegistrationDate CustomDate `json:"registration_date" validate:"required"`
}

type AddAutomobileDTO struct {
	ReleaseDate      CustomDate `json:"release_date" validate:"required"`
	Model            string     `json:"model" validate:"required,min=2,max=100"`
	LicensePlate     string     `json:"license_plate" validate:"required,license_plate,min=5,max=10"`
	RegistrationDate CustomDate `json:"registration_date" validate:"required"`
}

type UpdateAutomobileDTO struct {
	ReleaseDate      CustomDate `json:"release_date" validate:"required"`
	Model            string     `json:"model" validate:"required,min=2,max=100"`
	LicensePlate     string     `json:"license_plate" validate:"required,license_plate,min=5,max=10"`
	RegistrationDate CustomDate `json:"registration_date" validate:"required"`
}
