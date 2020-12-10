package handlers

import (
	"github.com/labstack/echo"
)

// HealthCheckHandler ...
type HealthCheckHandler struct {
}

// NewHealthCheckHandler ...
func NewHealthCheckHandler() *HealthCheckHandler {
	return &HealthCheckHandler{}
}

// HealthCheck ...
func (handler *HealthCheckHandler) HealthCheck() echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(200, nil)
	}
}
