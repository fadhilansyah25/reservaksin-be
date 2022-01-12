package citizen

import (
	"ca-reservaksin/businesses/citizen"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/citizen/request"
	"ca-reservaksin/controllers/citizen/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type CitizenController struct {
	citizenService citizen.Service
}

func NewCitizenController(uc citizen.Service) *CitizenController {
	return &CitizenController{
		citizenService: uc,
	}
}

func (ctrl *CitizenController) Register(c echo.Context) error {
	req := request.Citizen{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	data, err := ctrl.citizenService.Register(req.ToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "duplicate data") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *CitizenController) LoginByEmail(c echo.Context) error {
	req := request.CitizenLoginEmail{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.citizenService.LoginByEmail(req.Email, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect (email) or (password)") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *CitizenController) LoginByNIK(c echo.Context) error {
	req := request.CitizenLoginNIK{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.citizenService.LoginByNIK(req.Nik, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect (nik) or (password)") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *CitizenController) Update(c echo.Context) error {
	id := c.Param("id")
	req := request.Citizen{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.citizenService.Update(id, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))

}

func (ctrl *CitizenController) Delete(c echo.Context) error {
	id := c.Param("id")

	res, err := ctrl.citizenService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, res)
}
