package request

import "ca-reservaksin/businesses/booking"

type Booking struct {
	CitizenId   string `json:"citizen_id"`
	SessionId   string `json:"session_id"`
	Date        string `json:"date"`
	SessionTime string `json:"session_time"`
}

func (req *Booking) ToDomain() *booking.Domain {
	return &booking.Domain{
		CitizenId:   req.CitizenId,
		SessionId:   req.SessionId,
		Date:        req.Date,
		SessionTime: req.SessionTime,
	}
}
