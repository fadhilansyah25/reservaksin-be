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
	jwtAuth           *middlewares.ConfigJWT
}

func NewCitizenService(repoCitizen Repository, addressRepo currentAddress.Repository, jwtauth *middlewares.ConfigJWT) Service {
	return &citizenService{
		citizenRepository: repoCitizen,
		AddressRepository: addressRepo,
		jwtAuth:           jwtauth,
	}
}

func (repo *citizenService) Register(citizenDomain *Domain) (Domain, error) {
	existedCitizen, _ := repo.citizenRepository.GetByNIK(citizenDomain.Nik)
	if existedCitizen != (Domain{}) {
		return Domain{}, businesses.ErrDuplicateData
	}

	citizenDomain.CurrentAddress.Id, _ = nanoid.GenerateNanoId()
	newAddress, _ := repo.AddressRepository.Create(&citizenDomain.CurrentAddress)
	citizenDomain.CurrentAddressID = newAddress.Id

	citizenDomain.Id, _ = nanoid.GenerateNanoId()
	if citizenDomain.Password != "" {
		hashedPassword := encrypt.HashAndSalt([]byte(citizenDomain.Password))
		citizenDomain.Password = hashedPassword
	}

	result, err := repo.citizenRepository.Register(citizenDomain)
	if err != nil {
		return Domain{}, err
	}
	return result, nil
}

func (repo *citizenService) Login(email_or_nik, password string) (Domain, string, error) {

	if strings.TrimSpace(email_or_nik) == "" {
		return Domain{}, "", businesses.ErrEmailOrNIKNotFound
	}

	citizenDomain, err := repo.citizenRepository.GetByEmailOrNIK(email_or_nik)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, "", businesses.ErrEmailOrNIKNotFound
		}
		return Domain{}, "", businesses.ErrInternalServer
	}

	if !encrypt.ValidateHash(password, citizenDomain.Password) {
		return Domain{}, "", businesses.ErrIncorrectPassword
	}

	token := repo.jwtAuth.GenerateToken(citizenDomain.Id)
	return citizenDomain, token, nil
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

	message := fmt.Sprintf("account %s success to deleted", existed.FullName)
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

	data.Id = existed.Id
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

func (service *citizenService) GetByID(id string) (Domain, error) {
	resCitizen, err := service.citizenRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	return resCitizen, nil
}

func (service *citizenService) GetByAdminID(adminID string) ([]Domain, error) {
	resCitizen, err := service.citizenRepository.GetByAdminID(adminID)
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	return resCitizen, nil
}

func (service *citizenService) GetByNoKK(noKK string) ([]Domain, error) {
	resCitizen, err := service.citizenRepository.GetByNoKK(noKK)
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	return resCitizen, nil
}
