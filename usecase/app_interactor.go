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

// GetAppInfo ...
func (interactor *AppInteractor) GetAppInfo(id string) (app model.App, err error) {

	app, err = interactor.AppRepository.Where(id)
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}
