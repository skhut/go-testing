package app

import "github.com/skhut/go-testing/src/api/controllers"

func mapUrls() {
	router.GET("/locations/countries/:country_id", controllers.GetCountry)
}
