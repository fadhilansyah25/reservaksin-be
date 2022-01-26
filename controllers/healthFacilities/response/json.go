package response

import (
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/controllers/currentAddress/response"
	"time"
)

type HealthFacilities struct {
	ID               string                  `json:"id"`
	NameFacilites    string                  `json:"name_facilities"`
	NoTelp           string                  `json:"no_telp"`
	AdminId          string                  `json:"admin"`
	CurrentAddressID string                  `json:"current_address_id"`
	CurrentAddress   response.CurrentAddress `json:"current_Address"`
	CreatedAt        time.Time               `json:"created_at"`
	UpdatedAt        time.Time               `json:"updated_at"`
}

func FromDomain(data healthFacilities.Domain) *HealthFacilities {
	return &HealthFacilities{
		ID:               data.ID,
		NameFacilites:    data.NameFacilites,
		NoTelp:           data.NoTelp,
		AdminId:          data.AdminId,
		CurrentAddressID: data.CurrentAddressID,
		CurrentAddress:   response.CurrentAddress(data.CurrentAddress),
		CreatedAt:        data.CreatedAt,
		UpdatedAt:        data.UpdatedAt,
	}
}

func FromDomainArray(domain []healthFacilities.Domain) *[]HealthFacilities {
	res := []HealthFacilities{}

	for _, val := range domain {
		res = append(res, *FromDomain(val))
	}
	return &res
}

type forMapsResponse struct {
	ID            string  `json:"id"`
	NameFacilites string  `json:"name_facilities"`
	Lat           float64 `json:"lat"`
	Lng           float64 `json:"lng"`
	Address       string  `json:"address"`
}

func FromDomainToMapsResponse(data healthFacilities.Domain) *forMapsResponse {
	return &forMapsResponse{
		ID:            data.ID,
		NameFacilites: data.NameFacilites,
		Lat:           data.CurrentAddress.Lat,
		Lng:           data.CurrentAddress.Lng,
		Address: data.CurrentAddress.Alamat + ", " +
			data.CurrentAddress.Kelurahan + ", " +
			data.CurrentAddress.Kecamatan + ", " +
			data.CurrentAddress.Kota + ", " +
			data.CurrentAddress.Provinsi,
	}
}

func FromDomainArrayToMapsResponse(domain []healthFacilities.Domain) *[]forMapsResponse {
	res := []forMapsResponse{}
	for _, val := range domain {
		res = append(res, *FromDomainToMapsResponse(val))
	}
	return &res
}
