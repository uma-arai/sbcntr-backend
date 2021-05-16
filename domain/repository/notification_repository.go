package repository

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/interface/database"
)

// NotificationRepositoryInterface ...
type NotificationRepositoryInterface interface {
	Where(id string) (account model.Notifications, err error)
	FindAll() (notifications model.Notifications, err error)
	Count(query interface{}, args ...interface{}) (data model.NotificationCount, err error)
	Update(value map[string]interface{}, query interface{}, args ...interface{}) (notification model.Notification, err error)
}

// NotificationRepository ....
type NotificationRepository struct {
	database.SQLHandler
}

// Where ...
func (repo *NotificationRepository) Where(id string) (notifications model.Notifications, err error) {
	repo.SQLHandler.Where(&notifications.Data, "id = ?", id)
	return
}

// FindAll ...
func (repo *NotificationRepository) FindAll() (notifications model.Notifications, err error) {
	repo.SQLHandler.Scan(&notifications.Data, "id desc")
	return
}

// Count ...
func (repo *NotificationRepository) Count(query interface{}, args ...interface{}) (data model.NotificationCount, err error) {
	var count int
	repo.SQLHandler.Count(&count, &model.Notification{}, query, args...)

	return model.NotificationCount{Data: count}, nil
}

// Update ...
func (repo *NotificationRepository) Update(value map[string]interface{}, query interface{}, args ...interface{}) (notification model.Notification, err error) {
	// NOTE: When update with struct, GORM will only update non-zero fields,
	// you might want to use map to update attributes or use Select to specify fields to update
	repo.SQLHandler.Update(&notification, value, query, args...)

	return
}
