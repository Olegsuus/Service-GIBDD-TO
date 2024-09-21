package storage_models

import "time"

//type Inspection struct {
//	ID             int       `json:"id" db:"id"`
//	AutomobileID   int       `json:"automobile_id" db:"automobile_id"`
//	CardNumber     string    `json:"card_number" db:"card_number"`
//	InspectionDate time.Time `json:"inspection_date" db:"inspection_date"`
//	Notes          string    `json:"notes" db:"notes"`
//}

type Inspection struct {
	ID             int       `json:"id"`
	AutomobileID   int       `json:"automobile_id"`
	CardNumber     string    `json:"card_number"`
	InspectionDate time.Time `json:"inspection_date"`
	Notes          string    `json:"notes"`
}
