package response

import (
	"ca-reservaksin/businesses/vaccine"
	"time"
)

type Vaccine struct {
	Id         string    `json:"id"`
	NamaVaksin string    `json:"nama_vaksin"`
	AdminID    string    `json:"admin_id"`
	Stok       int       `json:"stok"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

func FromDomain(vaccine vaccine.Domain) *Vaccine {
	return &Vaccine{
		Id:         vaccine.Id,
		NamaVaksin: vaccine.NamaVaksin,
		AdminID:    vaccine.AdminID,
		Stok:       vaccine.Stok,
		CreatedAt:  vaccine.CreatedAt,
		UpdatedAt:  vaccine.UpdatedAt,
	}
}

func FromDomainArray(domain []vaccine.Domain) *[]Vaccine {
	res := []Vaccine{}
	for _, val := range domain {
		res = append(res, *FromDomain(val))
	}
	return &res
}
