package healthFacilities

import (
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/drivers/database/admin"
	"ca-reservaksin/drivers/database/currentAddress"

	"gorm.io/gorm"
)

type HealthFacilities struct {
	gorm.Model
	Id               string                        `json:"id" gorm:"primary key"`
	NameFacilites    string                        `json:"name_facilities"`
	AdminId          string                        `json:"admin_id"`
	Admin            admin.Admin                   `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:RESTRICT;"`
	CurrentAddressID string                        `json:"current_Address_id"`
	CurrentAddress   currentAddress.CurrentAddress `gorm:"constraint:OnUpdate:NO ACTION,OnDelete:NO ACTION;"`
	NoTelp           string                        `json:"no_telp"`
}

func (rec *HealthFacilities) ToDomain() healthFacilities.Domain {
	return healthFacilities.Domain{
		Id:               rec.Id,
		NameFacilites:    rec.NameFacilites,
		AdminId:          rec.AdminId,
		CurrentAddressID: rec.CurrentAddressID,
		NoTelp:           rec.NoTelp,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

func FromDomain(facilities healthFacilities.Domain) *HealthFacilities {
	return &HealthFacilities{
		Id:               facilities.Id,
		NameFacilites:    facilities.NameFacilites,
		AdminId:          facilities.AdminId,
		CurrentAddressID: facilities.CurrentAddressID,
		NoTelp:           facilities.NoTelp,
	}
}
