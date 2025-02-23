package check_handler

import (
	"fmt"

	"github.com/labstack/echo/v4"
	svc "github.com/rbennum/service-account/services/check"
	"github.com/rs/zerolog"
)

type CheckkHandler struct {
	svc    svc.CheckService
	logger zerolog.Logger
}

func New(svc svc.CheckService, logger zerolog.Logger) CheckkHandler {
	return CheckkHandler{
		svc:    svc,
		logger: logger,
	}
}

func (t *CheckkHandler) CheckBalance(c echo.Context) error {
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
