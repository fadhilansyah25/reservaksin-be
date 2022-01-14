package vaccine

import (
	"ca-reservaksin/businesses"
	"ca-reservaksin/helpers/nanoid"
	"fmt"
	"strings"
)

type vaccineService struct {
	vaccineRepository Repository
}

func NewVaccineService(vaccineRepo Repository) Service {
	return &vaccineService{
		vaccineRepository: vaccineRepo,
	}
}

func (service *vaccineService) Create(data *Domain) (Domain, error) {
	data.Id, _ = nanoid.GenerateNanoId()
	dataVaccine, err := service.vaccineRepository.Create(data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}
	return dataVaccine, nil
}

func (service *vaccineService) Update(id string, data *Domain) (Domain, error) {
	existed, err := service.vaccineRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	data.Id = existed.Id

	dataVaccine, err := service.vaccineRepository.Update(id, data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataVaccine, nil
}

func (service *vaccineService) Delete(id string) (string, error) {
	existed, err := service.vaccineRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", businesses.ErrIDNotFound
		}
		return "", businesses.ErrInternalServer
	}

	if _, err := service.vaccineRepository.Delete(id); err != nil {
		return "", err
	}

	message := fmt.Sprintf("vaccine %s success to deleted", existed.NamaVaksin)
	return message, nil
}

func (service *vaccineService) GetByID(id string) (Domain, error) {
	data, err := service.vaccineRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, businesses.ErrInternalServer
	}

	return data, nil
}

func (service *vaccineService) FetchAll() ([]Domain, error) {
	data, err := service.vaccineRepository.FetchAll()
	if err != nil {
		return []Domain{}, businesses.ErrInternalServer
	}

	if err != nil {
		if strings.Contains(err.Error(), "empty") {
			return []Domain{}, err
		}
		return []Domain{}, err
	}

	return data, nil
}

func (service *vaccineService) GetByAdminID(adminID string) ([]Domain, error) {
	dataVaccine, err := service.vaccineRepository.GetByAdminID(adminID)
	if err != nil {
		if strings.Contains(err.Error(), "empty") {
			return []Domain{}, err
		}
		return []Domain{}, err
	}
	return dataVaccine, nil
}
