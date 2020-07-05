package locations

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/federicoleon/golang-restclient/rest"
	"github.com/skhut/go-testing/src/api/domain/locations"
	"github.com/skhut/go-testing/src/api/utils/errors"
)

const (
	urlGetCountry = "https://api.mercadolibre.com/countries/%s"
)

//GetCountry to get the countries from Api
func GetCountry(countryID string) (*locations.Country, *errors.ApiError) {
	fmt.Println(fmt.Sprintf(urlGetCountry, countryID))
	resp := rest.Get(fmt.Sprintf(urlGetCountry, countryID))
	fmt.Printf("Obtained status code:%d", resp.StatusCode)
	//fmt.Println(rest.Head(urlGetCountry).String())
	if resp == nil || resp.Response == nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("Invalid Restclient error when trying to get country %s", countryID),
		}
	}

	if resp.StatusCode > 299 {
		var apiErr errors.ApiError
		if err := json.Unmarshal(resp.Bytes(), &apiErr); err != nil {
			return nil, &errors.ApiError{
				Status:  http.StatusInternalServerError,
				Message: fmt.Sprintf("Invalid error interface when getting country %s", countryID),
			}
		}
		return nil, &apiErr
	}

	var result locations.Country
	if err := json.Unmarshal(resp.Bytes(), &result); err != nil {
		return nil, &errors.ApiError{
			Status:  http.StatusInternalServerError,
			Message: fmt.Sprintf("error when trying to unmarshal country data for %s", countryID),
		}
	}
	return &result, nil
}
