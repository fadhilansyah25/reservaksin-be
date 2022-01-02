package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_middlewares "ca-reservaksin/app/middlewares"
	_routes "ca-reservaksin/app/routes"
	_driverFactory "ca-reservaksin/drivers"
	_dbDriver "ca-reservaksin/drivers/mysql"

	_adminService "ca-reservaksin/businesses/admin"
	_currentAddressService "ca-reservaksin/businesses/currentAddress"
	_vaccineService "ca-reservaksin/businesses/vaccine"
	_adminController "ca-reservaksin/controllers/admin"
	_currentAddressController "ca-reservaksin/controllers/currentAddress"
	_vaccineController "ca-reservaksin/controllers/vaccine"
	_AdminRepo "ca-reservaksin/drivers/database/admin"
	_currentAddressRepo "ca-reservaksin/drivers/database/currentAddress"
	_VaccineRepo "ca-reservaksin/drivers/database/vaccine"
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
		&_VaccineRepo.Vaccine{},
		&_currentAddressRepo.CurrentAddress{},
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

	vaccineRepo := _driverFactory.NewVaccineRepository(db)
	vaccineService := _vaccineService.NewVaccineService(vaccineRepo)
	vaccineCtrl := _vaccineController.NewVaccineController(vaccineService)

	currentAddressRepo := _driverFactory.NewCurrentAddressRepository(db)
	currentAddressService := _currentAddressService.NewCurrentAddressService(currentAddressRepo)
	currentAddressCtrl := _currentAddressController.NewCurrentAddressController(currentAddressService)

	routesInit := _routes.ControllerList{
		JwtMiddleware:            configJWT.Init(),
		AdminController:          *adminCtrl,
		VaccineController:        *vaccineCtrl,
		CurrentAddressController: *currentAddressCtrl,
	}
	e := echo.New()
	routesInit.RoutesRegister(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}
