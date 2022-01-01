package routes

import (
	"ca-reservaksin/controllers/admin"
	"ca-reservaksin/controllers/vaccine"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtMiddleware     middleware.JWTConfig
	AdminController   admin.AdminController
	VaccineController vaccine.VaccineController
}

func (cl *ControllerList) RoutesRegister(e *echo.Echo) {
	admin := e.Group("admin")
	admin.POST("/register", cl.AdminController.Register)
	admin.GET("/login", cl.AdminController.Login)
	admin.GET("/:id", cl.AdminController.GetByID)

	vaccine := e.Group("vaccine")
	vaccine.POST("", cl.VaccineController.Create)
	vaccine.GET("", cl.VaccineController.FetchAll)
	vaccine.GET("/:id", cl.VaccineController.GetByID)
	vaccine.PUT("/:id", cl.VaccineController.Update)
	vaccine.DELETE("/:id", cl.VaccineController.Delete)
}
