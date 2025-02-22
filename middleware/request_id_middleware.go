package middleware

import (
	"time"

	"github.com/google/uuid"
	echo "github.com/labstack/echo/v4"
)

func RequestIDMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Set(KeyRequestID, uuid.NewString())
			c.Set(KeyElapsedTime, time.Now())
			return next(c)
		}
	}
}
