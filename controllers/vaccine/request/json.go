package request

import "ca-reservaksin/businesses/vaccine"

type Vaccine struct {
	// Id         int       `json:"id"`
	NamaVaksin string `json:"nama_vaksin"`
	Stok       int    `json:"stok"`
	// CreatedAt  time.Time `json:"created_at"`
	// UpdatedAt  time.Time `json:"updated_at"`
}

func (req *Vaccine) ToDomain() *vaccine.Domain {
	return &vaccine.Domain{
		NamaVaksin: req.NamaVaksin,
		Stok:       req.Stok,
	}
}
