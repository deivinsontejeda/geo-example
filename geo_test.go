package main

import (
	"testing"
)

type GeocodeFetcherStub struct{}

func (GeocodeFetcherStub) Geocode(address string) *Point {
	return &Point{
		Lat: 10.00,
		Lng: 11.0,
	}
}

func (GeocodeFetcherStub) ReverseGeocode(p *Point) string {
	return "Some Address"
}

var stubGeocode GeocodeFetcherStub

func TestGeocode(t *testing.T) {
	point := stubGeocode.Geocode("Some Address")

	if point.Lat != 10.00 || point.Lng != 11.00 {
		t.Error("TestGeocode() failed", point)
	}
}

func TestReverseGeocode(t *testing.T) {
	point := &Point{
		Lat: 10.00,
		Lng: 11.00,
	}
	address := stubGeocode.ReverseGeocode(point)
	if address != "Some Address" {
		t.Error("TestReverseGeocode() failed ", address)
	}
}
