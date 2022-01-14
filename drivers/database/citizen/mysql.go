package citizen

import (
	"ca-reservaksin/businesses/citizen"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlCitizenRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) citizen.Repository {
	return &MysqlCitizenRepository{
		Conn: conn,
	}
}

func (mysqlrepo *MysqlCitizenRepository) Register(dataCitizen *citizen.Domain) (citizen.Domain, error) {
	recCitizen := FromDomain(*dataCitizen)

	err := mysqlrepo.Conn.Create(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.ToDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByEmail(email string) (citizen.Domain, error) {
	recCitizen := Citizen{}
	err := mysqlrepo.Conn.First(&recCitizen, "email = ?", email).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.ToDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByNIK(nik string) (citizen.Domain, error) {
	recCitizen := Citizen{}
	err := mysqlrepo.Conn.First(&recCitizen, "nik = ?", nik).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.ToDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByNoKK(nokk string) ([]citizen.Domain, error) {
	recCitizen := []Citizen{}
	err := mysqlrepo.Conn.Where("nokk = ?", nokk).First(&recCitizen).Error
	if err != nil {
		return []citizen.Domain{}, err
	}

	return ToArrayOfDomain(recCitizen), nil
}

func (mysqlRepo *MysqlCitizenRepository) Update(id string, data *citizen.Domain) (citizen.Domain, error) {
	recCitizen := FromDomain(*data)
	err := mysqlRepo.Conn.Model(&recCitizen).Where("id = ?", id).Updates(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.ToDomain(), nil
}

func (mysqlRepo *MysqlCitizenRepository) Delete(id string) (string, error) {
	recCitizen := Citizen{}
	err := mysqlRepo.Conn.Delete(&recCitizen, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}

func (mysqlRepo *MysqlCitizenRepository) GetByID(id string) (citizen.Domain, error) {
	recCitizen := Citizen{}

	if err := mysqlRepo.Conn.Preload(clause.Associations).First(&recCitizen, "id = ?", id).Error; err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.ToDomain(), nil
}

func (mysqlRepo *MysqlCitizenRepository) GetByEmailOrNIK(email_or_string string) (citizen.Domain, error) {
	recCitizen := Citizen{}

	if err := mysqlRepo.Conn.Preload(clause.Associations).
		First(&recCitizen, "email = ? OR nik = ?", email_or_string, email_or_string).Error; err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.ToDomain(), nil
}
