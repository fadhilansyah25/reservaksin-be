package booking

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/helpers/nanoid"
)

type bookingsessionService struct {
	bookingRepository Repository
}

func NewBookingSessionService(repoBooking Repository) Service {
	return &bookingsessionService{
		bookingRepository: repoBooking,
	}
}

func (service *bookingsessionService) BookingSession(dataBooking *Domain) (Domain, error) {
	dataBooking.Id, _ = nanoid.GenerateNanoId()

	getQueueNumber, err := service.bookingRepository.GetBySessionID(dataBooking.SessionId)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}
	dataBooking.NomorAntrian = len(getQueueNumber) + 1

	booking, err := service.bookingRepository.Create(dataBooking)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return booking, nil
}
