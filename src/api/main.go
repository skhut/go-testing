package main

import (
	"fmt"

	"github.com/skhut/go-testing/src/api/providers/locations"
)

func main() {
	country, err := locations.GetCountry("ARB")
	fmt.Println(err)
	fmt.Println(country)
}
