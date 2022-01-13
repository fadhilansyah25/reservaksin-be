package booking

import (
	"ca-reservaksin/businesses/booking"
	"ca-reservaksin/drivers/database/citizen"
	"ca-reservaksin/drivers/database/session"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	Id           string          `json:"id" gorm:"PrimaryKey;Unique"`
	NomorAntrian int             `json:"nomor_antrian"`
	Status       string          `json:"status" gorm:"default:'booked'"`
	CitizenId    string          `gorm:"size:191" json:"citizen_id"`
	Citizen      citizen.Citizen `gorm:"foreignKey:CitizenId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	SessionId    string          `gorm:"size:191" json:"session_id"`
	Session      session.Session `gorm:"foreignKey:SessionId;constraint:OnUpdate:CASCADE,OnDelete:RESTRICT;"`
	Date         string          `gorm:"type:Date" json:"date"`
	SessionTime  string          `json:"start_session" gorm:"type:Datetime" time_format:"2006-01-02 15:04"`
}

func (rec *Booking) ToDomain() booking.Domain {
	return booking.Domain{
		Id:           rec.Id,
		NomorAntrian: rec.NomorAntrian,
		Status:       rec.Status,
		CitizenId:    rec.CitizenId,
		Citizen:      rec.Citizen.ToDomain(),
		SessionId:    rec.SessionId,
		Session:      rec.Session.ToDomain(),
		Date:         rec.Date,
		SessionTime:  rec.SessionTime,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
	}
}

func FromDomain(dataBooking *booking.Domain) *Booking {
	return &Booking{
		Id:           dataBooking.Id,
		NomorAntrian: dataBooking.NomorAntrian,
		Status:       dataBooking.Status,
		CitizenId:    dataBooking.CitizenId,
		SessionId:    dataBooking.SessionId,
		Date:         dataBooking.Date,
		SessionTime:  dataBooking.SessionTime,
	}
}

func ToArrayOfDomain(rec []Booking) []booking.Domain {
	domainArray := []booking.Domain{}

	for _, val := range rec {
		domainArray = append(domainArray, val.ToDomain())
	}

	return domainArray
}
