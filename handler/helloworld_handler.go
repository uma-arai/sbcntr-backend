package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/uma-arai/sbcntr-backend/domain/model"
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
	body := &model.Hello{
		Data: "Hello world",
	}
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, body)
	}
}
