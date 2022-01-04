package currentAddress

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/helpers/nanoid"
	"strings"
)

type currentAddressService struct {
	currentAddressRepository Repository
}

func NewCurrentAddressService(currentAddressRepo Repository) Service {
	return &currentAddressService{
		currentAddressRepository: currentAddressRepo,
	}
}

func (service *currentAddressService) Create(data *Domain) (Domain, error) {
	data.Id, _ = nanoid.GenerateNanoId()
	dataAddress, err := service.currentAddressRepository.Create(data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}
	return dataAddress, nil
}

func (service *currentAddressService) GetByID(id string) (Domain, error) {
	dataAddress, err := service.currentAddressRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	return dataAddress, nil
}

func (service *currentAddressService) Update(id string, data *Domain) (Domain, error) {
	res, err := service.currentAddressRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	data.Id = res.Id
	dataAddress, err := service.currentAddressRepository.Update(id, data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataAddress, nil
}

func (service *currentAddressService) Delete(id string) (string, error) {
	_, err := service.currentAddressRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", businesses.ErrIDNotFound
		}
		return "", businesses.ErrInternalServer
	}

	if _, err := service.currentAddressRepository.Delete(id); err != nil {
		return "", businesses.ErrInternalServer
	}

	message := "current address success to deleted"
	return message, nil
}
