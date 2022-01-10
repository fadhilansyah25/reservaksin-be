package admin

import (
	"ca-reservaksin/businesses/admin"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/admin/request"
	"ca-reservaksin/controllers/admin/response"
	"errors"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type AdminController struct {
	AdminService admin.Service
}

func NewAdminController(service admin.Service) *AdminController {
	return &AdminController{
		AdminService: service,
	}
}

func (ctrl *AdminController) Register(c echo.Context) error {
	req := request.Admin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.AdminService.Register(req.ToDomain())
	if err != nil {
		if strings.Contains(err.Error(), "duplicate data") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}

func (ctrl *AdminController) Login(c echo.Context) error {
	req := request.AdminLogin{}
	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.AdminService.Login(req.Username, req.Password)
	if err != nil {
		if strings.Contains(err.Error(), "incorrect (Username) or (Password)") {
			return controllers.NewErrorResponse(c, http.StatusUnauthorized, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controllers.NewSuccesResponse(c, res)
}

func (ctrl *AdminController) GetByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, errors.New("id: id is empty"))
	}

	admin, err := ctrl.AdminService.GetByID(id)
	if err != nil {
		if !strings.Contains(err.Error(), "not found") {
			return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
		}
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccesResponse(c, response.FromDomain(admin))
}
