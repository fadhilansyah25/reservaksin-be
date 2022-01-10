package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_middlewares "ca-reservaksin/app/middlewares"
	_routes "ca-reservaksin/app/routes"
	_driverFactory "ca-reservaksin/drivers"
	_dbDriver "ca-reservaksin/drivers/mysql"

	_adminService "ca-reservaksin/businesses/admin"
	_currentAddressService "ca-reservaksin/businesses/currentAddress"
	_healthFacilitiesService "ca-reservaksin/businesses/healthFacilities"
	_sessionService "ca-reservaksin/businesses/session"
	_vaccineService "ca-reservaksin/businesses/vaccine"
	_adminController "ca-reservaksin/controllers/admin"
	_currentAddressController "ca-reservaksin/controllers/currentAddress"
	_healthFacilitiesController "ca-reservaksin/controllers/healthFacilities"
	_sessionController "ca-reservaksin/controllers/session"
	_vaccineController "ca-reservaksin/controllers/vaccine"
	_AdminRepo "ca-reservaksin/drivers/database/admin"
	_currentAddressRepo "ca-reservaksin/drivers/database/currentAddress"
	_healthFacilitiesRepo "ca-reservaksin/drivers/database/healthFacilities"
	_session "ca-reservaksin/drivers/database/session"
	_VaccineRepo "ca-reservaksin/drivers/database/vaccine"
)

func init() {
	viper.SetConfigFile(`config.json`)
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
		&_healthFacilitiesRepo.HealthFacilities{},
		&_session.Session{},
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

	healthFacilitiesRepo := _driverFactory.NewHealthFacilitiesRepository(db)
	healthFacilitiesService := _healthFacilitiesService.NewHealthFacilitiesService(healthFacilitiesRepo, currentAddressRepo)
	healthFacilitiesCtrl := _healthFacilitiesController.NewHealthFacilitiesController(healthFacilitiesService)

	sessionRepo := _driverFactory.NewSessionRepository(db)
	sessionService := _sessionService.NewSessionService(sessionRepo, currentAddressRepo)
	sessionCtrl := _sessionController.NewSessioncontroller(sessionService)

	routesInit := _routes.ControllerList{
		JwtMiddleware:              configJWT.Init(),
		AdminController:            *adminCtrl,
		VaccineController:          *vaccineCtrl,
		CurrentAddressController:   *currentAddressCtrl,
		HealthFacilitiesController: *healthFacilitiesCtrl,
		SessionController:          *sessionCtrl,
	}
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
		AllowedHeaders: []string{"*"},
		Debug:          true,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	routesInit.RoutesRegister(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	port := os.Getenv("PORT")
	log.Fatal(e.Start(":" + port))
}
