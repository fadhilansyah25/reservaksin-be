package drivers

import (
	adminDomain "ca-reservaksin/businesses/admin"
	vaccineDomain "ca-reservaksin/businesses/vaccine"
	adminDB "ca-reservaksin/drivers/database/admin"
	vaccineDB "ca-reservaksin/drivers/database/vaccine"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlRepository(conn)
}

func NewVaccineRepository(conn *gorm.DB) vaccineDomain.Repository {
	return vaccineDB.NewMysqlRepository(conn)
}
