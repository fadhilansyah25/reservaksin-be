package vaccine

import (
	"ca-reservaksin/businesses/vaccine"
	"errors"

	"gorm.io/gorm"
)

type mysqlVaccineRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) vaccine.Repository {
	return &mysqlVaccineRepository{
		Conn: conn,
	}
}

func (mysqlRepo *mysqlVaccineRepository) Create(data *vaccine.Domain) (vaccine.Domain, error) {
	recVaccine := fromDomain(data)

	err := mysqlRepo.Conn.Create(&recVaccine).Error
	if err != nil {
		return vaccine.Domain{}, err
	}

	return recVaccine.ToDomain(), nil
}

func (mysqlRepo *mysqlVaccineRepository) Update(id string, data *vaccine.Domain) (vaccine.Domain, error) {
	recVaccine := fromDomain(data)

	err := mysqlRepo.Conn.Save(&recVaccine).Error
	if err != nil {
		return vaccine.Domain{}, err
	}

	return recVaccine.ToDomain(), nil
}

func (mysqlRepo *mysqlVaccineRepository) Delete(id string) (string, error) {
	recVaccine := Vaccine{}
	err := mysqlRepo.Conn.Delete(&recVaccine, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}

func (mysqlRepo *mysqlVaccineRepository) GetByID(id string) (vaccine.Domain, error) {
	recVaccine := Vaccine{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&recVaccine).Error
	if err != nil {
		return vaccine.Domain{}, err
	}

	return recVaccine.ToDomain(), nil
}

func (mysqlRepo *mysqlVaccineRepository) FetchAll() ([]vaccine.Domain, error) {
	recVaccine := []Vaccine{}

	results := mysqlRepo.Conn.Find(&recVaccine)
	if results.Error != nil {
		return []vaccine.Domain{}, results.Error
	}

	if len(recVaccine) == 0 {
		err := errors.New("data is empty")
		return []vaccine.Domain{}, err
	}

	return ToArrayOfDomain(recVaccine), nil
}

func (mysqlRepo *mysqlVaccineRepository) GetByAdminID(adminID string) ([]vaccine.Domain, error) {
	dataVaccine := []Vaccine{}
	if err := mysqlRepo.Conn.Find(&dataVaccine, "admin_id = ?", adminID).Error; err != nil {
		return []vaccine.Domain{}, err
	}
	if len(dataVaccine) == 0 {
		err := errors.New("data is empty")
		return []vaccine.Domain{}, err
	}
	return ToArrayOfDomain(dataVaccine), nil
}
