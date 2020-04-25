package services

import (
	"challenge/models"
	"encoding/json"
	"github.com/go-errors/errors"
	"io/ioutil"
	"strings"
)

func GetParking(query string) (*models.Parking, error) {
	var parkings []*models.Parking

	value, err := ioutil.ReadFile("parking.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(value, &parkings)
	if err != nil {
		return nil, err
	}

	var found *models.Parking
	for _, parking := range parkings {
		if strings.Contains(strings.ToLower(parking.Name), strings.ToLower(query)) {
			found = parking
			break
		}
	}

	if found == nil {
		return nil, errors.New(errors.Errorf("%s", "Parking not found"))
	}

	return found, nil
}
