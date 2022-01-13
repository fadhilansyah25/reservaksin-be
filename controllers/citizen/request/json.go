package request

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/controllers/currentAddress/request"
)

type Citizen struct {
	Email              string                 `json:"email"`
	NoHp               string                 `json:"nohp"`
	FullName           string                 `json:"fullname"`
	Password           string                 `json:"password"`
	NoKK               string                 `json:"nokk"`
	Nik                string                 `json:"nik"`
	DateofBirth        string                 `json:"dob"`
	Gender             string                 `json:"gender"`
	Role               string                 `json:"role"`
	FamilyRelationship string                 `json:"relationship"`
	MarriageStatus     string                 `json:"status"`
	CurrentAddress     request.CurrentAddress `json:"current_address"`
	ImageURI           string                 `json:"image_url"`
}

type CitizenLoginEmail struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CitizenLoginNIK struct {
	Nik      string `json:"nik"`
	Password string `json:"password"`
}

func (req *Citizen) ToDomain() *citizen.Domain {
	return &citizen.Domain{
		Email:              req.Email,
		NoHp:               req.NoHp,
		FullName:           req.FullName,
		Password:           req.Password,
		NoKK:               req.NoKK,
		Nik:                req.Nik,
		DateofBirth:        req.DateofBirth,
		FamilyRelationship: req.FamilyRelationship,
		Gender:             req.Gender,
		MarriageStatus:     req.MarriageStatus,
		Role:               req.Role,
		CurrentAddress:     *req.CurrentAddress.ToDomain(),
		ImageURI:           req.ImageURI,
	}
}

func (req *CitizenLoginEmail) ToDomain() *citizen.Domain {
	return &citizen.Domain{
		Email:    req.Email,
		Password: req.Password,
	}
}

func (req *CitizenLoginNIK) ToDomain() *citizen.Domain {
	return &citizen.Domain{
		Nik:      req.Nik,
		Password: req.Password,
	}
}
