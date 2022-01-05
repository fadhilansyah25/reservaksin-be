package request

import (
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/controllers/currentAddress/request"
)

type HealthFacilities struct {
	AdminId        string                 `json:"admin_id"`
	NameFacilites  string                 `json:"name_facilities"`
	NoTelp         string                 `json:"no_telp"`
	CurrentAddress request.CurrentAddress `json:"current_address"`
}

func (req *HealthFacilities) ToDomain() *healthFacilities.Domain {
	return &healthFacilities.Domain{
		AdminId:        req.AdminId,
		NameFacilites:  req.NameFacilites,
		NoTelp:         req.NoTelp,
		CurrentAddress: *req.CurrentAddress.ToDomain(),
	}
}
