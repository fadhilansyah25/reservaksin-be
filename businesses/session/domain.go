package session

import (
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/businesses/vaccine"
	"time"
)

type Domain struct {
	Id                string
	HealthFacilitesID string
	HealthFacilites   healthFacilities.Domain
	NameSession       string
	Capacity          int
	VaccineID         string
	Vaccine           vaccine.Domain
	Date              string
	Tahap             string
	StartSession      string
	EndSession        string
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Result struct {
	Session  Domain
	Distance float64
}

type Service interface {
	Create(data *Domain) (Domain, error)
	FetchAll() ([]Domain, error)
	GetByLatLong(lat float64, lng float64) ([]Result, error)
	GetByID(id string) (Domain, error)
	// Update(id string, data *Domain) (Domain, error)
	// Delete(id string) (string, error)
}

type Repository interface {
	Create(data *Domain) (Domain, error)
	FetchAll() ([]Domain, error)
	GetByLatLong(lat, lng float64) ([]Result, error)
	GetByID(id string) (Domain, error)
	// Update(id string, data *Domain) (Domain, error)
	// Delete(id string) (string, error)
}
