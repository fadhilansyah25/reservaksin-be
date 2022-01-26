package booking

import (
	"ca-reservaksin/businesses/booking"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type MysqlBookingRepository struct {
	Conn *gorm.DB
}

func NewMysqlRepository(conn *gorm.DB) booking.Repository {
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

	mysqlRepo.Conn.Preload("Session.HealthFacilites").
		Preload("Session.HealthFacilites.CurrentAddress").
		Preload("Citizen.CurrentAddress").
		Preload("Session.Vaccine").
		Preload(clause.Associations).Find(&recBooking)

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

	err := mysqlRepo.Conn.Preload("Session.HealthFacilites").
		Preload("Session.HealthFacilites.CurrentAddress").
		Preload("Citizen.CurrentAddress").
		Preload("Session.Vaccine").
		Preload(clause.Associations).
		Find(&recBooking, "id = ?", id).Error
	if err != nil {
		return booking.Domain{}, err
	}

	return recBooking.ToDomain(), nil
}

func (mysqlRepo *MysqlBookingRepository) GetByCitizenID(citizenID string) ([]booking.Domain, error) {
	recBooking := []Booking{}

	err := mysqlRepo.Conn.Preload("Session.HealthFacilites").
		Preload("Session.HealthFacilites.CurrentAddress").
		Preload("Citizen.CurrentAddress").
		Preload("Session.Vaccine").
		Preload(clause.Associations).
		Find(&recBooking, "citizen_id = ?", citizenID).Error
	if err != nil {
		return []booking.Domain{}, err
	}

	return ToArrayOfDomain(recBooking), nil
}

func (mysqlRepo *MysqlBookingRepository) GetByNoKK(noKK string) ([]booking.Domain, error) {
	recBooking := []Booking{}

	sqlQuery := `SELECT
			bookings.*
		FROM
			bookings
			INNER JOIN citizens ON bookings.citizen_id = citizens.id
			INNER JOIN current_addresses ON citizens.current_address_id = current_addresses.id
			INNER JOIN sessions ON bookings.session_id = sessions.id
			INNER JOIN health_facilities ON sessions.health_facilites_id = health_facilities.id
		WHERE
			no_kk = ?
		ORDER BY
		status DESC, created_at DESC`

	err := mysqlRepo.Conn.Raw(sqlQuery, noKK).Preload("Session.HealthFacilites").
		Preload("Session.HealthFacilites.CurrentAddress").
		Preload("Citizen.CurrentAddress").
		Preload("Session.Vaccine").
		Preload(clause.Associations).
		Find(&recBooking).Error
	if err != nil {
		return []booking.Domain{}, err
	}

	return ToArrayOfDomain(recBooking), nil
}

func (mysqlRepo *MysqlBookingRepository) UpdateStatusByID(id, status string) (booking.Domain, error) {
	recBooking := Booking{}

	err := mysqlRepo.Conn.Model(&recBooking).Where("id = ?", id).Update("status", status).Error
	if err != nil {
		return booking.Domain{}, err
	}

	mysqlRepo.Conn.Preload("Session.HealthFacilites").
		Preload("Session.HealthFacilites.CurrentAddress").
		Preload("Citizen.CurrentAddress").
		Preload("Session.Vaccine").
		Preload(clause.Associations).First(&recBooking, "id = ?", id)

	return recBooking.ToDomain(), nil
}
