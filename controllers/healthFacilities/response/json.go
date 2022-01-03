package response

import (
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/businesses/healthFacilities"
	"time"
)

type HealthFacilities struct {
	Id             string         `json:"id"`
	NameFacilites  string         `json:"name_facilities"`
	NoTelp         string         `json:"no_telp"`
	AdminId        string         `json:"admin"`
	CurrentAddress CurrentAddress `json:"current_Address"`
	CreatedAt      time.Time      `json:"created_at"`
	UpdatedAt      time.Time      `json:"updated_at"`
}

type CurrentAddress struct {
	Id        string  `json:"id"`
	Alamat    string  `json:"alamat"`
	Provinsi  string  `json:"provinsi"`
	Kota      string  `json:"kota"`
	Kecamatan string  `json:"kecamatan"`
	Kelurahan string  `json:"kelurahan"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

func FromDomain(data *healthFacilities.Domain, address *currentAddress.Domain) *HealthFacilities {
	return &HealthFacilities{
		Id:            data.Id,
		NameFacilites: data.NameFacilites,
		NoTelp:        data.NoTelp,
		AdminId:       data.AdminId,
		CurrentAddress: CurrentAddress{
			Id:        data.CurrentAddressID,
			Alamat:    address.Alamat,
			Provinsi:  address.Provinsi,
			Kota:      address.Kota,
			Kecamatan: address.Kecamatan,
			Kelurahan: address.Kelurahan,
			Lat:       address.Lat,
			Lng:       address.Lng,
		},
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
