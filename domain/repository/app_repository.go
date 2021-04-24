package repository

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/interface/database"
)

// AppRepositoryInterface ...
type AppRepositoryInterface interface {
	FindAll() (items model.Items, err error)
	Create(item model.Item) (out model.Item, err error)
	Update(value map[string]interface{}, query interface{}, args ...interface{}) (notification model.Notification, err error)
}

// AppRepository ...
type AppRepository struct {
	database.SQLHandler
}

// FindAll ...
func (repo *AppRepository) FindAll() (items model.Items, err error) {
	repo.SQLHandler.Scan(&items.Data, "id desc")
	return
}

// Create ...
func (repo *AppRepository) Create(item model.Item) (out model.Item, err error) {
	repo.SQLHandler.Create(item)
	return
}

// Update ...
func (repo *AppRepository) Update(value map[string]interface{}, query interface{}, args ...interface{}) (notification model.Notification, err error) {
	// NOTE: When update with struct, GORM will only update non-zero fields,
	// you might want to use map to update attributes or use Select to specify fields to update
	repo.SQLHandler.Update(&notification, value, query, args...)

	return
}
