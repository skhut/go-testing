package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	rest "github.com/federicoleon/golang-restclient/rest"
	"github.com/gin-gonic/gin"
	"github.com/skhut/go-testing/src/api/domain/locations"
	"github.com/skhut/go-testing/src/api/services"
	"github.com/skhut/go-testing/src/api/utils/errors"
	"github.com/stretchr/testify/assert"
)

var (
	getCountryFunc func(countryId string) (*locations.Country, *errors.ApiError)
)

type locationsServiceMock struct{}

func (*locationsServiceMock) GetCountry(countryId string) (*locations.Country, *errors.ApiError) {
	return getCountryFunc(countryId)
}
func TestMain(m *testing.M) {
	rest.StartMockupServer()
	os.Exit(m.Run())
}

func TestGetCountryNotFound(t *testing.T) {
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return nil, &errors.ApiError{Status: http.StatusNotFound, Message: "Country not found"}
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "https://api.mercadolibre.com/countries/AR", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusNotFound, response.Code)
	var apiError errors.ApiError
	err := json.Unmarshal(response.Body.Bytes(), &apiError)

	assert.Nil(t, err)
	assert.EqualValues(t, http.StatusNotFound, apiError.Status)
	assert.EqualValues(t, "Country not found", apiError.Message)
}

func TestGetCountryNoError(t *testing.T) {
	getCountryFunc = func(countryId string) (*locations.Country, *errors.ApiError) {
		return &locations.Country{Id: "AR", Name: "Argentina"}, nil
	}
	services.LocationsService = &locationsServiceMock{}

	response := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(response)
	c.Request, _ = http.NewRequest(http.MethodGet, "https://api.mercadolibre.com/countries/AR", nil)
	c.Params = gin.Params{
		{Key: "country_id", Value: "AR"},
	}

	GetCountry(c)

	assert.EqualValues(t, http.StatusOK, response.Code)
	var country locations.Country
	err := json.Unmarshal(response.Body.Bytes(), &country)
	fmt.Println(string(response.Body.Bytes()))
	assert.Nil(t, err)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
}
