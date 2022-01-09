package currentAddress

import (
	"ca-reservaksin/businesses/currentAddress"

	"gorm.io/gorm"
)

type CurrentAddress struct {
	gorm.Model
	Id        string  `json:"id" gorm:"PrimaryKey; NOT NULL"`
	Alamat    string  `json:"alamat"`
	Provinsi  string  `json:"provinsi"`
	Kota      string  `json:"kota"`
	Kecamatan string  `json:"kecamatan"`
	Kelurahan string  `json:"kelurahan"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

func (rec *CurrentAddress) ToDomain() currentAddress.Domain {
	return currentAddress.Domain{
		Id:        rec.Id,
		Alamat:    rec.Alamat,
		Provinsi:  rec.Provinsi,
		Kota:      rec.Kota,
		Kecamatan: rec.Kecamatan,
		Kelurahan: rec.Kelurahan,
		Lat:       rec.Lat,
		Lng:       rec.Lng,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func FromDomain(address currentAddress.Domain) *CurrentAddress {
	return &CurrentAddress{
		Id:        address.Id,
		Alamat:    address.Alamat,
		Provinsi:  address.Provinsi,
		Kota:      address.Kota,
		Kecamatan: address.Kecamatan,
		Kelurahan: address.Kelurahan,
		Lat:       address.Lat,
		Lng:       address.Lng,
	}
}
