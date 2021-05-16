package handlers

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/uma-arai/sbcntr-backend/domain/repository"
	"github.com/uma-arai/sbcntr-backend/interface/database"
	"github.com/uma-arai/sbcntr-backend/usecase"
	"github.com/uma-arai/sbcntr-backend/utils"
)

// NotificationHandler ...
type NotificationHandler struct {
	Interactor usecase.NotificationInteractor
}

// NewNotificationHandler ...
func NewNotificationHandler(sqlHandler database.SQLHandler) *NotificationHandler {
	return &NotificationHandler{
		Interactor: usecase.NotificationInteractor{
			NotificationRepository: &repository.NotificationRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetNotifications ...
func (handler *NotificationHandler) GetNotifications() echo.HandlerFunc {
	return func(c echo.Context) (err error) {

		id := c.QueryParam("id")
		//data := &model.Notification{
		//	ID:          "1",
		//	Title:       "notif",
		//	Description: "notification detail",
		//	Category:    "info",
		//	Read:        false,
		//	CreatedAt:   "2021/03/17 10:39:06",
		//	UpdatedAt:   "2021/03/17 10:39:06",
		//}
		//body := &model.Notifications{Data: []model.Notification{
		//	*data,
		//}}
		//return c.JSON(http.StatusOK, body)
		//

		resJSON, err := handler.Interactor.GetNotifications(id)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// GetUnreadNotificationCount ...
func (handler *NotificationHandler) GetUnreadNotificationCount() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		//body := &model.NotificationCount{
		//	Data: 5,
		//}
		//return c.JSON(http.StatusOK, body)
		resJSON, err := handler.Interactor.GetUnreadNotificationCount()
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// PostNotificationsRead ...
func (handler *NotificationHandler) PostNotificationsRead() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		//id := c.FormValue("id")
		//if id == "" {
		//	return c.JSON(400, "Invalid parameter id.")
		//}

		resJSON, err := handler.Interactor.MarkNotificationsRead()
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}
