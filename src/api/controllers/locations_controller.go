package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/skhut/go-testing/src/api/services"
)

//GetCountry to get details of a country
func GetCountry(c *gin.Context) {
	fmt.Println("Inside controller")
	country, err := services.LocationsService.GetCountry(c.Param("country_id"))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, country)
}
