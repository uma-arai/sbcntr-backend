package repository

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/interface/database"
)

// NotificationRepositoryInterface ...
type NotificationRepositoryInterface interface {
	Where(id string) (account model.Notifications, err error)
	FindAll() (notifications model.Notifications, err error)
	Count() (data model.NotificationCount, err error)
}

// NotificationRepository ....
type NotificationRepository struct {
	database.SQLHandler
}

// Where ...
func (repo *NotificationRepository) Where(id string) (app model.Notifications, err error) {
	repo.SQLHandler.Where(&app, "id = ?", id)
	return
}

// FindAll ...
func (repo *NotificationRepository) FindAll() (notifications model.Notifications, err error) {
	repo.SQLHandler.Scan(&notifications.Data, "id desc")
	return
}

// Count ...
func (repo *NotificationRepository) Count() (data model.NotificationCount, err error) {
	var count int
	repo.SQLHandler.Count(&count, &model.Notification{})

	return model.NotificationCount{Data: count}, nil
}
