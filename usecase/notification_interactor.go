package usecase

import (
	"context"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/horsewin/echo-playground-v2/domain/model/errors"
	"github.com/horsewin/echo-playground-v2/domain/repository"
)

// NotificationInteractor ...
type NotificationInteractor struct {
	NotificationRepository repository.NotificationRepositoryInterface
}

// GetNotifications ...
func (interactor *NotificationInteractor) GetNotifications(ctx context.Context, id string) (app model.Notifications, err error) {
	if id == "" {
		app, err = interactor.NotificationRepository.FindAll(ctx)
		if err != nil {
			return app, errors.NewBusinessError("10001E", err)
		}

	} else {
		app, err = interactor.NotificationRepository.Find(ctx, id)
		if err != nil {
			return app, errors.NewBusinessError("10001E", err)
		}
	}

	return
}

// MarkNotificationsRead ...
func (interactor *NotificationInteractor) MarkNotificationsRead(ctx context.Context, notificationId string) (err error) {
	var clause string
	var args map[string]interface{}

	if notificationId != "" {
		// 特定の通知のみを既読にする
		clause = "id = :id AND is_read = :is_read"
		args = map[string]interface{}{
			"id":      notificationId,
			"is_read": false,
		}
	} else {
		// 全ての未読通知を既読にする（従来の動作）
		clause = "is_read = :is_read"
		args = map[string]interface{}{"is_read": false}
	}

	err = interactor.NotificationRepository.Update(ctx, map[string]interface{}{"is_read": true}, clause, args)

	if err != nil {
		return errors.NewBusinessError("10001E", err)
	}

	return
}
