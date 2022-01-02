package drivers

import (
	adminDomain "ca-reservaksin/businesses/admin"
	currentAddressDomain "ca-reservaksin/businesses/currentAddress"
	vaccineDomain "ca-reservaksin/businesses/vaccine"
	adminDB "ca-reservaksin/drivers/database/admin"
	currentAddressDB "ca-reservaksin/drivers/database/currentAddress"
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
