package healthfacilities

import (
	"time"
)

type Domain struct {
	Id               int
	AdminId          int
	NameFacilites    string
	CurrentAddressID int
	NoTelp           int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
