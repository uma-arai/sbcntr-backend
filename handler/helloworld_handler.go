package handlers

import (
	"net/http"

	"github.com/rs/zerolog"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/labstack/echo/v4"
)

// HelloWorldHandler ...
type HelloWorldHandler struct {
}

// NewHelloWorldHandler ...
func NewHelloWorldHandler() *HelloWorldHandler {
	return &HelloWorldHandler{}
}

// SayHelloWorld ...
func (handler *HelloWorldHandler) SayHelloWorld() echo.HandlerFunc {
	body := &model.Hello{
		Message: "Hello world",
	}

	// ドメインモデルをJSONにして返却
	return func(c echo.Context) error {
		// Get logger from context
		ctx := c.Request().Context()
		logger := zerolog.Ctx(ctx)

		// スパンを作成
		tracer := otel.Tracer("helloworld-handler")
		_, span := tracer.Start(ctx, "SayHelloWorld",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		// スパンに属性を追加
		span.SetAttributes(
			attribute.String("message", body.Message),
		)

		return c.JSON(http.StatusOK, model.APIResponse{
			Data: body,
		})
	}
}

// SayError ...
func (handler *HelloWorldHandler) SayError() echo.HandlerFunc {
	return func(c echo.Context) error {
		// Return a client-friendly error using echo.NewHTTPError
		// This will be handled by our custom error handler
		return echo.NewHTTPError(http.StatusNotImplemented, "Invalid endpoint")
	}
}
