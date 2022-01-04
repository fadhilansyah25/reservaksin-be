package healthFacilities

import (
	"ca-reservaksin/businesses/currentAddress"
	"ca-reservaksin/businesses/healthFacilities"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/healthFacilities/request"
	"ca-reservaksin/controllers/healthFacilities/response"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type HealthFacilitiesController struct {
	FacilitiesService healthFacilities.Service
	AddressService    currentAddress.Service
}

func NewHealthFacilitiesController(service healthFacilities.Service, addressService currentAddress.Service) *HealthFacilitiesController {
	return &HealthFacilitiesController{
		FacilitiesService: service,
		AddressService:    addressService,
	}
}

func (ctrl *HealthFacilitiesController) Create(c echo.Context) error {
	req := request.HealthFacilities{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.FacilitiesService.Create(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	address, err := ctrl.AddressService.GetByID(data.CurrentAddressID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(&data, &address))
}

func (ctrl *HealthFacilitiesController) GetByID(c echo.Context) error {
	id := c.Param("id")

	data, err := ctrl.FacilitiesService.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate data") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	address, err := ctrl.AddressService.GetByID(data.CurrentAddressID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(&data, &address))
}

func (ctrl *HealthFacilitiesController) Update(c echo.Context) error {
	id := c.Param("id")
	req := request.HealthFacilities{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	facilitiesDomain, addressDomain := req.ToDomain()
	data, address, err := ctrl.FacilitiesService.Update(id, facilitiesDomain, addressDomain)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(&data, &address))

}
