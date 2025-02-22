package daftar_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	dm "github.com/rbennum/service-account/models/daftar"
	svc "github.com/rbennum/service-account/services/daftar"
	"github.com/rs/zerolog"
)

type DaftarHandler struct {
	service svc.DaftarService
	logger  zerolog.Logger
}

func New(service svc.DaftarService, logger zerolog.Logger) DaftarHandler {
	return DaftarHandler{
		service: service,
		logger:  logger,
	}
}

func (d *DaftarHandler) PostDaftar(c echo.Context) error {
	var input dm.RequestBody
	if err := c.Bind(&input); err != nil {
		return c.JSON(
			http.StatusBadRequest,
			dm.ResponseBody{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: err.Error(),
			},
		)
	}
	res, err := d.service.CreateCustomer(
		c.Request().Context(),
		input,
	)
	if err != nil {
		return echo.NewHTTPError(res.StatusCode, res)
	}

	return c.JSON(res.StatusCode, res)
}
