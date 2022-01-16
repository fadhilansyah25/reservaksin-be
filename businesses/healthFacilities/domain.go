package healthFacilities

import (
	"ca-reservaksin/businesses/currentAddress"
	"time"
)

type Domain struct {
	ID               string
	AdminId          string
	NameFacilites    string
	CurrentAddressID string
	NoTelp           string
	CurrentAddress   currentAddress.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	Create(data *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByIdAdmin(id string) ([]Domain, error)
	FetchAll() ([]Domain, error)
}

type Repository interface {
	Create(data *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByIdAdmin(id string) ([]Domain, error)
	FetchAll() ([]Domain, error)
}
