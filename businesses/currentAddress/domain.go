package currentaddress

import (
	"time"
)

type Domain struct {
	Id        int
	Alamat    string
	Desa      string
	Kota      string
	Kecamatan string
	Provinsi  string
	Latd      float64
	Lngtd     float64
	CreatedAt time.Time
	UpdatedAt time.Time
}
