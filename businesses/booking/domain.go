package booking

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/businesses/session"
	"time"
)

type Domain struct {
	Id           string
	NomorAntrian int
	Status       string
	CitizenId    string
	Citizen      citizen.Domain
	SessionId    string
	Session      session.Domain
	Date         string
	SessionTime  string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type Service interface {
	BookingSession(dataBooking *Domain) (Domain, error)
	GetBySessionID(sessionID string) ([]Domain, error)
	GetByCitizenID(citizenID string) ([]Domain, error)
	GetByNoKK(noKK string) ([]Domain, error)
	UpdateStatusByID(id string, status string) (Domain, error)
}

type Repository interface {
	GetBySessionID(sessionID string) ([]Domain, error)
	GetByCitizenID(citizenID string) ([]Domain, error)
	GetByStatus(status string) ([]Domain, error)
	GetByNoKK(noKK string) ([]Domain, error)
	GetByID(id string) (Domain, error)
	Create(data *Domain) (Domain, error)
	UpdateStatusByID(id string, status string) (Domain, error)
	// Update(data *Domain) (Domain, error)
	// Delete(id string) (string, error)
}
