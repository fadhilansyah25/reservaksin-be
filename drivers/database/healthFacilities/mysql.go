package healthFacilities

import (
	"ca-reservaksin/businesses/healthFacilities"
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlHealthFacilitiesRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) healthFacilities.Repository {
	return &MysqlHealthFacilitiesRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlHealthFacilitiesRepository) Create(data *healthFacilities.Domain) (healthFacilities.Domain, error) {
	dataHealthFacilities := FromDomain(*data)

	err := mysqlRepo.Conn.Create(&dataHealthFacilities).Error
	if err != nil {
		return healthFacilities.Domain{}, err
	}

	if err := mysqlRepo.Conn.Preload(clause.Associations).Find(&dataHealthFacilities).Error; err != nil {
		return healthFacilities.Domain{}, err
	}

	return dataHealthFacilities.ToDomain(), nil
}

func (mysqlRepo *MysqlHealthFacilitiesRepository) GetByID(id string) (healthFacilities.Domain, error) {
	recFacilities := HealthFacilities{}

	if err := mysqlRepo.Conn.Preload(clause.Associations).First(&recFacilities, "id = ?", id).Error; err != nil {
		return healthFacilities.Domain{}, err
	}

	return recFacilities.ToDomain(), nil
}

func (mysqlRepo *MysqlHealthFacilitiesRepository) Update(id string, data *healthFacilities.Domain) (healthFacilities.Domain, error) {
	recFacilities := FromDomain(*data)
	err := mysqlRepo.Conn.Model(&recFacilities).Where("id = ?", id).Updates(&recFacilities).Error
	if err != nil {
		return healthFacilities.Domain{}, err
	}

	return recFacilities.ToDomain(), nil
}

func (mysqlRepo *MysqlHealthFacilitiesRepository) Delete(id string) (string, error) {
	recFacilities := HealthFacilities{}
	err := mysqlRepo.Conn.Delete(&recFacilities, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}

func (mysqlRepo *MysqlHealthFacilitiesRepository) GetByIdAdmin(id string) ([]healthFacilities.Domain, error) {
	dataFaskes := []HealthFacilities{}
	if err := mysqlRepo.Conn.Preload(clause.Associations).Find(&dataFaskes, "admin_id = ?", id).Error; err != nil {
		return []healthFacilities.Domain{}, err
	}
	if len(dataFaskes) == 0 {
		err := errors.New("data is empty")
		return []healthFacilities.Domain{}, err
	}
	return ToArrayOfDomain(dataFaskes), nil
}

func (mysqlRepo *MysqlHealthFacilitiesRepository) FetchAll() ([]healthFacilities.Domain, error) {
	dataFaskes := []HealthFacilities{}
	if err := mysqlRepo.Conn.Preload(clause.Associations).Find(&dataFaskes).Error; err != nil {
		return []healthFacilities.Domain{}, err
	}
	if len(dataFaskes) == 0 {
		err := errors.New("data is empty")
		return []healthFacilities.Domain{}, err
	}
	return ToArrayOfDomain(dataFaskes), nil
}
