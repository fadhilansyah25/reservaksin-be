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
	dataAddress := fromDomain(*data)
	err := mysqlRepo.Conn.Create(&dataAddress).Error
	if err != nil {
		return currentAddress.Domain{}, err
	}

	return dataAddress.toDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) GetByID(id string) (currentAddress.Domain, error) {
	recAddress := CurrentAddress{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&recAddress).Error
	if err != nil {
		return currentAddress.Domain{}, err
	}

	return recAddress.toDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) Update(id string, data *currentAddress.Domain) (currentAddress.Domain, error) {
	recAddress := fromDomain(*data)

	err := mysqlRepo.Conn.Save(&recAddress).Error
	if err != nil {
		return currentAddress.Domain{}, err
	}

	return recAddress.toDomain(), nil
}

func (mysqlRepo *MysqlAdressRepository) Delete(id string) (string, error) {
	recAddress := CurrentAddress{}
	err := mysqlRepo.Conn.Delete(&recAddress, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}
