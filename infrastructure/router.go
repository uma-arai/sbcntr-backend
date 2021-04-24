package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	handlers "github.com/uma-arai/sbcntr-backend/handler"
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
	// TODO アイテムを追加
	e.GET("/v1/Items", AppHandler.GetItems())
	//e.POST("/v1/Item", AppHandler.)
	e.POST("/v1/Item/favorite", AppHandler.GetItems())

	e.GET("/v1/Notifications", NotificationHandler.GetNotifications())
	e.GET("/v1/Notification/Count", NotificationHandler.GetUnreadNotificationCount())
	e.POST("/v1/Notifications/Read", NotificationHandler.PostNotificationsRead())
	return e
}
