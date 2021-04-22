package repository

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/interface/database"
)

// NotificationRepositoryInterface ...
type NotificationRepositoryInterface interface {
	Where(id string) (account model.Notifications, err error)
}

// NotificationRepository ...
type NotificationRepository struct {
	database.SQLHandler
}

// Where ...
func (repo *NotificationRepository) Where(id string) (app model.Notifications, err error) {
	repo.SQLHandler.Where(&app, "id = ?", id)
	return
}
