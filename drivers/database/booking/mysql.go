package booking

import (
	"ca-reservaksin/businesses/booking"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlBookingRepository struct {
	Conn *gorm.DB
}

func NewMysqlBookingRepository(conn *gorm.DB) booking.Repository {
	return &MysqlBookingRepository{
		Conn: conn,
	}
}

func (mysqlRepo *MysqlBookingRepository) Create(data *booking.Domain) (booking.Domain, error) {
	recBooking := FromDomain(data)

	err := mysqlRepo.Conn.Create(&recBooking).Error
	if err != nil {
		return booking.Domain{}, err
	}

	mysqlRepo.Conn.Preload("Session.HealthFacilites").Preload("Session.Vaccine").Preload(clause.Associations).Find(&recBooking)

	return recBooking.ToDomain(), nil
}

func (mysqlRepo *MysqlBookingRepository) GetBySessionID(sessionID string) ([]booking.Domain, error) {
	recBooking := []Booking{}

	err := mysqlRepo.Conn.Preload("Citizen.CurrentAddress").Preload(clause.Associations).Find(&recBooking, "session_id = ?", sessionID).Error
	if err != nil {
		return []booking.Domain{}, err
	}

	return ToArrayOfDomain(recBooking), nil
}

func (mysqlRepo *MysqlBookingRepository) GetByStatus(status string) ([]booking.Domain, error) {
	recBooking := []Booking{}

	err := mysqlRepo.Conn.Find(&recBooking, "session_id = ?", status).Error
	if err != nil {
		return []booking.Domain{}, err
	}

	return ToArrayOfDomain(recBooking), nil
}

func (mysqlRepo *MysqlBookingRepository) GetByID(id string) (booking.Domain, error) {
	recBooking := Booking{}

	err := mysqlRepo.Conn.Find(&recBooking, "session_id = ?", id).Error
	if err != nil {
		return booking.Domain{}, err
	}

	return recBooking.ToDomain(), nil
}

func (mysqlRepo *MysqlBookingRepository) GetByCitizenID(sessionID string) ([]booking.Domain, error) {
	recBooking := []Booking{}

	err := mysqlRepo.Conn.Preload("Session.HealthFacilites").Preload("Session.Vaccine").Preload(clause.Associations).Find(&recBooking, "citizen_id = ?", sessionID).Error
	if err != nil {
		return []booking.Domain{}, err
	}

	return ToArrayOfDomain(recBooking), nil
}
