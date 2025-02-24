package tarik_handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
	tm "github.com/rbennum/service-account/models/accounts"
	as "github.com/rbennum/service-account/services/accounts"
	"github.com/rs/zerolog"
)

type TarikHandler struct {
	svc    as.AccountsService
	logger zerolog.Logger
}

func New(svc as.AccountsService, logger zerolog.Logger) TarikHandler {
	return TarikHandler{
		svc:    svc,
		logger: logger,
	}
}

func (t *TarikHandler) WithdrawBalance(c echo.Context) error {
	var input tm.RequestBody
	if err := c.Bind(&input); err != nil {
		return echo.NewHTTPError(
			http.StatusBadRequest,
			tm.ResponseBody{
				StatusCode:   http.StatusBadRequest,
				ErrorMessage: err.Error(),
			},
		)
	}

	res, err := t.svc.WithdrawBalance(c.Request().Context(), input.Transferred, input.AccountNumber)
	if err != nil {
		return echo.NewHTTPError(res.StatusCode, res)
	}
	return c.JSON(res.StatusCode, res)
}
