package usecase

import (
	"github.com/iselegant/cnappdemo/domain/model"
	"github.com/iselegant/cnappdemo/domain/repository"
	"github.com/iselegant/cnappdemo/utils"
)

// AppInteractor ...
type AppInteractor struct {
	AppRepository repository.AppRepositoryInterface
}

// GetAppInfo ...
func (interactor *AppInteractor) GetAppInfo(id string) (app model.App, err error) {

	app, err = interactor.AppRepository.Where(id)
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}
