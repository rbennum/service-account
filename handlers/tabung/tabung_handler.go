package tabung_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	mdl "github.com/rbennum/service-account/models/tabung"
	svc "github.com/rbennum/service-account/services/tabung"
	"github.com/rs/zerolog"
)

type TabungHandler struct {
	svc    svc.TabungService
	logger zerolog.Logger
}

func New(svc svc.TabungService, logger zerolog.Logger) TabungHandler {
	return TabungHandler{
		svc:    svc,
		logger: logger,
	}
}

func (t *TabungHandler) DepositBalance(c echo.Context) error {
	var input mdl.RequestBody
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			mdl.ResponseBody{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: err.Error(),
			},
		)
	}

	res, err := t.svc.DepositBalance(c.Request().Context(), input.Transferred, input.AccountNumber)
	if err != nil {
		return echo.NewHTTPError(res.StatusCode, res)
	}
	return c.JSON(res.StatusCode, res)
}
