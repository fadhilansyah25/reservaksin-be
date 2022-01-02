package request

import "ca-reservaksin/businesses/currentAddress"

type CurrentAddress struct {
	Alamat    string  `json:"alamat"`
	Provinsi  string  `json:"provinsi"`
	Kota      string  `json:"kota"`
	Kecamatan string  `json:"kecamatan"`
	Kelurahan string  `json:"kelurahan"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

func (req *CurrentAddress) ToDomain() *currentAddress.Domain {
	return &currentAddress.Domain{
		Alamat:    req.Alamat,
		Provinsi:  req.Provinsi,
		Kota:      req.Kota,
		Kecamatan: req.Kecamatan,
		Kelurahan: req.Kelurahan,
		Lat:       req.Lat,
		Lng:       req.Lng,
	}
}
