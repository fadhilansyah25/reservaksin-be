package response

import (
	"ca-reservaksin/businesses/booking"
	"time"
)

type BookingCitizen struct {
	Id           string `json:"id"`
	NomorAntrian int    `json:"nomor_antrian"`
	Status       string `json:"status"`
	VaccineName  string `json:"vaccine_name"`
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

func FromDomainBookingCitizen(data booking.Domain) *BookingCitizen {
	return &BookingCitizen{
		Id:           data.Id,
		NomorAntrian: data.NomorAntrian,
		Status:       data.Status,
		VaccineName:  data.Session.Vaccine.NamaVaksin,
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

func FromDomainOfArrayBookingCitizen(data []booking.Domain) *[]BookingCitizen {
	res := []BookingCitizen{}

	for _, val := range data {
		res = append(res, *FromDomainBookingCitizen(val))
	}
	return &res
}

type BookingSessionID struct {
	CitizenName string `json:"citizen_name"`
	Address     string `json:"address"`
	NIK         string `json:"nik"`
	NoTelp      string `json:"no_telp"`
	Status      string `json:"status"`
}

func FromDomainBookingSession(data booking.Domain) *BookingSessionID {
	return &BookingSessionID{
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
