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

	return controllers.NewSuccesResponse(c, response.FromDomain(data))
}
