package citizen

import (
	"ca-reservaksin/businesses/currentAddress"
	"time"
)

type Domain struct {
	Id                 string
	Email              string
	NoHp               string
	FullName           string
	Password           string
	NoKK               string
	Nik                string
	DateofBirth        string
	Gender             string
	Role               string
	FamilyRelationship string
	MarriageStatus     string
	CurrentAddressID   string
	CurrentAddress     currentAddress.Domain
	ImageURI           string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	Login(email_nik, password string) (Domain, string, error)
	Update(id string, data *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByID(id string) (Domain, error)
	GetByAdminID(adminID string) ([]Domain, error)
	GetByNoKK(noKK string) ([]Domain, error)
}

type Repository interface {
	Register(data *Domain) (Domain, error)
	GetByID(id string) (Domain, error)
	GetByEmail(email string) (Domain, error)
	GetByNIK(nik string) (Domain, error)
	GetByEmailOrNIK(email_or_nik string) (Domain, error)
	GetByNoKK(noKK string) ([]Domain, error)
	GetByAdminID(adminID string) ([]Domain, error)
	Delete(id string) (string, error)
	Update(id string, data *Domain) (Domain, error)
}
