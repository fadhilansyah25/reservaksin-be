package healthFacilities

import (
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/drivers/database/admin"
	"ca-reservaksin/drivers/database/currentAddress"

	"gorm.io/gorm"
)

type HealthFacilities struct {
	gorm.Model
	ID               string                        `json:"id" gorm:"primaryKey; NOT NULL"`
	NameFacilites    string                        `json:"name_facilities"`
	AdminId          string                        `gorm:"size:191" json:"admin_id"`
	CurrentAddressID string                        `gorm:"size:191" json:"current_Address_id"`
	NoTelp           string                        `json:"no_telp"`
	Admin            admin.Admin                   `gorm:"foreignKey:AdminId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CurrentAddress   currentAddress.CurrentAddress `gorm:"foreignKey:CurrentAddressID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}

func (rec *HealthFacilities) ToDomain() healthFacilities.Domain {
	return healthFacilities.Domain{
		ID:               rec.ID,
		NameFacilites:    rec.NameFacilites,
		AdminId:          rec.AdminId,
		CurrentAddressID: rec.CurrentAddressID,
		CurrentAddress:   rec.CurrentAddress.ToDomain(),
		NoTelp:           rec.NoTelp,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
	}
}

func FromDomain(facilities healthFacilities.Domain) *HealthFacilities {
	return &HealthFacilities{
		ID:               facilities.ID,
		NameFacilites:    facilities.NameFacilites,
		AdminId:          facilities.AdminId,
		CurrentAddressID: facilities.CurrentAddressID,
		NoTelp:           facilities.NoTelp,
		CurrentAddress:   *currentAddress.FromDomain(facilities.CurrentAddress),
	}
}
