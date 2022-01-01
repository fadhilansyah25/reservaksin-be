package drivers

import (
	adminDomain "ca-reservaksin/businesses/admin"
	adminDB "ca-reservaksin/drivers/database/admin"

	"gorm.io/gorm"
)

func NewAdminRepository(conn *gorm.DB) adminDomain.Repository {
	return adminDB.NewMysqlRepository(conn)
}
