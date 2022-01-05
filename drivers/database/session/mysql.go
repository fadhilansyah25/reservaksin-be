package session

import (
	"ca-reservaksin/businesses/session"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlSessionRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) session.Repository {
	return &MysqlSessionRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlSessionRepository) Create(data *session.Domain) (session.Domain, error) {
	dataSession := FromDomain(data)

	err := mysqlRepo.Conn.Create(&dataSession).Error
	if err != nil {
		return session.Domain{}, err
	}

	if err := mysqlRepo.Conn.Preload(clause.Associations).First(&dataSession).Error; err != nil {
		return session.Domain{}, err
	}

	return dataSession.ToDomain(), nil
}

// func (mysqlRepo *MysqlSessionRepository) FetchAll() ([]session.Domain, error) {

// }

// func (mysqlRepo *MysqlSessionRepository) Update(id string, data *session.Domain) (session.Domain, error) {

// }

func (mysqlRepo *MysqlSessionRepository) GetByID(id string) (session.Domain, error) {
	dataSession := Session{}

	err := mysqlRepo.Conn.Preload(clause.Associations).First(&dataSession, "id = ?", id).Error
	if err != nil {
		return session.Domain{}, err
	}

	return dataSession.ToDomain(), nil
}

// func (mysqlRepo *MysqlSessionRepository) Delete(id string) (string, error) {

// }
