package models

type ParkingRequest struct {
	Name string `json:"name" binding:"required"`
}

type Parking struct {
	Name       string `json:"name"`
	Status     string `json:"status"`
	FreeSpaces int    `json:"free_spaces"`
	Time       int64  `json:"time"`
}
