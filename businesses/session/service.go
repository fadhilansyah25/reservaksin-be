package session

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/helpers/nanoid"
	"strings"
)

type SessionService struct {
	SessionRepository Repository
}

func NewSessionService(sessionRepo Repository) Service {
	return &SessionService{
		SessionRepository: sessionRepo,
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
