package usecase

import (
	"github.com/uma-arai/sbcntr-backend/domain/model"
	"github.com/uma-arai/sbcntr-backend/domain/repository"
	"github.com/uma-arai/sbcntr-backend/utils"
)

// NotificationInteractor ...
type NotificationInteractor struct {
	NotificationRepository repository.NotificationRepositoryInterface
}

// GetNotifications ...
func (interactor *NotificationInteractor) GetNotifications(id string) (app model.Notifications, err error) {

	app, err = interactor.NotificationRepository.Where(id)
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}
