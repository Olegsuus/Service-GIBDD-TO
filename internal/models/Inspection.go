package models

import "time"

type Inspection struct {
	ID             int       `json:"id"`
	AutomobileID   int       `json:"automobile_id"`
	CardNumber     string    `json:"card_number"`
	InspectionDate time.Time `json:"inspection_date"`
	Notes          string    `json:"notes"`
}
