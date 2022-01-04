package currentAddress

import (
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/currentAddress/request"
	"ca-reservaksin/controllers/currentAddress/response"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type CurrentAddressController struct {
	CurrentAddressService currentAddress.Service
}

func NewCurrentAddressController(service currentAddress.Service) *CurrentAddressController {
	return &CurrentAddressController{
		CurrentAddressService: service,
	}
}

func (ctrl *CurrentAddressController) Create(c echo.Context) error {
	req := request.CurrentAddress{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.CurrentAddressService.Create(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *CurrentAddressController) GetByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errors.New("id: id is empty"))
	}

	data, err := ctrl.CurrentAddressService.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *CurrentAddressController) Update(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errors.New("id: id is empty"))
	}

	req := request.CurrentAddress{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.CurrentAddressService.Update(id, req.ToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}

		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *CurrentAddressController) Delete(c echo.Context) error {
	id := c.Param("id")

	res, err := ctrl.CurrentAddressService.Delete(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, res)
}
