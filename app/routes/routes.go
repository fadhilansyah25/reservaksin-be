package routes

import (
	"ca-reservaksin/controllers/admin"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtMiddleware   middleware.JWTConfig
	AdminController admin.AdminController
}

func (cl *ControllerList) RoutesRegister(e *echo.Echo) {
	admin := e.Group("admin")
	admin.POST("/register", cl.AdminController.Register)
	admin.GET("/login", cl.AdminController.Login)
	admin.GET("/:id", cl.AdminController.GetByID)
}
