package main

import (
	"fmt"
	geo "github.com/kellydunn/golang-geo"
)

func main() {

	g := &geo.GoogleGeocoder{}

	point, err := g.Geocode("Lord Cochrane 220")

	if err != nil {
		panic(err)
	}

	fmt.Println("Lat: ", point.Lat())
	fmt.Println("Long: ", point.Lng())

	address, err := g.ReverseGeocode(point)

	if err != nil {
		panic(err)
	}

	fmt.Println("Address: ", address)
}
