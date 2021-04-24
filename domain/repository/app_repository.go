package repository

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/interface/database"
)

// AppRepositoryInterface ...
type AppRepositoryInterface interface {
	FindAll() (items model.Items, err error)
	Find(query interface{}, args ...interface{}) (items model.Items, err error)
	Create(input model.Item) (out model.Response, err error)
	Update(value map[string]interface{}, query interface{}, args ...interface{}) (item model.Item, err error)
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

// Find ...
func (repo *AppRepository) Find(query interface{}, args ...interface{}) (items model.Items, err error) {
	repo.SQLHandler.Where(&items.Data, query, args...)
	return
}

// Create ...
func (repo *AppRepository) Create(input model.Item) (out model.Response, err error) {
	_, err = repo.SQLHandler.Create(&input)

	if err != nil {
		return model.Response{
			Code:    400,
			Message: "Create error",
		}, err
	}

	return model.Response{
		Code:    200,
		Message: "OK",
	}, nil
}

// Update ...
func (repo *AppRepository) Update(value map[string]interface{}, query interface{}, args ...interface{}) (item model.Item, err error) {
	// NOTE: When update with struct, GORM will only update non-zero fields,
	// you might want to use map to update attributes or use Select to specify fields to update
	repo.SQLHandler.Update(&item, value, query, args...)

	return
}
