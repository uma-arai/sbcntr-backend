package usecase

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/domain/repository"
	"github.com/uma-arai/sbcntr-backend/utils"
)

// AppInteractor ...
type AppInteractor struct {
	AppRepository repository.AppRepositoryInterface
}

// GetItems ...
func (interactor *AppInteractor) GetItems(favorite string) (app model.Items, err error) {
	var query string
	var args interface{}
	if favorite == "true" {
		query = "favorite = ?"
		args = true
	} else if favorite == "false" {
		query = "favorite = ?"
		args = false
	} else {
		query = ""
		args = ""
	}

	app, err = interactor.AppRepository.Find(query, args)
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}

// CreateItem ...
func (interactor *AppInteractor) CreateItem(input model.Item) (response model.Response, err error) {
	response, err = interactor.AppRepository.Create(input)
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}

// UpdateFavoriteAttr ...
func (interactor *AppInteractor) UpdateFavoriteAttr(input model.Item) (item model.Item, err error) {
	item, err = interactor.AppRepository.Update(map[string]interface{}{"Favorite": input.Favorite}, "id = ?", input.ID)

	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}
