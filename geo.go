package main

import (
	"fmt"
	geo "github.com/kellydunn/golang-geo"
)

type Geocoder interface {
	Geocode(query string) *Point
	ReverseGeocode(point *Point) string
}

type GeocodeFetcher struct{}

type Point struct {
	Lat float64
	Lng float64
}

func (GeocodeFetcher) Geocode(address string) *Point {
	g := &geo.GoogleGeocoder{}

	point, err := g.Geocode(address)

	if err != nil {
		panic(err)
	}

	return &Point{
		Lat: point.Lat(),
		Lng: point.Lng(),
	}
}

func (GeocodeFetcher) ReverseGeocode(point *Point) string {
	g := &geo.GoogleGeocoder{}

	geoPoint := geo.NewPoint(point.Lat, point.Lng)

	address, err := g.ReverseGeocode(geoPoint)

	if err != nil {
		panic(err)
	}

	return address
}

func main() {
	var customGeocode GeocodeFetcher
	point := customGeocode.Geocode("Cochrane 220")

	fmt.Println("Lat: ", point.Lat)
	fmt.Println("Long: ", point.Lng)

	address := customGeocode.ReverseGeocode(point)

	fmt.Println("Address: ", address)

}
