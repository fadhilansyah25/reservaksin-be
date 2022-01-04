package vaccine

import (
	"ca-reservaksin/businesses/vaccine"
	"time"

	"gorm.io/gorm"
)

type Vaccine struct {
	gorm.Model
	Id         string    `json:"id" gorm:"PrimaryKey; Not Null"`
	NamaVaksin string    `json:"nama_vaksin"`
	Stok       int       `json:"stok"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func (rec *Vaccine) toDomain() vaccine.Domain {
	return vaccine.Domain{
		Id:         rec.Id,
		NamaVaksin: rec.NamaVaksin,
		Stok:       rec.Stok,
		CreatedAt:  rec.CreatedAt,
		UpdatedAt:  rec.UpdatedAt,
	}
}

func fromDomain(vaccineDomain *vaccine.Domain) *Vaccine {
	return &Vaccine{
		Id:         vaccineDomain.Id,
		NamaVaksin: vaccineDomain.NamaVaksin,
		Stok:       vaccineDomain.Stok,
	}
}

func ToArrayOfDomain(rec []Vaccine) []vaccine.Domain {
	domainArray := []vaccine.Domain{}

	for _, val := range rec {
		domainArray = append(domainArray, val.toDomain())
	}

	return domainArray
}
