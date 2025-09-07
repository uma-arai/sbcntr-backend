package infrastructure

import (
	"context"
	"os"
	"time"

	handlers "github.com/horsewin/echo-playground-v2/handler"
	"github.com/horsewin/echo-playground-v2/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog"
	zerologlog "github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

const (
	// TODO: 適切な名前に変更をする
	projectName = "echo-playground-v2"
)

// configureOpenTelemetry OpenTelemetryの設定を行う
func configureOpenTelemetry(apiConfig *utils.APIConfig, logger zerolog.Logger) {
	logger.Info().Msgf("configureOpenTelemetry start : %v", apiConfig.EnableTracing)

	ctx := context.Background()
	tp, err := SetupOpenTelemetry(ctx, projectName, "2.14.0", logger, apiConfig)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to configure OpenTelemetry")
		return
	}

	if tp != nil {
		logger.Info().Msg("OpenTelemetry configured successfully")
	} else {
		logger.Info().Msg("OpenTelemetry is disabled")
	}
}

// setupRequestLogger リクエストロガーミドルウェアを設定
func setupRequestLogger(logger zerolog.Logger) echo.MiddlewareFunc {
	return middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:       true,
		LogStatus:    true,
		LogUserAgent: true,
		LogMethod:    true,
		LogLatency:   true,
		LogRemoteIP:  true,
		LogRequestID: true,
		LogError:     true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			event := logger.Info()
			if v.Status >= 500 {
				event = logger.Error()
			} else if v.Status >= 400 {
				event = logger.Warn()
			}

			rid := c.Response().Header().Get(echo.HeaderXRequestID)

			event.Str("request_id", rid).
				Str("time", time.Now().Format(time.RFC3339Nano)).
				Str("remote_ip", v.RemoteIP).
				Str("method", v.Method).
				Str("latency", v.Latency.String()).
				Str("uri", v.URI).
				Int("status", v.Status).
				Str("user_agent", v.UserAgent)

			// エラーが発生した場合、エラー情報をログに記録
			if v.Error != nil {
				event.Err(v.Error)
				event.Msg("failed api request")

				return v.Error
			}

			event.Msg("completed api request")

			return nil
		},
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/healthcheck"
		},
	})
}

// setupOpenTelemetryMiddleware OpenTelemetryミドルウェアを設定
func setupOpenTelemetryMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			rid := c.Response().Header().Get(echo.HeaderXRequestID)
			// Create request-specific logger with request details
			reqLogger := zerologlog.With().
				Str("request_id", rid).
				Str("method", c.Request().Method).
				Str("path", c.Request().URL.Path).
				Logger()

			// healthcheckはトレース作成を行わない
			if c.Path() == "/healthcheck" {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			ctx := req.Context()

			// OpenTelemetryトレーサーを取得
			tracer := otel.Tracer(projectName)

			// スパンを開始（サーバースパンとして）
			ctx, span := tracer.Start(ctx, c.Request().URL.Path,
				trace.WithSpanKind(trace.SpanKindServer),
				trace.WithAttributes(
					attribute.String("http.method", c.Request().Method),
					attribute.String("http.url", c.Request().URL.String()),
					attribute.String("http.target", c.Request().URL.Path),
					attribute.String("http.host", c.Request().Host),
					attribute.String("http.scheme", c.Request().URL.Scheme),
					attribute.String("http.user_agent", c.Request().UserAgent()),
					attribute.String("http.request_id", rid),
				),
			)
			defer func() {
				defer span.End()
			}()

			// リクエストのコンテキストを更新
			c.SetRequest(req.WithContext(reqLogger.WithContext(ctx)))

			// 次のハンドラーを呼び出し
			err := next(c)

			// レスポンスステータスを記録
			span.SetAttributes(
				attribute.Int("http.status_code", res.Status),
			)

			// エラーが発生した場合、スパンにエラー情報を記録
			if err != nil {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())
			} else {
				span.SetStatus(codes.Ok, "")
			}

			return err
		}
	}
}

// registerRoutes ルートを登録する
func registerRoutes(e *echo.Echo) {
	healthCheckHandler := handlers.NewHealthCheckHandler()
	helloWorldHandler := handlers.NewHelloWorldHandler()

	// ---------------------------
	// APIルートの定義
	// ---------------------------
	e.GET("/", healthCheckHandler.HealthCheck())
	e.GET("/healthcheck", healthCheckHandler.HealthCheck())
	e.GET("/v1/helloworld", helloWorldHandler.SayHelloWorld())
	e.GET("/v1/helloworld/error", helloWorldHandler.SayError())
	if os.Getenv("DB_CONN") == "1" {
		sqlHandler := NewSQLHandler()
		petHandler := handlers.NewPetHandler(sqlHandler)
		notificationHandler := handlers.NewNotificationHandler(sqlHandler)

		e.GET("/v1/pets", petHandler.GetPets())
		e.POST("/v1/pets/:id/like", petHandler.UpdateLike())
		e.POST("/v1/pets/:id/reservation", petHandler.Reservation())

		e.GET("/v1/notifications", notificationHandler.GetNotifications())
		e.POST("/v1/notifications/read", notificationHandler.PostNotificationsRead())

	}
}

// Router ...
func Router() *echo.Echo {
	// Setup
	logger := zerolog.New(os.Stdout)
	e := echo.New()
	apiConfig := utils.NewAPIConfig()

	// Configure OpenTelemetry
	configureOpenTelemetry(apiConfig, logger)

	// Configure Echo settings
	e.HideBanner = true
	e.HidePort = false

	// Setup middlewares
	setupMiddlewares(e, logger)

	// Register routes
	registerRoutes(e)

	return e
}

// setupMiddlewares ミドルウェアを設定
func setupMiddlewares(e *echo.Echo, logger zerolog.Logger) {
	// リクエストIDの生成
	e.Use(middleware.RequestID())

	// ログ出力設定
	e.Use(setupRequestLogger(logger))

	// OpenTelemetryミドルウェア
	e.Use(setupOpenTelemetryMiddleware())

	// recoveryミドルウェアの設定
	e.Use(middleware.Recover())
}
