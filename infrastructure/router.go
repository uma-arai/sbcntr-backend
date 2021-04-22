package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/uma-arai/sbcntr-backend/handler"
)

// Router ...
func Router() *echo.Echo {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	AppHandler := handlers.NewAppHandler(NewSQLHandler())
	healthCheckHandler := handlers.NewHealthCheckHandler()
	helloWorldHandler := handlers.NewHelloWorldHandler()
	NotificationHandler := handlers.NewNotificationHandler(NewSQLHandler())

	e.GET("/", healthCheckHandler.HealthCheck())
	e.GET("/healthcheck", healthCheckHandler.HealthCheck())
	e.GET("/v1/helloworld", helloWorldHandler.SayHelloWorld())
	e.GET("/v1/app", AppHandler.GetAppInfo())

	e.GET("/v1/Notifications", NotificationHandler.GetNotifications())
	e.GET("/v1/Notification/Count", NotificationHandler.GetNotificationCount())
	return e
}
