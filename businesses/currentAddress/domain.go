package currentAddress

import (
	"time"
)

type Domain struct {
	Id        string
	Alamat    string
	Provinsi  string
	Kota      string
	Kecamatan string
	Kelurahan string
	Lat       float64
	Lng       float64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Service interface {
	Create(dataAddress *Domain) (Domain, error)
	Update(id string, dataAddress *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByID(id string) (Domain, error)
}

type Repository interface {
	Create(dataAddress *Domain) (Domain, error)
	Update(id string, dataAddress *Domain) (Domain, error)
	Delete(id string) (string, error)
	GetByID(id string) (Domain, error)
}
