package check_handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	as "github.com/rbennum/service-account/services/accounts"
	"github.com/rs/zerolog"
)

type CheckHandler struct {
	svc    as.AccountsService
	logger zerolog.Logger
}

func New(svc as.AccountsService, logger zerolog.Logger) CheckHandler {
	return CheckHandler{
		svc:    svc,
		logger: logger,
	}
}

func (t *CheckHandler) CheckBalance(c echo.Context) error {
	accountNum := c.Param("no_rekening")
	fmt.Println(accountNum)
	res, err := t.svc.CheckBalance(
		c.Request().Context(),
		accountNum,
	)
	if err != nil {
		return echo.NewHTTPError(res.StatusCode, res)
	}
	return c.JSON(res.StatusCode, res)
}
