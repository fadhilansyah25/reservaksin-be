package request

import (
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/controllers/currentAddress/request"
)

type HealthFacilities struct {
	AdminId        string                 `json:"admin_id"`
	NameFacilites  string                 `json:"name_facilities"`
	NoTelp         string                 `json:"no_telp"`
	CurrentAddress request.CurrentAddress `json:"current_address"`
}

func (req *HealthFacilities) ToDomain() (*healthFacilities.Domain, *currentAddress.Domain) {
	return &healthFacilities.Domain{
			AdminId:       req.AdminId,
			NameFacilites: req.NameFacilites,
			NoTelp:        req.NoTelp,
		}, &currentAddress.Domain{
			Alamat:    req.CurrentAddress.Alamat,
			Provinsi:  req.CurrentAddress.Provinsi,
			Kota:      req.CurrentAddress.Kota,
			Kecamatan: req.CurrentAddress.Kecamatan,
			Kelurahan: req.CurrentAddress.Kelurahan,
			Lat:       req.CurrentAddress.Lat,
			Lng:       req.CurrentAddress.Lng,
		}
}
