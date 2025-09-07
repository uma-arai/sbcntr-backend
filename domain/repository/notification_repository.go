package repository

import (
	"context"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/horsewin/echo-playground-v2/interface/database"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// NotificationRepositoryInterface ...
type NotificationRepositoryInterface interface {
	Find(ctx context.Context, id string) (account model.Notifications, err error)
	FindAll(ctx context.Context) (notifications model.Notifications, err error)
	Count(ctx context.Context, query string, args map[string]interface{}) (data model.NotificationCount, err error)
	Update(ctx context.Context, in map[string]interface{}, query string, args map[string]interface{}) (err error)
}

// NotificationRepository ....
type NotificationRepository struct {
	database.SQLHandler
}

const NotificationTable = "notifications"

// Find ...
func (repo *NotificationRepository) Find(ctx context.Context, id string) (notifications model.Notifications, err error) {
	// スパンを作成
	tracer := otel.Tracer("notification-repository")
	_, span := tracer.Start(ctx, "NotificationRepository.Find",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("id", id),
	)

	whereClause := "id = :id"
	whereArgs := map[string]interface{}{"id": id}
	err = repo.SQLHandler.Where(ctx, &notifications.Data, NotificationTable, whereClause, whereArgs)
	if err != nil {
		span.RecordError(err)
	}
	return
}

// FindAll ...
func (repo *NotificationRepository) FindAll(ctx context.Context) (notifications model.Notifications, err error) {
	// スパンを作成
	tracer := otel.Tracer("notification-repository")
	_, span := tracer.Start(ctx, "NotificationRepository.FindAll",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	err = repo.SQLHandler.Scan(ctx, &notifications.Data, NotificationTable, "id desc")
	if err != nil {
		span.RecordError(err)
	}
	return notifications, err
}

// Count ...
func (repo *NotificationRepository) Count(ctx context.Context, query string, args map[string]interface{}) (data model.NotificationCount, err error) {
	// スパンを作成
	tracer := otel.Tracer("notification-repository")
	_, span := tracer.Start(ctx, "NotificationRepository.Count",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("query", query),
	)

	var count int
	err = repo.SQLHandler.Count(ctx, &count, NotificationTable, query, args)
	if err != nil {
		span.RecordError(err)
	}
	return model.NotificationCount{Data: count}, err
}

// Update ...
func (repo *NotificationRepository) Update(ctx context.Context, in map[string]interface{}, query string, args map[string]interface{}) (err error) {
	// スパンを作成
	tracer := otel.Tracer("notification-repository")
	_, span := tracer.Start(ctx, "NotificationRepository.Update",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("query", query),
	)

	// 新しいSQLHandlerのシグネチャに合わせて、SET用とWHERE用のパラメータを分離して渡す
	err = repo.SQLHandler.Update(ctx, in, NotificationTable, query, args)
	if err != nil {
		span.RecordError(err)
	}
	return
}
