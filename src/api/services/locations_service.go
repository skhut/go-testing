package services

import (
	"fmt"

	locations "github.com/skhut/go-testing/src/api/domain/locations"
	locations_provider "github.com/skhut/go-testing/src/api/providers/locations"
	"github.com/skhut/go-testing/src/api/utils/errors"
)

type locationsService struct{}

type locationsServiceInterface interface {
	GetCountry(countryId string) (*locations.Country, *errors.ApiError)
}

var (
	LocationsService locationsServiceInterface
)

func init() {
	LocationsService = &locationsService{}
}

func (s locationsService) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	fmt.Println("Inside service")
	return locations_provider.GetCountry(countryId)
}
