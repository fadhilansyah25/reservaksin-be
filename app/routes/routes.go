package routes

import (
	"ca-reservaksin/controllers/admin"
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

	session := e.Group("session")
	session.POST("", cl.SessionController.Create)
	session.GET("/:id", cl.SessionController.GetByID)
	session.GET("", cl.SessionController.FetchAll)
	session.GET("/nearfacilities", cl.SessionController.NearFacilities)
	session.PUT("/:id", cl.SessionController.Update)
	session.DELETE("/:id", cl.SessionController.Delete)
	session.GET("/current", cl.SessionController.FetchSessionCurrent)
	session.GET("/history", cl.SessionController.FetchSessionHistory)
	session.GET("/upcoming", cl.SessionController.FetchSessionUpcoming)
}
