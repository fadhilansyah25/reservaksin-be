package drivers

import (
	adminDomain "ca-reservaksin/businesses/admin"
	bookingDomain "ca-reservaksin/businesses/booking"
	citizenDomain "ca-reservaksin/businesses/citizen"
	currentAddressDomain "ca-reservaksin/businesses/currentAddress"
	healthFacilitiesDomain "ca-reservaksin/businesses/healthFacilities"
	sessionDomain "ca-reservaksin/businesses/session"
	vaccineDomain "ca-reservaksin/businesses/vaccine"
	adminDB "ca-reservaksin/drivers/database/admin"
	bookingDB "ca-reservaksin/drivers/database/booking"
	citizenDB "ca-reservaksin/drivers/database/citizen"
	currentAddressDB "ca-reservaksin/drivers/database/currentAddress"
	healthFacilitiesDB "ca-reservaksin/drivers/database/healthFacilities"
	sessionDB "ca-reservaksin/drivers/database/session"
	vaccineDB "ca-reservaksin/drivers/database/vaccine"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlRepository(conn)
}

func NewVaccineRepository(conn *gorm.DB) vaccineDomain.Repository {
	return vaccineDB.NewMysqlRepository(conn)
}

func NewCurrentAddressRepository(conn *gorm.DB) currentAddressDomain.Repository {
	return currentAddressDB.NewMysqlRepository(conn)
}

func NewHealthFacilitiesRepository(conn *gorm.DB) healthFacilitiesDomain.Repository {
	return healthFacilitiesDB.NewMysqlRepository(conn)
}

func NewSessionRepository(conn *gorm.DB) sessionDomain.Repository {
	return sessionDB.NewMysqlRepository(conn)
}

func NewCitizenRepository(conn *gorm.DB) citizenDomain.Repository {
	return citizenDB.NewMysqlRepository(conn)
}

func NewBookingRepository(conn *gorm.DB) bookingDomain.Repository {
	return bookingDB.NewMysqlRepository(conn)
}
