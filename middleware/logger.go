package middleware

import (
	"time"

	"github.com/labstack/echo/v4"
	logger "github.com/rbennum/service-account/utils/log"
)

func ResponseLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			requestId := c.Get(KeyRequestID).(string)
			startTime := c.Get(KeyElapsedTime).(time.Time)
			elapsed := time.Since(startTime)
			logger := logger.Logger.Info().
				Str("request_id", requestId).
				Str("method", c.Request().Method).
				Str("path", c.Request().URL.Path).
				Int("status", c.Response().Status).
				Dur("elapsed", elapsed)

			err := next(c)
			if err != nil {
				logger.Err(err).Msg("request completed with error")
			} else {
				logger.Msg("request processed")
			}
			return err
		}
	}
}
