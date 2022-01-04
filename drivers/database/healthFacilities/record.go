package healthFacilities

import (
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/drivers/database/admin"
	"ca-reservaksin/drivers/database/currentAddress"

	"gorm.io/gorm"
)

type HealthFacilities struct {
	gorm.Model
	Id               string                        `json:"id" gorm:"primaryKey"`
	NameFacilites    string                        `json:"name_facilities"`
	AdminId          string                        `gorm:"size:191" json:"admin_id"`
	CurrentAddressID string                        `gorm:"size:191" json:"current_Address_id"`
	NoTelp           string                        `json:"no_telp"`
	Admin            admin.Admin                   `gorm:"foreignKey:AdminId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	CurrentAddress   currentAddress.CurrentAddress `gorm:"foreignKey:CurrentAddressID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
