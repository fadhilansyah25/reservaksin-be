package bookingsession

import (
	"time"
)

type Domain struct {
	Id           int
	CitizenId    int
	SessionId    int
	NomorAntrian int
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
