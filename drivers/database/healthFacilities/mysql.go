package healthFacilities

import (
	"ca-reservaksin/businesses/healthFacilities"

	"gorm.io/gorm"
)

type MysqlAdressRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) healthFacilities.Repository {
	return &MysqlAdressRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlAdressRepository) Create(data *healthFacilities.Domain) (healthFacilities.Domain, error) {
	dataHealthFacilities := FromDomain(*data)

	err := mysqlRepo.Conn.Create(&dataHealthFacilities).Error
	if err != nil {
		return healthFacilities.Domain{}, err
	}

	return dataHealthFacilities.ToDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) GetByID(id string) (healthFacilities.Domain, error) {
	recFacilities := HealthFacilities{}

	err := mysqlRepo.Conn.Where("id = ?", id).First(&recFacilities).Error
	if err != nil {
		return healthFacilities.Domain{}, err
	}

	return recFacilities.ToDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) Update(id string, data *healthFacilities.Domain) (healthFacilities.Domain, error) {
	recFacilities := FromDomain(*data)
	err := mysqlRepo.Conn.Model(&recFacilities).Where("id = ?", id).Updates(&recFacilities).Error
	if err != nil {
		return healthFacilities.Domain{}, err
	}

	return recFacilities.ToDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) Delete(id string) (string, error) {
	recFacilities := HealthFacilities{}
	err := mysqlRepo.Conn.Delete(&recFacilities, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}
