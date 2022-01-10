package session

import (
	"ca-reservaksin/businesses/session"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/session/request"
	"ca-reservaksin/controllers/session/response"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type Sessioncontroller struct {
	SessionService session.Service
}

func NewSessioncontroller(service session.Service) *Sessioncontroller {
	return &Sessioncontroller{
		SessionService: service,
	}
}

func (ctrl *Sessioncontroller) Create(c echo.Context) error {
	req := request.Session{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.SessionService.Create(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *Sessioncontroller) GetByID(c echo.Context) error {
	id := c.Param("id")

	data, err := ctrl.SessionService.GetByID(id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *Sessioncontroller) FetchAll(c echo.Context) error {
	data, err := ctrl.SessionService.FetchAll()
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainArray(data))
}

func (ctrl *Sessioncontroller) NearFacilities(c echo.Context) error {
	lat, _ := strconv.ParseFloat(c.QueryParam("lat"), 64)
	lng, _ := strconv.ParseFloat(c.QueryParam("lng"), 64)

	res, err := ctrl.SessionService.GetByLatLong(lat, lng)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomainArrayResult(res))
}

func (ctrl *Sessioncontroller) Update(c echo.Context) error {
	id := c.Param("id")
	req := request.Session{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.SessionService.Update(id, req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *Sessioncontroller) Delete(c echo.Context) error {
	id := c.Param("id")
	res, err := ctrl.SessionService.Delete(id)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *Sessioncontroller) FetchSessionHistory(c echo.Context) error {
	adminID := c.Param("id")
	res, err := ctrl.SessionService.FetchByHistory(adminID, "history")
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(res) == 0 {
		return controllers.NewEmptyDataResponse(c, res)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainArraySimpleRes(res))
}

func (ctrl *Sessioncontroller) FetchSessionCurrent(c echo.Context) error {
	adminID := c.Param("id")
	res, err := ctrl.SessionService.FetchByHistory(adminID, "current")
	if err != nil {
		return controllers.NewEmptyDataResponse(c, res)
	}

	if len(res) == 0 {
		return controllers.NewEmptyDataResponse(c, response.FromDomainArraySimpleRes(res))
	}

	return controllers.NewSuccesResponse(c, response.FromDomainArraySimpleRes(res))
}

func (ctrl *Sessioncontroller) FetchSessionUpcoming(c echo.Context) error {
	adminID := c.Param("id")
	res, err := ctrl.SessionService.FetchByHistory(adminID, "upcoming")
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(res) == 0 {
		return controllers.NewEmptyDataResponse(c, response.FromDomainArraySimpleRes(res))
	}

	return controllers.NewSuccesResponse(c, response.FromDomainArraySimpleRes(res))
}

func (ctrl *Sessioncontroller) FetchSessionByAdminId(c echo.Context) error {
	adminID := c.Param("id")
	res, err := ctrl.SessionService.FetchAllByAdminID(adminID)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	if len(res) == 0 {
		return controllers.NewEmptyDataResponse(c, response.FromDomainArraySimpleRes(res))
	}

	return controllers.NewSuccesResponse(c, response.FromDomainArraySimpleRes(res))
}
