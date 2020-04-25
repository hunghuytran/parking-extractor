package main

import (
	"challenge/extractor/services"
)

func main() {
	err := services.ExtractParkingData()
	if err != nil {
		panic("Error: unable to extract data")
	}
}
