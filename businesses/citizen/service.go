package citizen

import (
	"ca-reservaksin/app/middlewares"
	"ca-reservaksin/businesses"
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/helpers/encrypt"
	"ca-reservaksin/helpers/nanoid"
	"fmt"
	"strings"
)

type citizenService struct {
	citizenRepository Repository
	AddressRepository currentAddress.Repository
	// contextTimeout    time.Duration
	jwtAuth *middlewares.ConfigJWT
}

func NewCitizenService(repoCitizen Repository, jwtauth *middlewares.ConfigJWT) Service {
	return &citizenService{
		citizenRepository: repoCitizen,
		// contextTimeout:    timeout,
		jwtAuth: jwtauth,
	}
}

func (repo *citizenService) Register(citizenDomain *Domain) (Domain, error) {
	existedCitizen, err := repo.citizenRepository.GetByNIK(citizenDomain.Nik)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return Domain{}, err
		}
	}

	citizenDomain.ID, err = nanoid.GenerateNanoId()
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	if existedCitizen != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	citizenDomain.Password = encrypt.HashAndSalt([]byte(citizenDomain.Password))
	result, err := repo.citizenRepository.Register(citizenDomain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (repo *citizenService) LoginByEmail(email, password string) (string, error) {

	if strings.TrimSpace(email) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	citizenDomain, err := repo.citizenRepository.GetByEmail(email)

	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, citizenDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := repo.jwtAuth.GenerateToken(citizenDomain.ID)
	return token, nil

}

func (repo *citizenService) LoginByNIK(nik, password string) (string, error) {

	if strings.TrimSpace(nik) == "" && strings.TrimSpace(password) == "" {
		return "", businesses.ErrEmailPasswordNotFound
	}

	citizenDomain, err := repo.citizenRepository.GetByNIK(nik)

	if err != nil {
		return "", err
	}

	if !encrypt.ValidateHash(password, citizenDomain.Password) {
		return "", businesses.ErrInternalServer
	}

	token := repo.jwtAuth.GenerateToken(citizenDomain.ID)
	return token, nil

}

func (service *citizenService) Delete(id string) (string, error) {
	existed, err := service.citizenRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", businesses.ErrIDNotFound
		}
		return "", businesses.ErrInternalServer
	}

	if _, err := service.citizenRepository.Delete(id); err != nil {
		return "", businesses.ErrInternalServer
	}

	if _, err := service.AddressRepository.Delete(existed.CurrentAddressID); err != nil {
		return "", businesses.ErrInternalServer
	}

	message := fmt.Sprintf("account %s success to deleted", existed.Username)
	return message, nil
}

func (service *citizenService) Update(id string, data *Domain) (Domain, error) {
	existed, err := service.citizenRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	data.ID = existed.ID
	data.CurrentAddressID = existed.CurrentAddressID
	data.CurrentAddress.Id = existed.CurrentAddressID

	if _, err := service.AddressRepository.Update(data.CurrentAddress.Id, &data.CurrentAddress); err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	dataHealthFacilities, err := service.citizenRepository.Update(id, data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataHealthFacilities, nil
}
