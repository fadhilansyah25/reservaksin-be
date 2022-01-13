package vaccine

import (
	"ca-reservaksin/businesses/vaccine"
	"ca-reservaksin/drivers/database/admin"
	"time"

	"gorm.io/gorm"
)

type Vaccine struct {
	gorm.Model
	Id         string      `json:"id" gorm:"PrimaryKey; Not Null"`
	NamaVaksin string      `json:"nama_vaksin"`
	AdminID    string      `json:"admin_id" gorm:"size:191"`
	Admin      admin.Admin `gorm:"foreignKey:AdminID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Stok       int         `json:"stok"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
}

func (rec *Vaccine) ToDomain() vaccine.Domain {
	return vaccine.Domain{
		Id:         rec.Id,
		NamaVaksin: rec.NamaVaksin,
		AdminID:    rec.AdminID,
		Stok:       rec.Stok,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(vaccineDomain *vaccine.Domain) *Vaccine {
	return &Vaccine{
		Id:         vaccineDomain.Id,
		AdminID:    vaccineDomain.AdminID,
		NamaVaksin: vaccineDomain.NamaVaksin,
		Stok:       vaccineDomain.Stok,
	}
}

func ToArrayOfDomain(rec []Vaccine) []vaccine.Domain {
	domainArray := []vaccine.Domain{}

	for _, val := range rec {
		domainArray = append(domainArray, val.ToDomain())
	}

	return domainArray
}
