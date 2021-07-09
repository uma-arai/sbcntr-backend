package infrastructure

import (
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	handlers "github.com/uma-arai/sbcntr-backend/handler"
)

// Router ...
func Router() *echo.Echo {
	e := echo.New()

	// Middleware
	logger := middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"id":"${id}","time":"${time_rfc3339}","remote_ip":"${remote_ip}",` +
			`"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",` +
			`"status":${status},"error":"${error}"}` + "\n",
		Output: os.Stdout,
	})
	e.Use(logger)
	e.Use(middleware.Recover())
	e.Logger.SetLevel(log.INFO)
	e.HideBanner = true
	e.HidePort = false

	AppHandler := handlers.NewAppHandler(NewSQLHandler())
	healthCheckHandler := handlers.NewHealthCheckHandler()
	helloWorldHandler := handlers.NewHelloWorldHandler()
	NotificationHandler := handlers.NewNotificationHandler(NewSQLHandler())

	e.GET("/", healthCheckHandler.HealthCheck())
	e.GET("/healthcheck", healthCheckHandler.HealthCheck())
	e.GET("/v1/helloworld", helloWorldHandler.SayHelloWorld())

	e.GET("/v1/Items", AppHandler.GetItems())
	e.POST("/v1/Item", AppHandler.CreateItem())

	// TODO: RESTを意識した形式とする
	e.POST("/v1/Item/favorite", AppHandler.UpdateFavoriteAttr())

	e.GET("/v1/Notifications", NotificationHandler.GetNotifications())
	e.GET("/v1/Notifications/Count", NotificationHandler.GetUnreadNotificationCount())
	e.POST("/v1/Notifications/Read", NotificationHandler.PostNotificationsRead())

	return e
}

//func logFormat() string {
//	var format string
//
//	format += "{"
//	format += `"time":"${time_rfc3339_nano}",`
//	format += `"id":"${id}",`
//	format += `"uri":"${uri}",`
//	format += `"remote_ip":"${remote_ip}",`
//	format += `"host":"${host}",`
//	format += `"method":"${method}",`
//	format += `"user_agent":"${user_agent}",`
//	format += `"status":"${status}",`
//	format += `"error":"${error}",`
//	format += `"latency":"${latency}",`
//	format += `"log_type":"AccessLog"`
//	format += "}\n"
//
//	return format
//}
