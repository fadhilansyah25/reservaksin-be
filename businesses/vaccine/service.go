package vaccine

import (
	"ca-reservaksin/businesses"
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
	dataVaccine, err := service.vaccineRepository.Create(data)
	if err != nil {
		return Domain{}, err
	}
	return dataVaccine, err
}

func (service *vaccineService) Update(id int, data *Domain) (Domain, error) {
	res, err := service.vaccineRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return Domain{}, businesses.ErrIDNotFound
		}
		return Domain{}, err
	}

	data.Id = res.Id

	dataVaccine, err := service.vaccineRepository.Update(id, data)
	if err != nil {
		return Domain{}, businesses.ErrInternalServer
	}

	return dataVaccine, nil
}

func (service *vaccineService) Delete(id int) (string, error) {
	existed, err := service.vaccineRepository.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return "", businesses.ErrIDNotFound
		}
		return "", businesses.ErrIDNotFound
	}

	if _, err := service.vaccineRepository.Delete(id); err != nil {
		return "", err
	}

	message := fmt.Sprintf("vaccine %s success to deleted", existed.NamaVaksin)
	return message, nil
}

func (service *vaccineService) GetByID(id int) (Domain, error) {
	data, err := service.vaccineRepository.GetByID(id)
	if err != nil {
		return Domain{}, businesses.ErrIDNotFound
	}

	return data, nil
}

func (service *vaccineService) FetchAll() ([]Domain, error) {
	data, err := service.vaccineRepository.FetchAll()
	if err != nil {
		return []Domain{}, err
	}

	return data, err
}
