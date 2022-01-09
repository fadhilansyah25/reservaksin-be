package response

import (
	"ca-reservaksin/businesses/currentAddress"
	"time"
)

type CurrentAddress struct {
	Id        string    `json:"id"`
	Alamat    string    `json:"alamat"`
	Provinsi  string    `json:"provinsi"`
	Kota      string    `json:"kota"`
	Kecamatan string    `json:"kecamatan"`
	Kelurahan string    `json:"kelurahan"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"lng"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(data currentAddress.Domain) *CurrentAddress {
	return &CurrentAddress{
		Id:        data.Id,
		Alamat:    data.Alamat,
		Provinsi:  data.Provinsi,
		Kota:      data.Kota,
		Kecamatan: data.Kecamatan,
		Kelurahan: data.Kelurahan,
		Lat:       data.Lat,
		Lng:       data.Lng,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
