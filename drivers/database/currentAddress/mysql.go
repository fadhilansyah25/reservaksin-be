package currentAddress

import (
	"ca-reservaksin/businesses/currentAddress"

	"gorm.io/gorm"
)

type MysqlAdressRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) currentAddress.Repository {
	return &MysqlAdressRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlAdressRepository) Create(data *currentAddress.Domain) (currentAddress.Domain, error) {
	dataAddress := FromDomain(*data)
	err := mysqlRepo.Conn.Create(&dataAddress).Error
	if err != nil {
		return currentAddress.Domain{}, err
	}

	return dataAddress.ToDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) GetByID(id string) (currentAddress.Domain, error) {
	recAddress := CurrentAddress{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&recAddress).Error
	if err != nil {
		return currentAddress.Domain{}, err
	}

	return recAddress.ToDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) Update(id string, data *currentAddress.Domain) (currentAddress.Domain, error) {
	recAddress := FromDomain(*data)

	err := mysqlRepo.Conn.Model(&recAddress).Where("id = ?", id).Updates(&recAddress).Error
	if err != nil {
		return currentAddress.Domain{}, err
	}

	return recAddress.ToDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) Delete(id string) (string, error) {
	recAddress := CurrentAddress{}
	err := mysqlRepo.Conn.Delete(&recAddress, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}
