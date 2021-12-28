package vaccine

import "time"

type Domain struct {
	Id        int
	Nama      string
	Stok      int
	CreatedAt time.Time
	UpdatedAt time.Time
}
