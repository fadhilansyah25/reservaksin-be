package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rs/cors"
	"github.com/spf13/viper"
	"gorm.io/gorm"

	_middlewares "ca-reservaksin/app/middlewares"
	_routes "ca-reservaksin/app/routes"
	_driverFactory "ca-reservaksin/drivers"
	_dbDriver "ca-reservaksin/drivers/mysql"

	_adminService "ca-reservaksin/businesses/admin"
	_bookingService "ca-reservaksin/businesses/booking"
	_citizenService "ca-reservaksin/businesses/citizen"
	_currentAddressService "ca-reservaksin/businesses/currentAddress"
	_healthFacilitiesService "ca-reservaksin/businesses/healthFacilities"
	_sessionService "ca-reservaksin/businesses/session"
	_vaccineService "ca-reservaksin/businesses/vaccine"
	_adminController "ca-reservaksin/controllers/admin"
	_bookingController "ca-reservaksin/controllers/booking"
	_citizenController "ca-reservaksin/controllers/citizen"
	_currentAddressController "ca-reservaksin/controllers/currentAddress"
	_healthFacilitiesController "ca-reservaksin/controllers/healthFacilities"
	_sessionController "ca-reservaksin/controllers/session"
	_vaccineController "ca-reservaksin/controllers/vaccine"
	_AdminRepo "ca-reservaksin/drivers/database/admin"
	_bookingRepo "ca-reservaksin/drivers/database/booking"
	_citizenRepo "ca-reservaksin/drivers/database/citizen"
	_currentAddressRepo "ca-reservaksin/drivers/database/currentAddress"
	_healthFacilitiesRepo "ca-reservaksin/drivers/database/healthFacilities"
	_sessionRepo "ca-reservaksin/drivers/database/session"
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
		&_healthFacilitiesRepo.HealthFacilities{},
		&_sessionRepo.Session{},
		&_citizenRepo.Citizen{},
		&_bookingRepo.Booking{},
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

	citizenRepo := _driverFactory.NewCitizenRepository(db)
	citizenService := _citizenService.NewCitizenService(citizenRepo, currentAddressRepo, &configJWT)
	citizenCtrl := _citizenController.NewCitizenController(citizenService)

	bookingRepo := _driverFactory.NewBookingRepository(db)
	bookingService := _bookingService.NewBookingSessionService(bookingRepo)
	bookingCtrl := _bookingController.NewBookingController(bookingService)

	routesInit := _routes.ControllerList{
		JwtMiddleware:              configJWT.Init(),
		AdminController:            *adminCtrl,
		VaccineController:          *vaccineCtrl,
		CurrentAddressController:   *currentAddressCtrl,
		HealthFacilitiesController: *healthFacilitiesCtrl,
		SessionController:          *sessionCtrl,
		CitizenController:          *citizenCtrl,
		BookingController:          *bookingCtrl,
	}
	e := echo.New()

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodPatch},
		AllowedHeaders: []string{"*"},
		Debug:          false,
	})
	e.Use(echo.WrapMiddleware(corsMiddleware.Handler))

	routesInit.RoutesRegister(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!\n")
	})

	log.Fatal(e.Start(viper.GetString("server.address")))
}
