package session

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/helpers/nanoid"
	"fmt"
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

	resSession, err := service.SessionRepository.FetchAll()
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	return resSession, nil
}

func (service *SessionService) GetByLatLong(lat float64, lng float64) ([]SessionDistance, error) {
	resultSession, err := service.SessionRepository.GetByLatLong(lat, lng)
	if err != nil {
		return []SessionDistance{}, businesses.ErrInternalServer
	}

	return resultSession, nil
}

func (service *SessionService) Update(id string, data *Domain) (Domain, error) {
	existed, err := service.SessionRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	data.Id = existed.Id
	dataSession, err := service.SessionRepository.Update(id, data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataSession, nil
}

func (service *SessionService) Delete(id string) (string, error) {
	existed, err := service.SessionRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", businesses.ErrIDNotFound
		}
		return "", businesses.ErrInternalServer
	}

	if _, err := service.SessionRepository.Delete(id); err != nil {
		return "", businesses.ErrInternalServer
	}

	message := fmt.Sprintf("session %s success to deleted", existed.NameSession)
	return message, nil
}

func (service *SessionService) FetchByHistory(adminID, history string) ([]Domain, error) {

	param := strings.ToLower(history)
	resSession, err := service.SessionRepository.FetchByHistory(adminID, param)
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	return resSession, nil
}

func (service *SessionService) FetchAllByAdminID(adminID string) ([]Domain, error) {
	resSession, err := service.SessionRepository.FetchAllByAdminID(adminID)
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	return resSession, nil
}
