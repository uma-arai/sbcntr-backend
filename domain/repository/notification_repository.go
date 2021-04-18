package repository

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/interface/database"
)

// NotificationRepositoryInterface ...
type NotificationRepositoryInterface interface {
	Where(id string) (account model.Notification, err error)
}

// AppRepository ...
type NotificationRepository struct {
	database.SQLHandler
}

// Where ...
func (repo *NotificationRepository) Where(id string) (app model.Notification, err error) {
	repo.SQLHandler.Where(&app, "id = ?", id)
	return
}
