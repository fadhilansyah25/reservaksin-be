package booking

import (
	"ca-reservaksin/businesses/booking"
	"ca-reservaksin/controllers"
	"ca-reservaksin/controllers/booking/request"
	"ca-reservaksin/controllers/booking/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BookingController struct {
	bookingService booking.Service
}

func NewBookingController(service booking.Service) *BookingController {
	return &BookingController{
		bookingService: service,
	}
}

func (ctrl *BookingController) BookingSession(c echo.Context) error {
	req := request.Booking{}

	if err := c.Bind(&req); err != nil {
		return controllers.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.bookingService.BookingSession(req.ToDomain())
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainBookingCitizen(data))
}

func (ctrl *BookingController) GetByCitizenID(c echo.Context) error {
	citizenId := c.Param("id")

	data, err := ctrl.bookingService.GetByCitizenID(citizenId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainOfArrayBookingCitizen(data))
}

func (ctrl *BookingController) GetBySessionID(c echo.Context) error {
	sessionId := c.Param("id")

	data, err := ctrl.bookingService.GetBySessionID(sessionId)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainOfArrayBookingSession(data))
}

func (ctrl *BookingController) GetByNoKK(c echo.Context) error {
	noKK := c.Param("id")

	data, err := ctrl.bookingService.GetByNoKK(noKK)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controllers.NewSuccesResponse(c, response.FromDomainOfArrayBookingCitizen(data))
}
