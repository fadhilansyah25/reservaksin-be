package routes

import (
	"ca-reservaksin/controllers/admin"
	"ca-reservaksin/controllers/booking"
	"ca-reservaksin/controllers/citizen"
	"ca-reservaksin/controllers/currentAddress"
	"ca-reservaksin/controllers/healthFacilities"
	"ca-reservaksin/controllers/session"
	"ca-reservaksin/controllers/vaccine"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtMiddleware              middleware.JWTConfig
	AdminController            admin.AdminController
	VaccineController          vaccine.VaccineController
	CurrentAddressController   currentAddress.CurrentAddressController
	HealthFacilitiesController healthFacilities.HealthFacilitiesController
	SessionController          session.Sessioncontroller
	CitizenController          citizen.CitizenController
	BookingController          booking.BookingController
}

func (cl *ControllerList) RoutesRegister(e *echo.Echo) {
	admin := e.Group("admin")
	admin.POST("/register", cl.AdminController.Register)
	admin.POST("/login", cl.AdminController.Login)
	admin.GET("/:id", cl.AdminController.GetByID)

	vaccine := e.Group("vaccine")
	vaccine.POST("", cl.VaccineController.Create)
	vaccine.GET("", cl.VaccineController.FetchAll)
	vaccine.GET("/:id", cl.VaccineController.GetByID)
	vaccine.PUT("/:id", cl.VaccineController.Update)
	vaccine.DELETE("/:id", cl.VaccineController.Delete)
	vaccine.GET("/admin/:id", cl.VaccineController.GetByAdminID)

	address := e.Group("address")
	address.POST("", cl.CurrentAddressController.Create)
	address.GET("/:id", cl.CurrentAddressController.GetByID)
	address.PUT("/:id", cl.CurrentAddressController.Update)
	address.DELETE("/:id", cl.CurrentAddressController.Delete)

	healthFacilities := e.Group("health-facilities")
	healthFacilities.POST("", cl.HealthFacilitiesController.Create)
	healthFacilities.GET("/:id", cl.HealthFacilitiesController.GetByID)
	healthFacilities.PUT("/:id", cl.HealthFacilitiesController.Update)
	healthFacilities.DELETE("/:id", cl.HealthFacilitiesController.Delete)
	healthFacilities.GET("/admin/:id", cl.HealthFacilitiesController.GetByIdAdmin)
	healthFacilities.GET("/maps-place", cl.HealthFacilitiesController.FetchAllForMapsResponse)

	session := e.Group("session")
	session.POST("", cl.SessionController.Create)
	session.GET("/:id", cl.SessionController.GetByID)
	session.GET("", cl.SessionController.FetchAll)
	session.GET("/nearest-facilities", cl.SessionController.NearFacilities)
	session.PUT("/:id", cl.SessionController.Update)
	session.DELETE("/:id", cl.SessionController.Delete)
	session.GET("/current/admin/:id", cl.SessionController.FetchSessionCurrent)
	session.GET("/history/admin/:id", cl.SessionController.FetchSessionHistory)
	session.GET("/upcoming/admin/:id", cl.SessionController.FetchSessionUpcoming)
	session.GET("/admin/:id", cl.SessionController.FetchSessionByAdminId)

	citizen := e.Group("citizen")
	citizen.POST("/register", cl.CitizenController.Register)
	citizen.POST("/login", cl.CitizenController.Login)
	citizen.GET("/:id", cl.CitizenController.GetCitizenByID)
	citizen.GET("/admin/:id", cl.CitizenController.FetchCitizenByAdminID)
	citizen.GET("/family", cl.CitizenController.FetchCitizenByNoKK)
	citizen.PUT("/:id", cl.CitizenController.Update)

	booking := e.Group("booking")
	booking.POST("", cl.BookingController.BookingSession)
	booking.GET("/citizen/:id", cl.BookingController.GetByCitizenID)
	booking.GET("/session/:id", cl.BookingController.GetBySessionID)
	booking.GET("/nokk/:id", cl.BookingController.GetByNoKK)
	booking.PATCH("/status/:id", cl.BookingController.UpdateBookingStatusByID)
}
