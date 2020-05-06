package models

//Bus is the object
type Bus struct {
	LicenseNo  string
	Location   Location
	Attributes Attributes
}

//Attributes keeps track bus attributes
type Attributes struct {
	Availability int
	PathNo       int
	AC           int
}

//Location keeps track of location
type Location struct {
	Latitude  float32
	Longitude float32
	Speed     float64
	Heading   float64
}
