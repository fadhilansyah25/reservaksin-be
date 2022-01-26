package response

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/controllers/currentAddress/response"
	"time"
)

type CitizenResponse struct {
	Id                 string                  `json:"id"`
	Email              string                  `json:"email"`
	NoHp               string                  `json:"nohp"`
	Fullname           string                  `json:"fullname"`
	NoKK               string                  `json:"nokk"`
	Nik                string                  `json:"nik"`
	DateofBirth        string                  `json:"dob"`
	FamilyRelationship string                  `json:"relationship"`
	Gender             string                  `json:"gender"`
	MarriageStatus     string                  `json:"status"`
	Role               string                  `json:"role"`
	CurrentAddressID   string                  `json:"current_address_id"`
	CurrentAddress     response.CurrentAddress `json:"current_Address"`
	CreatedAt          time.Time               `json:"created_at"`
	UpdatedAt          time.Time               `json:"updated_at"`
}

func FromDomain(domain citizen.Domain) *CitizenResponse {
	return &CitizenResponse{
		Id:                 domain.Id,
		Email:              domain.Email,
		NoHp:               domain.NoHp,
		Fullname:           domain.FullName,
		NoKK:               domain.NoKK,
		Nik:                domain.Nik,
		DateofBirth:        domain.DateofBirth,
		FamilyRelationship: domain.FamilyRelationship,
		Gender:             domain.Gender,
		MarriageStatus:     domain.MarriageStatus,
		Role:               domain.Role,
		CurrentAddressID:   domain.CurrentAddressID,
		CurrentAddress:     response.CurrentAddress(domain.CurrentAddress),
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
	}
}

func FromDomainOfArray(domainArray []citizen.Domain) *[]CitizenResponse {
	res := []CitizenResponse{}

	for _, val := range domainArray {
		res = append(res, *FromDomain(val))
	}
	return &res
}
