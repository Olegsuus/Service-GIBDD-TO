package models

type Inspection struct {
	ID             int        `json:"id"`
	AutomobileID   int        `json:"automobile_id" validate:"required"`
	CardNumber     string     `json:"card_number" validate:"required,min=5,max=20"`
	InspectionDate CustomDate `json:"inspection_date" validate:"required"`
	Notes          string     `json:"notes"`
}

type AddInspectionDTO struct {
	AutomobileID   int        `json:"automobile_id" validate:"required"`
	CardNumber     string     `json:"card_number" validate:"required,min=5,max=20"`
	InspectionDate CustomDate `json:"inspection_date" validate:"required"`
	Notes          string     `json:"notes"`
}

type UpdateInspectionDTO struct {
	AutomobileID   int        `json:"automobile_id" validate:"required"`
	CardNumber     string     `json:"card_number" validate:"required,min=5,max=20"`
	InspectionDate CustomDate `json:"inspection_date" validate:"required"`
	Notes          string     `json:"notes"`
}
