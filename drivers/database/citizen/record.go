package citizen

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/drivers/database/currentAddress"

	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	Id                 string                        `json:"id" gorm:"Primarykey; Not Null"`
	Email              string                        `json:"email"`
	NoHp               string                        `json:"nohp"`
	FullName           string                        `json:"fullname"`
	Password           string                        `json:"password"`
	NoKK               string                        `json:"nokk"`
	Nik                string                        `json:"nik" gorm:"unique"`
	DateofBirth        string                        `json:"dob"`
	FamilyRelationship string                        `json:"relationship"`
	Gender             string                        `json:"gender"`
	MarriageStatus     string                        `json:"status"`
	Role               string                        `json:"role"`
	CurrentAddressID   string                        `gorm:"size:191" json:"current_address_id"`
	CurrentAddress     currentAddress.CurrentAddress `gorm:"foreignKey:CurrentAddressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ImageURI           string                        `json:"imageURI"`
}

func (rec *Citizen) ToDomain() citizen.Domain {
	return citizen.Domain{
		Id:                 rec.Id,
		Email:              rec.Email,
		NoHp:               rec.NoHp,
		FullName:           rec.FullName,
		Password:           rec.Password,
		NoKK:               rec.NoKK,
		Nik:                rec.Nik,
		DateofBirth:        rec.DateofBirth,
		FamilyRelationship: rec.FamilyRelationship,
		Gender:             rec.Gender,
		MarriageStatus:     rec.MarriageStatus,
		Role:               rec.Role,
		CurrentAddressID:   rec.CurrentAddressID,
		CurrentAddress:     rec.CurrentAddress.ToDomain(),
		ImageURI:           rec.ImageURI,
	}
}

func FromDomain(domain citizen.Domain) *Citizen {
	return &Citizen{
		Id:                 domain.Id,
		Email:              domain.Email,
		NoHp:               domain.NoHp,
		FullName:           domain.FullName,
		Password:           domain.Password,
		NoKK:               domain.NoKK,
		Nik:                domain.Nik,
		DateofBirth:        domain.DateofBirth,
		FamilyRelationship: domain.FamilyRelationship,
		Gender:             domain.Gender,
		MarriageStatus:     domain.MarriageStatus,
		Role:               domain.Role,
		CurrentAddress:     *currentAddress.FromDomain(domain.CurrentAddress),
		ImageURI:           domain.ImageURI,
	}
}

func ToArrayOfDomain(rec []Citizen) []citizen.Domain {
	domainArray := []citizen.Domain{}

	for _, val := range rec {
		domainArray = append(domainArray, val.ToDomain())
	}

	return domainArray
}
