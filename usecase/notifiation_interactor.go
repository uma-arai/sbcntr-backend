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
func (interactor *NotificationInteractor) GetNotifications() (app model.Notifications, err error) {

	app, err = interactor.NotificationRepository.FindAll()
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}

// GetNotificationCount ...
func (interactor *NotificationInteractor) GetNotificationCount() (count model.NotificationCount, err error) {

	count, err = interactor.NotificationRepository.Count()
	if err != nil {
		err = utils.SetErrorMassage("10001E")
		return
	}

	return
}
