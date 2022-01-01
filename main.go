package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_middlewares "ca-reservaksin/app/middlewares"
	"ca-reservaksin/app/routes"
	_driverFactory "ca-reservaksin/drivers"
	_dbDriver "ca-reservaksin/drivers/mysql"

	_adminService "ca-reservaksin/businesses/admin"
	_adminController "ca-reservaksin/controllers/admin"
	_AdminRepo "ca-reservaksin/drivers/database/admin"
)

func init() {
	viper.SetConfigFile(`app/config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug Mode")
	}
}

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_AdminRepo.Admin{},
	)
}

func main() {
	configDB := _dbDriver.ConfigDB{
		DB_Username: viper.GetString(`database.user`),
		DB_Password: viper.GetString(`database.pass`),
		DB_Host:     viper.GetString(`database.host`),
		DB_Port:     viper.GetString(`database.port`),
		DB_Database: viper.GetString(`database.name`),
	}

	db := configDB.InitialDB()
	dbMigrate(db)

	configJWT := _middlewares.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	adminRepo := _driverFactory.NewAdminRepository(db)
	adminService := _adminService.NewAdminService(adminRepo, &configJWT)
	adminCtrl := _adminController.NewAdminController(adminService)

	routesInit := routes.ControllerList{
		JwtMiddleware:   configJWT.Init(),
		AdminController: *adminCtrl,
	}
	e := echo.New()
	routesInit.RoutesRegister(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}
