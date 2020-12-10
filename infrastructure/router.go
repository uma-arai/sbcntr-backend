package infrastructure

import (
	"github.com/iselegant/cnappdemo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// Router ...
func Router() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	basePath := "cnappdemo"

	AppHandler := handlers.NewAppHandler(NewSQLHandler())
	healthCheckHandler := handlers.NewHealthCheckHandler()
	helloWorldHandler := handlers.NewHelloWorldHandler()

	e.GET("/healthcheck", healthCheckHandler.HealthCheck())
	e.GET(basePath+"/v1/helloworld", helloWorldHandler.SayHelloWorld())
	e.GET(basePath+"/v1/app", AppHandler.GetAppInfo())
	return e
}
