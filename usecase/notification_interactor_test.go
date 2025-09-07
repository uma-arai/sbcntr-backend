package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/horsewin/echo-playground-v2/domain/model"
	business_errors "github.com/horsewin/echo-playground-v2/domain/model/errors"
)

// MockNotificationRepository はテスト用のモックリポジトリ
type MockNotificationRepository struct {
	FindResult    model.Notifications
	FindError     error
	FindAllResult model.Notifications
	FindAllError  error
	UpdateError   error
}

func (m *MockNotificationRepository) Find(ctx context.Context, id string) (model.Notifications, error) {
	return m.FindResult, m.FindError
}

func (m *MockNotificationRepository) FindAll(ctx context.Context) (model.Notifications, error) {
	return m.FindAllResult, m.FindAllError
}

func (m *MockNotificationRepository) Count(ctx context.Context, query string, args map[string]interface{}) (model.NotificationCount, error) {
	return model.NotificationCount{}, nil
}

func (m *MockNotificationRepository) Update(ctx context.Context, in map[string]interface{}, query string, args map[string]interface{}) error {
	return m.UpdateError
}

func TestNotificationInteractor_GetNotifications_WithID(t *testing.T) {
	mockRepo := &MockNotificationRepository{
		FindResult: model.Notifications{
			Data: []model.Notification{
				{ID: 1, UserId: "user1", Title: "Test Notification", Message: "Test message", IsRead: false},
			},
		},
		FindError: nil,
	}

	interactor := &NotificationInteractor{
		NotificationRepository: mockRepo,
	}

	ctx := context.Background()
	result, err := interactor.GetNotifications(ctx, "1")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result.Data) != 1 {
		t.Errorf("expected 1 notification but got %d", len(result.Data))
	}

	if result.Data[0].ID != 1 {
		t.Errorf("expected notification ID 1 but got %d", result.Data[0].ID)
	}
}

func TestNotificationInteractor_GetNotifications_WithoutID(t *testing.T) {
	mockRepo := &MockNotificationRepository{
		FindAllResult: model.Notifications{
			Data: []model.Notification{
				{ID: 1, UserId: "user1", Title: "Test Notification 1", Message: "Test message 1", IsRead: false},
				{ID: 2, UserId: "user2", Title: "Test Notification 2", Message: "Test message 2", IsRead: true},
			},
		},
		FindAllError: nil,
	}

	interactor := &NotificationInteractor{
		NotificationRepository: mockRepo,
	}

	ctx := context.Background()
	result, err := interactor.GetNotifications(ctx, "")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}

	if len(result.Data) != 2 {
		t.Errorf("expected 2 notifications but got %d", len(result.Data))
	}
}

func TestNotificationInteractor_GetNotifications_FindError(t *testing.T) {
	mockRepo := &MockNotificationRepository{
		FindError: errors.New("database error"),
	}

	interactor := &NotificationInteractor{
		NotificationRepository: mockRepo,
	}

	ctx := context.Background()
	_, err := interactor.GetNotifications(ctx, "1")

	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	var be business_errors.BusinessError
	if !errors.As(err, &be) {
		t.Errorf("expected BusinessError but got %T", err)
		return
	}

	if be.Code() != "10001E" {
		t.Errorf("expected error code 10001E but got %s", be.Code())
	}
}

func TestNotificationInteractor_MarkNotificationsRead_WithID(t *testing.T) {
	mockRepo := &MockNotificationRepository{
		UpdateError: nil,
	}

	interactor := &NotificationInteractor{
		NotificationRepository: mockRepo,
	}

	ctx := context.Background()
	err := interactor.MarkNotificationsRead(ctx, "1")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNotificationInteractor_MarkNotificationsRead_WithoutID(t *testing.T) {
	mockRepo := &MockNotificationRepository{
		UpdateError: nil,
	}

	interactor := &NotificationInteractor{
		NotificationRepository: mockRepo,
	}

	ctx := context.Background()
	err := interactor.MarkNotificationsRead(ctx, "")

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestNotificationInteractor_MarkNotificationsRead_UpdateError(t *testing.T) {
	mockRepo := &MockNotificationRepository{
		UpdateError: errors.New("database error"),
	}

	interactor := &NotificationInteractor{
		NotificationRepository: mockRepo,
	}

	ctx := context.Background()
	err := interactor.MarkNotificationsRead(ctx, "1")

	if err == nil {
		t.Errorf("expected error but got nil")
		return
	}

	var be business_errors.BusinessError
	if !errors.As(err, &be) {
		t.Errorf("expected BusinessError but got %T", err)
		return
	}

	if be.Code() != "10001E" {
		t.Errorf("expected error code 10001E but got %s", be.Code())
	}
}
