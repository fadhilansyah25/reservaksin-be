package response

import (
	"ca-reservaksin/businesses/booking"
	_responseCitizen "ca-reservaksin/controllers/citizen/response"
	_responseSession "ca-reservaksin/controllers/session/response"
	"time"
)

type BookingCitizen struct {
	BookingID    string                           `json:"booking_id"`
	NomorAntrian int                              `json:"nomor_antrian"`
	Status       string                           `json:"status"`
	VaccineName  string                           `json:"vaccine_name"`
	CitizenId    string                           `json:"citizen_id"`
	Citizen      _responseCitizen.CitizenResponse `json:"citizen"`
	SessionId    string                           `json:"session_id"`
	Session      _responseSession.Session         `json:"session"`
	Date         string                           `json:"date"`
	SessionTime  string                           `json:"session_time"`
	CreatedAt    time.Time                        `json:"created_at"`
	UpdatedAt    time.Time                        `json:"updated_at"`
}

func FromDomainBookingCitizen(data booking.Domain) *BookingCitizen {
	return &BookingCitizen{
		BookingID:    data.Id,
		NomorAntrian: data.NomorAntrian,
		Status:       data.Status,
		VaccineName:  data.Session.Vaccine.NamaVaksin,
		CitizenId:    data.CitizenId,
		Citizen:      *_responseCitizen.FromDomain(data.Citizen),
		SessionId:    data.SessionId,
		Session:      *_responseSession.FromDomain(data.Session),
		Date:         data.Date,
		SessionTime:  data.SessionTime,
		CreatedAt:    data.CreatedAt,
		UpdatedAt:    data.UpdatedAt,
	}
}

func FromDomainOfArrayBookingCitizen(data []booking.Domain) *[]BookingCitizen {
	res := []BookingCitizen{}

	for _, val := range data {
		res = append(res, *FromDomainBookingCitizen(val))
	}
	return &res
}

type BookingSessionID struct {
	BookingID   string `json:"booking_id"`
	CitizenName string `json:"citizen_name"`
	Address     string `json:"address"`
	NIK         string `json:"nik"`
	NoTelp      string `json:"no_telp"`
	Status      string `json:"status"`
}

func FromDomainBookingSession(data booking.Domain) *BookingSessionID {
	return &BookingSessionID{
		BookingID:   data.Id,
		CitizenName: data.Citizen.FullName,
		Address:     data.Citizen.CurrentAddress.Alamat,
		NIK:         data.Citizen.Nik,
		NoTelp:      data.Citizen.NoHp,
		Status:      data.Status,
	}
}

func FromDomainOfArrayBookingSession(data []booking.Domain) *[]BookingSessionID {
	res := []BookingSessionID{}

	for _, val := range data {
		res = append(res, *FromDomainBookingSession(val))
	}
	return &res
}
