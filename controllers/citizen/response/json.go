package response

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/controllers/currentAddress/response"
	"time"
)

type CitizenResponse struct {
	ID                 string                  `json:"id"`
	Email              string                  `json:"email"`
	NoHp               string                  `json:"nohp"`
	Username           string                  `json:"username"`
	NoKK               string                  `json:"nokk"`
	Nik                string                  `json:"nik"`
	DateofBirth        string                  `json:"dob"`
	FamilyRelationship string                  `json:"relationship"`
	Gender             string                  `json:"gender"`
	MarriageStatus     string                  `json:"status"`
	Role               string                  `json:"role"`
	Address            string                  `json:"alamat"`
	Desa               string                  `json:"desa"`
	Kota               string                  `json:"kota"`
	Kecamatan          string                  `json:"kecamatan"`
	Provinsi           string                  `json:"provinsi"`
	CurrentAddressID   string                  `json:"current_address_id"`
	CurrentAddress     response.CurrentAddress `json:"current_Address"`
	CreatedAt          time.Time               `json:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at"`
	Token              string                  `json:"token"`
}

func FromDomain(domain citizen.Domain) *CitizenResponse {
	return &CitizenResponse{
		ID:                 domain.ID,
		Email:              domain.Email,
		NoHp:               domain.NoHp,
		Username:           domain.Username,
		NoKK:               domain.NoKK,
		Nik:                domain.Nik,
		DateofBirth:        domain.DateofBirth,
		FamilyRelationship: domain.FamilyRelationship,
		Gender:             domain.Gender,
		MarriageStatus:     domain.MarriageStatus,
		Role:               domain.Role,
		Address:            domain.Address,
		Desa:               domain.Desa,
		Kota:               domain.Kota,
		Kecamatan:          domain.Kecamatan,
		Provinsi:           domain.Provinsi,
		CurrentAddressID:   domain.CurrentAddressID,
		CurrentAddress:     response.CurrentAddress(domain.CurrentAddress),
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}
