package request

import (
	"ca-reservaksin/businesses/session"
)

type Session struct {
	HealthFacilitesID string `json:"health_facilities_id"`
	NameSession       string `json:"name_session"`
	Capacity          int    `json:"capacity"`
	VaccineID         string `json:"vaccine_id"`
	Date              string `json:"date"`
	Tahap             string `json:"tahap"`
	StartSession      string `json:"start_session"`
	EndSession        string `json:"end_session"`
}

func (req *Session) ToDomain() *session.Domain {
	return &session.Domain{
		HealthFacilitesID: req.HealthFacilitesID,
		NameSession:       req.NameSession,
		Capacity:          req.Capacity,
		VaccineID:         req.VaccineID,
		Date:              req.Date,
		Tahap:             req.Tahap,
		StartSession:      req.StartSession,
		EndSession:        req.EndSession,
	}
}
