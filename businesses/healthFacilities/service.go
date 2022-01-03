package healthfacilities

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/helpers/nanoid"
	"fmt"
	"strings"
)

type HealthFacilitiesService struct {
	FacilitiesRepository Repository
	AddressRepository    currentAddress.Repository
}

func NewHealthFacilitiesService(facilitiesRepo Repository, addressRepo currentAddress.Repository) Service {
	return &HealthFacilitiesService{
		FacilitiesRepository: facilitiesRepo,
		AddressRepository:    addressRepo,
	}
}

func (service *HealthFacilitiesService) Create(data *Domain, address *currentAddress.Domain) (Domain, error) {
	address.Id, _ = nanoid.GenerateNanoId()
	newAddress, _ := service.AddressRepository.Create(address)

	data.Id, _ = nanoid.GenerateNanoId()
	data.CurrentAddressID = newAddress.Id

	dataHealthFacilities, err := service.FacilitiesRepository.Create(data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataHealthFacilities, nil
}

func (service *HealthFacilitiesService) GetByID(id string) (Domain, error) {
	dataHealthFacilities, err := service.FacilitiesRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	return dataHealthFacilities, nil
}

func (service *HealthFacilitiesService) Update(id string, data *Domain) (Domain, error) {
	existed, err := service.FacilitiesRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	data.Id = existed.Id
	dataHealthFacilities, err := service.FacilitiesRepository.Update(id, data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataHealthFacilities, nil
}

func (service *HealthFacilitiesService) Delete(id string) (string, error) {
	existed, err := service.FacilitiesRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", businesses.ErrIDNotFound
		}
		return "", businesses.ErrInternalServer
	}

	if _, err := service.FacilitiesRepository.Delete(id); err != nil {
		return "", businesses.ErrInternalServer
	}

	message := fmt.Sprintf("health facilities %s success to deleted", existed.NameFacilites)
	return message, nil
}
