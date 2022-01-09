package citizen

import (
	"ca-reservaksin/businesses/currentAddress"
	"time"
)

type Domain struct {
	Id                 string
	Email              string
	NoHp               string
	Username           string
	Password           string
	NoKK               uint64
	Nik                uint64
	DateofBirth        string
	FamilyRelationship string
	Gender             string
	Role               string
	CurrentAddressID   string
	CurrentAddress     currentAddress.Domain
	ImageURI           string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	Login(email string, password string) (string, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
}

type Repository interface {
	GetByEmail(email string) (Domain, error)
	Register(data *Domain) (Domain, error)
	GetByNIK(nik string) (Domain, error)
	GetByNoKK(noKK string) ([]Domain, error)
	Update(id string) (Domain, error)
}
