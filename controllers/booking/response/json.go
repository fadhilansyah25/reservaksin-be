package response

import (
	"ca-reservaksin/businesses/booking"
	"time"
)

type Booking struct {
	Id           string `json:"id"`
	NomorAntrian int    `json:"nomor_antrian"`
	Status       string `json:"status"`
	CitizenId    string `json:"citizen_id"`
	// Citizen      _responseCitizen.CitizenResponse `json:"citizen"`
	Citizen   string `json:"citizen"`
	SessionId string `json:"session_id"`
	// Session     _responseSession.Session `json:"session"`
	Session     string    `json:"session_place"`
	Date        string    `json:"date"`
	SessionTime string    `json:"session_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FromDomain(data booking.Domain) *Booking {
	return &Booking{
		Id:           data.Id,
		NomorAntrian: data.NomorAntrian,
		Status:       data.Status,
		CitizenId:    data.CitizenId,
		Citizen:      data.Citizen.FullName,
		SessionId:    data.SessionId,
		Session:      data.Session.HealthFacilites.NameFacilites,
		Date:         data.Date,
		SessionTime:  data.SessionTime,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}
