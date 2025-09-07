package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/horsewin/echo-playground-v2/domain/model/errors"

	"github.com/horsewin/echo-playground-v2/domain/repository"
	"github.com/horsewin/echo-playground-v2/interface/database"
	"github.com/horsewin/echo-playground-v2/usecase"
	"github.com/labstack/echo/v4"
)

// NotificationReadRequest JSON形式のリクエストボディをバインドするための構造体
type NotificationReadRequest struct {
	ID string `json:"id" form:"id"`
}

// NotificationHandler ...
type NotificationHandler struct {
	Interactor usecase.NotificationInteractor
}

// NewNotificationHandler ...
func NewNotificationHandler(sqlHandler database.SQLHandler) *NotificationHandler {
	return &NotificationHandler{
		Interactor: usecase.NotificationInteractor{
			NotificationRepository: &repository.NotificationRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetNotifications ...
func (handler *NotificationHandler) GetNotifications() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// スパンを作成
		ctx := c.Request().Context()
		tracer := otel.Tracer("notification-handler")
		ctx, span := tracer.Start(ctx, "GetNotifications",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		id := c.QueryParam("id")

		// スパンに属性を追加
		if id != "" {
			span.SetAttributes(
				attribute.String("id", id),
			)
		}

		// contextを渡す
		resJSON, err := handler.Interactor.GetNotifications(ctx, id)
		if err != nil {
			span.RecordError(err)
			return errors.NewEchoHTTPError(ctx, err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// PostNotificationsRead ...
func (handler *NotificationHandler) PostNotificationsRead() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// スパンを作成
		ctx := c.Request().Context()
		tracer := otel.Tracer("notification-handler")
		ctx, span := tracer.Start(ctx, "PostNotificationsRead",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		// JSONリクエストボディからnotificationIdを取得
		var req NotificationReadRequest
		if err := c.Bind(&req); err != nil {
			span.RecordError(err)
			return errors.NewEchoHTTPError(ctx, err)
		}
		notificationId := req.ID
		logger := zerolog.Ctx(ctx)
		logger.Debug().Str("notificationId", notificationId).Msg("Processing notification read request")

		// スパンに属性を追加
		span.SetAttributes(
			attribute.String("id", notificationId),
		)

		// contextを渡す
		err = handler.Interactor.MarkNotificationsRead(ctx, notificationId)
		if err != nil {
			span.RecordError(err)
			return errors.NewEchoHTTPError(ctx, err)
		}

		return c.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: "OK",
		})
	}
}
