package citizen

import (
	"time"
)

type Domain struct {
	Id                 int
	Email              string
	NoHp               int16
	Username           string
	Password           string
	NoKK               uint64
	Nik                uint64
	DateofBirth        string
	FamilyRelationship string
	Gender             string
	MarriageStatus     string
	Role               string
	Address            string
	Desa               string
	Kota               string
	Kecamatan          string
	Provinsi           string
	ImageURI           string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type Service interface {
	Register(data *Domain) (Domain, error)
	Login(email string, password string) (string, error)
}

type Repository interface {
	GetByEmail(email string) (Domain, error)
}
