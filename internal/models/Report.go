package models

type Report struct {
	TotalCars       int `json:"total_cars"`
	CarsOlder3Years int `json:"cars_older_3_years"`
	CarsNewer3Years int `json:"cars_newer_3_years"`
}
