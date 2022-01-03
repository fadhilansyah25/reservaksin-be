package healthfacilities

import (
	"ca-reservaksin/businesses/currentAddress"
	"time"
)

type Domain struct {
	Id               string
	AdminId          string
	NameFacilites    string
	CurrentAddressID string
	NoTelp           string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Service interface {
	Create(data *Domain, address *currentAddress.Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
}

type Repository interface {
	Create(data *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
}
