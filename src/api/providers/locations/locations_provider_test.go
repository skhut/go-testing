package locations

import (
	"net/http"
	"testing"

	"github.com/skhut/go-testing/src/api/providers/locations"
	"github.com/stretchr/testify/assert"
)

func TestGetCountryRestclientError(t *testing.T) {
	//Init
	//Execution
	country, err := locations.GetCountry("AR")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid Restclient error when trying to get country AR", err.Message)
	//Validation
}

func TestGetCountryNotFound(t *testing.T) {
	country, err := locations.GetCountry("AT")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusNotFound, err.Status)
	assert.EqualValues(t, "country not found", err.Message)
}

func TestGetCountryInvalidErrorInterface(t *testing.T) {
	country, err := locations.GetCountry("AT")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "Invalid error interface when getting country AT", err.Message)
}

func TestGetCountryInvalidJsonResponse(t *testing.T) {
	country, err := locations.GetCountry("AT")
	assert.Nil(t, country)
	assert.NotNil(t, err)
	assert.EqualValues(t, http.StatusInternalServerError, err.Status)
	assert.EqualValues(t, "error when trying to unmarshal country data for AT", err.Message)
}

func TestGetCountryNoError(t *testing.T) {
	country, err := locations.GetCountry("AR")
	assert.Nil(t, err)
	assert.NotNil(t, country)
	assert.EqualValues(t, "AR", country.Id)
	assert.EqualValues(t, "Argentina", country.Name)
	assert.EqualValues(t, "GMT-03:00", country.TimeZone)
	assert.EqualValues(t, 24, len(country.States))
}
