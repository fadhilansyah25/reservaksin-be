package session

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/helpers/nanoid"
	"strings"
)

type SessionService struct {
	SessionRepository Repository
	AddressRepository currentAddress.Repository
}

func NewSessionService(sessionRepo Repository, addressRepo currentAddress.Repository) Service {
	return &SessionService{
		SessionRepository: sessionRepo,
		AddressRepository: addressRepo,
	}
}

func (service *SessionService) Create(data *Domain) (Domain, error) {
	data.Id, _ = nanoid.GenerateNanoId()

	dataSession, err := service.SessionRepository.Create(data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataSession, nil
}

func (service *SessionService) GetByID(id string) (Domain, error) {
	dataSession, err := service.SessionRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	return dataSession, nil
}

func (service *SessionService) FetchAll() ([]Domain, error) {

	res, err := service.SessionRepository.FetchAll()
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	return res, nil
}

func (service *SessionService) GetByLatLong(lat float64, lng float64) ([]Result, error) {
	res, err := service.SessionRepository.GetByLatLong(lat, lng)
	if err != nil {
		return []Result{}, err
	}

	return res, nil
}
