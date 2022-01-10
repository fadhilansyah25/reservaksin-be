package response

import (
	"ca-reservaksin/businesses/session"
	_responseHF "ca-reservaksin/controllers/healthFacilities/response"
	_responseVC "ca-reservaksin/controllers/vaccine/response"
	"time"
)

type Session struct {
	Id                string                       `json:"id"`
	HealthFacilitesID string                       `json:"health_facilities_id"`
	HealthFacilites   _responseHF.HealthFacilities `json:"health_facilities"`
	NameSession       string                       `json:"name_session"`
	Capacity          int                          `json:"capacity"`
	VaccineID         string                       `json:"vaccine_id"`
	Vaccine           _responseVC.Vaccine          `json:"vaccine"`
	Date              string                       `json:"date"`
	Tahap             string                       `json:"tahap"`
	StartSession      string                       `json:"start_session"`
	EndSession        string                       `json:"end_session"`
	CreatedAt         time.Time                    `json:"created_at"`
	UpdatedAt         time.Time                    `json:"updated_at"`
}

func FromDomain(data session.Domain) *Session {
	return &Session{
		Id:                data.Id,
		HealthFacilitesID: data.HealthFacilitesID,
		HealthFacilites:   *_responseHF.FromDomain(data.HealthFacilites),
		NameSession:       data.NameSession,
		Capacity:          data.Capacity,
		VaccineID:         data.VaccineID,
		Vaccine:           *_responseVC.FromDomain(data.Vaccine),
		Date:              data.Date,
		Tahap:             data.Tahap,
		StartSession:      data.StartSession,
		EndSession:        data.EndSession,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}

func FromDomainArray(domain []session.Domain) *[]Session {
	res := []Session{}

	for _, val := range domain {
		res = append(res, *FromDomain(val))
	}
	return &res
}

// Response Session With Distance
type Result struct {
	Session  `json:"session"`
	Distance float64 `json:"distance"`
}

func FromDomainResult(domain session.SessionDistance) *Result {
	return &Result{
		Session:  *FromDomain(domain.Session),
		Distance: domain.Distance,
	}
}

func FromDomainArrayResult(domain []session.SessionDistance) *[]Result {
	res := []Result{}

	for _, val := range domain {
		res = append(res, *FromDomainResult(val))
	}
	return &res
}

// Response Session Simple
type SimpleResSession struct {
	Id                string    `json:"id"`
	HealthFacilitesID string    `json:"health_facilities_id"`
	HealthFacilites   string    `json:"health_facilities"`
	NameSession       string    `json:"name_session"`
	Capacity          int       `json:"capacity"`
	VaccineID         string    `json:"vaccine_id"`
	Vaccine           string    `json:"vaccine"`
	Date              string    `json:"date"`
	Tahap             string    `json:"tahap"`
	StartSession      string    `json:"start_session"`
	EndSession        string    `json:"end_session"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func FromDomainSimpleRes(data session.Domain) *SimpleResSession {
	return &SimpleResSession{
		Id:                data.Id,
		HealthFacilitesID: data.HealthFacilitesID,
		HealthFacilites:   data.HealthFacilites.NameFacilites,
		NameSession:       data.NameSession,
		Capacity:          data.Capacity,
		VaccineID:         data.VaccineID,
		Vaccine:           data.Vaccine.NamaVaksin,
		Date:              data.Date,
		Tahap:             data.Tahap,
		StartSession:      data.StartSession,
		EndSession:        data.EndSession,
		CreatedAt:         data.CreatedAt,
		UpdatedAt:         data.UpdatedAt,
	}
}

func FromDomainArraySimpleRes(domain []session.Domain) *[]SimpleResSession {
	res := []SimpleResSession{}

	for _, val := range domain {
		res = append(res, *FromDomainSimpleRes(val))
	}
	return &res
}
