package models

type UpdateRes struct {
	Success bool `json:"success"`
}

type DeleteRes struct {
	Success bool `json:"success"`
}

type AddRes struct {
	ID int `json:"id"`
}
