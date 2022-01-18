package request

import "ca-reservaksin/businesses/booking"

type Booking struct {
	CitizenId   string `json:"citizen_id" validate:"required"`
	SessionId   string `json:"session_id" validate:"required"`
	Date        string `json:"date" validate:"required"`
	SessionTime string `json:"session_time" validate:"required"`
}

func (req *Booking) ToDomain() *booking.Domain {
	return &booking.Domain{
		CitizenId:   req.CitizenId,
		SessionId:   req.SessionId,
		Date:        req.Date,
		SessionTime: req.SessionTime,
	}
}
