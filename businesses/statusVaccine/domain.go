package statusvaccine

import "time"

type Domain struct {
	Id               int
	AdminID          int
	BookingSessionID int
	StatusVaksin     int
	CreatedAt        time.Time
}
