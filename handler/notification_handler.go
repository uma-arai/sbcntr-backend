package handlers

import (
	"github.com/labstack/echo"
	"github.com/uma-arai/sbcntr-backend/domain/repository"
	"github.com/uma-arai/sbcntr-backend/interface/database"
	"github.com/uma-arai/sbcntr-backend/usecase"
	"github.com/uma-arai/sbcntr-backend/utils"
)

// AppHandler ...
type NotificationHandler struct {
	Interactor usecase.NotificationInteractor
}

// NewNotifiationHandler ...
func NewNotificationHandler(sqlHandler database.SQLHandler) *NotificationHandler {
	return &NotificationHandler{
		Interactor: usecase.NotificationInteractor{
			NotificationRepository: &repository.NotificationRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetAppInfo ...
func (handler *NotificationHandler) GetNotifications() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		id := c.QueryParam("id")
		if id == "" {
			return c.JSON(400, "Invalid parameter id.")
		}

		//err = utils.HeaderCheck(c, echo.HeaderContentType, echo.MIMEApplicationJSON)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		//err = utils.ClientIDCheck(c)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		resJSON, err := handler.Interactor.GetNotifications(id)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(200, resJSON)
	}
}
