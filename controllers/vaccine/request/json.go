package request

import "ca-reservaksin/businesses/vaccine"

type Vaccine struct {
	AdminID    string `json:"admin_id"`
	NamaVaksin string `json:"nama_vaksin"`
	Stok       int    `json:"stok"`
}

func (req *Vaccine) ToDomain() *vaccine.Domain {
	return &vaccine.Domain{
		AdminID:    req.AdminID,
		NamaVaksin: req.NamaVaksin,
		Stok:       req.Stok,
	}
}
