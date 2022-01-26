package admin

import (
	"ca-reservaksin/businesses/admin"

	"gorm.io/gorm"
)

type MysqlAdminRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) admin.Repository {
	return &MysqlAdminRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlAdminRepository) Register(dataAdmin *admin.Domain) (admin.Domain, error) {
	recAdmin := FromDomain(*dataAdmin)

	err := mysqlRepo.Conn.Create(&recAdmin).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return recAdmin.ToDomain(), nil
}

func (mysqlRepo *MysqlAdminRepository) GetByUsername(username string) (admin.Domain, error) {
	recAdmin := Admin{}
	err := mysqlRepo.Conn.First(&recAdmin, "username = ?", username).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return recAdmin.ToDomain(), nil
}

func (mysqlRepo *MysqlAdminRepository) GetByID(id string) (admin.Domain, error) {
	recAdmin := Admin{}
	err := mysqlRepo.Conn.Where("id = ?", id).First(&recAdmin).Error
	if err != nil {
		return admin.Domain{}, err
	}

	return recAdmin.ToDomain(), nil
}
