package controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type BaseResponse struct {
	Meta struct {
		Status   int      `json:"status"`
		Message  string   `json:"message"`
		Messages []string `json:"messages,omitempty"`
	} `json:"meta"`
	Data interface{} `json:"data"`
}

func NewSuccesResponse(c echo.Context, data interface{}) error {
	res := BaseResponse{}
	res.Meta.Status = http.StatusOK
	res.Meta.Message = "succes"
	res.Data = data

	return c.JSON(http.StatusOK, res)
}

func NewErrorResponse(c echo.Context, status int, err error) error {
	res := BaseResponse{}
	res.Meta.Status = status
	res.Meta.Message = "Error"
	res.Meta.Messages = []string{err.Error()}

	return c.JSON(status, res)
}

func NewEmptyDataResponse(c echo.Context, data interface{}) error {
	res := BaseResponse{}
	res.Meta.Status = http.StatusOK
	res.Meta.Message = "Data is Empty"
	res.Data = data

	return c.JSON(http.StatusOK, res)
}
