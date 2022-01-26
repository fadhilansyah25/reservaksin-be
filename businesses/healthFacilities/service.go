package healthFacilities

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

func (service *HealthFacilitiesService) Create(data *Domain) (Domain, error) {
	data.CurrentAddress.Id, _ = nanoid.GenerateNanoId()
	newAddress, _ := service.AddressRepository.Create(&data.CurrentAddress)

	data.ID, _ = nanoid.GenerateNanoId()
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

	data.ID = existed.ID
	data.CurrentAddressID = existed.CurrentAddressID
	data.CurrentAddress.Id = existed.CurrentAddressID

	if _, err := service.AddressRepository.Update(data.CurrentAddress.Id, &data.CurrentAddress); err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

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

	if _, err := service.AddressRepository.Delete(existed.CurrentAddressID); err != nil {
		return "", businesses.ErrInternalServer
	}

	message := fmt.Sprintf("health facilities %s success to deleted", existed.NameFacilites)
	return message, nil
}

func (service *HealthFacilitiesService) GetByIdAdmin(id string) ([]Domain, error) {
	dataFaskes, err := service.FacilitiesRepository.GetByIdAdmin(id)
	if err != nil {
		if strings.Contains(err.Error(), "empty") {
			return []Domain{}, err
		}
		return []Domain{}, err
	}
	return dataFaskes, nil
}

func (service *HealthFacilitiesService) FetchAll() ([]Domain, error) {
	dataFaskes, err := service.FacilitiesRepository.FetchAll()
	if err != nil {
		if strings.Contains(err.Error(), "empty") {
			return []Domain{}, err
		}
		return []Domain{}, err
	}
	return dataFaskes, nil
}
