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

func (ctrl *CitizenController) Login(c echo.Context) error {
	req := request.CitizenLogin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	dataCitizen, token, err := ctrl.citizenService.Login(req.EmailOrNIK, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect (Email) or (NIK)") {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, err)
		}
		if strings.Contains(err.Error(), "incorrect (Password)") {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		DataCitizen response.CitizenResponse
		Token       string `json:"token"`
	}{DataCitizen: *response.FromDomain(dataCitizen), Token: token}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *CitizenController) Update(c echo.Context) error {
	id := c.Param("id")
	req := request.CitizenEdit{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.citizenService.Update(id, req.ToDomainCitizenEdit())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
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

func (ctrl *CitizenController) GetCitizenByID(c echo.Context) error {
	id := c.Param("id")

	res, err := ctrl.citizenService.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(res))
}

func (ctrl *CitizenController) FetchCitizenByAdminID(c echo.Context) error {
	adminID := c.Param("id")
	res, err := ctrl.citizenService.GetByAdminID(adminID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(res) == 0 {
		return controllers.NewEmptyDataResponse(c, response.FromDomainOfArray(res))
	}

	return controllers.NewSuccesResponse(c, response.FromDomainOfArray(res))
}

func (ctrl *CitizenController) FetchCitizenByNoKK(c echo.Context) error {
	noKK := c.QueryParam("nokk")
	res, err := ctrl.citizenService.GetByNoKK(noKK)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(res) == 0 {
		return controllers.NewEmptyDataResponse(c, response.FromDomainOfArray(res))
	}

	return controllers.NewSuccesResponse(c, response.FromDomainOfArray(res))
}
