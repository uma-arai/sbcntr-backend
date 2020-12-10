package handlers

import (
	"github.com/labstack/echo"
)

// HelloWorldHandler ...
type HelloWorldHandler struct {
}

// NewHelloWorldHandler ...
func NewHelloWorldHandler() *HelloWorldHandler {
	return &HelloWorldHandler{}
}

// SayHelloWorld ...
func (handler *HelloWorldHandler) SayHelloWorld() echo.HandlerFunc {

	return func(c echo.Context) error {
		return c.JSON(200, "Hello world!")
	}
}
