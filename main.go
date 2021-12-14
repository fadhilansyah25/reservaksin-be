package main

import (
	"Reservaksin-BE/config/database"
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	database.Connection()
	e.GET("/", HelloController)

	e.Logger.Fatal(e.Start(":8000"))
}

func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World Jakarta")
}
