package citizen

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/drivers/database/currentAddress"

	"gorm.io/gorm"
)

type Citizen struct {
	gorm.Model
	ID                 string                        `json:"id" gorm:"Primarykey; Not Null"`
	Email              string                        `json:"email"`
	NoHp               string                        `json:"nohp"`
	Username           string                        `json:"username"`
	Password           string                        `json:"password"`
	NoKK               string                        `json:"nokk"`
	Nik                string                        `json:"nik" gorm:"unique"`
	DateofBirth        string                        `json:"dob"`
	FamilyRelationship string                        `json:"relationship"`
	Gender             string                        `json:"gender"`
	MarriageStatus     string                        `json:"status"`
	Role               string                        `json:"role"`
	Address            string                        `json:"alamat"`
	Desa               string                        `json:"desa"`
	Kota               string                        `json:"kota"`
	Kecamatan          string                        `json:"kecamatan"`
	Provinsi           string                        `json:"provinsi"`
	CurrentAddressID   string                        `gorm:"size:191" json:"current_Address_id"`
	CurrentAddress     currentAddress.CurrentAddress `gorm:"foreignKey:CurrentAddressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ImageURI           string                        `json:"imageURI"`
}

func (rec *Citizen) toDomain() citizen.Domain {
	return citizen.Domain{
		ID:                 rec.ID,
		Email:              rec.Email,
		NoHp:               rec.NoHp,
		Username:           rec.Username,
		Password:           rec.Password,
		NoKK:               rec.NoKK,
		Nik:                rec.Nik,
		DateofBirth:        rec.DateofBirth,
		FamilyRelationship: rec.FamilyRelationship,
		Gender:             rec.Gender,
		MarriageStatus:     rec.MarriageStatus,
		Role:               rec.Role,
		Address:            rec.Address,
		Desa:               rec.Desa,
		Kota:               rec.Kota,
		Kecamatan:          rec.Kecamatan,
		Provinsi:           rec.Provinsi,
		CurrentAddressID:   rec.CurrentAddressID,
		CurrentAddress:     rec.CurrentAddress.ToDomain(),
		ImageURI:           rec.ImageURI,
	}
}

func fromDomain(domain citizen.Domain) *Citizen {
	return &Citizen{
		ID:                 domain.ID,
		Email:              domain.Email,
		NoHp:               domain.NoHp,
		Username:           domain.Username,
		Password:           domain.Password,
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
		CurrentAddress:     *currentAddress.FromDomain(domain.CurrentAddress),
		ImageURI:           domain.ImageURI,
	}
}

func ToArrayOfDomain(rec []Citizen) []citizen.Domain {
	domainArray := []citizen.Domain{}

	for _, val := range rec {
		domainArray = append(domainArray, val.toDomain())
	}

	return domainArray
}
