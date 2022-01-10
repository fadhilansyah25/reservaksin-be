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
	recCitizen := fromDomain(*dataCitizen)

	err := mysqlrepo.Conn.Create(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.toDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByEmail(email string) (citizen.Domain, error) {
	recCitizen := Citizen{}
	err := mysqlrepo.Conn.First(&recCitizen, "email = ?", email).Error
	// err:=mysqlrepo.Conn.Create(&recCitizen).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.toDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByNIK(nik string) (citizen.Domain, error) {
	recCitizen := Citizen{}
	// s, err := mysqlrepo.Conn.Where("nik = ?", strconv.Atoi(nik)).First(&recCitizen).Error;
	// err := mysqlrepo.Conn.Where("nik = ?", nik).First(&recCitizen).Error
	err := mysqlrepo.Conn.First(&recCitizen, "nik = ?", nik).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recCitizen.toDomain(), nil
}

func (mysqlrepo *MysqlCitizenRepository) GetByNoKK(nokk string) ([]citizen.Domain, error) {
	recCitizen := []Citizen{}
	// s, err := mysqlrepo.Conn.Where("nik = ?", strconv.Atoi(nik)).First(&recCitizen).Error;
	err := mysqlrepo.Conn.Where("nokk = ?", nokk).First(&recCitizen).Error
	if err != nil {
		return []citizen.Domain{}, err
	}

	return ToArrayOfDomain(recCitizen), nil
}

func (mysqlRepo *MysqlCitizenRepository) Update(id string, data *citizen.Domain) (citizen.Domain, error) {
	recFacilities := fromDomain(*data)
	err := mysqlRepo.Conn.Model(&recFacilities).Where("id = ?", id).Updates(&recFacilities).Error
	if err != nil {
		return citizen.Domain{}, err
	}

	return recFacilities.toDomain(), nil
}

func (mysqlRepo *MysqlCitizenRepository) Delete(id string) (string, error) {
	recFacilities := Citizen{}
	err := mysqlRepo.Conn.Delete(&recFacilities, "id = ?", id).Error
	if err != nil {
		return "", err
	}

	return "", nil
}

func (mysqlRepo *MysqlCitizenRepository) GetByID(id string) (citizen.Domain, error) {
	recFacilities := Citizen{}

	if err := mysqlRepo.Conn.Preload(clause.Associations).First(&recFacilities, "id = ?", id).Error; err != nil {
		return citizen.Domain{}, err
	}

	return recFacilities.toDomain(), nil
}
