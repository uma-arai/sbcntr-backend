package infrastructure

import (
	"context"
	"fmt"
	"github.com/horsewin/echo-playground-v2/utils"
	"github.com/rs/zerolog"
	"os"

	awsxray "go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

var Tracer trace.Tracer

// SetupOpenTelemetry OpenTelemetryのトレーサーを設定
func SetupOpenTelemetry(ctx context.Context, serviceName string, serviceVersion string, logger zerolog.Logger, apiConfig *utils.APIConfig) (*sdktrace.TracerProvider, error) {
	if !apiConfig.EnableTracing {
		// トレーシングが無効の場合はnoopトレーサープロバイダーを返す
		tp := noop.NewTracerProvider()
		otel.SetTracerProvider(tp)
		Tracer = tp.Tracer(serviceName)
		return nil, nil
	}

	// OTLPエクスポーターのエンドポイント設定
	exporterEndpoint := os.Getenv("OTEL_EXPORTER_OTLP_ENDPOINT")
	if exporterEndpoint == "" {
		exporterEndpoint = "127.0.0.1:4318" // デフォルトのOTLP/HTTPエンドポイント
	}

	// HTTPトレースエクスポーターを作成
	var exporterOptions []otlptracehttp.Option
	exporterOptions = append(exporterOptions, otlptracehttp.WithEndpoint(exporterEndpoint))
	exporterOptions = append(exporterOptions, otlptracehttp.WithInsecure())

	// HTTPトレースエクスポーターを作成
	exporter, err := otlptracehttp.New(ctx,
		otlptracehttp.WithEndpoint(exporterEndpoint),
		otlptracehttp.WithInsecure(), // サイドカーコンテナに向けて送出するためHTTPモードを使用
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create OTLP trace exporter: %w", err)
	}

	// X-Rayリモートサンプラーを作成
	sampler := sdktrace.AlwaysSample()
	// Note: xray.NewRemoteSamplerのAPIが変更されたため、一時的にAlwaysSampleを使用

	// リソース検出器を設定
	res := resource.Default()

	// サービス情報をリソースに追加
	serviceResource, err := resource.New(ctx,
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.ServiceVersionKey.String(serviceVersion),
		),
	)
	if err == nil {
		res, _ = resource.Merge(res, serviceResource)
	}

	// トレーサープロバイダーを作成（X-Ray ID generatorを含む）
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithSampler(sampler),
		sdktrace.WithResource(res),
		sdktrace.WithIDGenerator(awsxray.NewIDGenerator()),
	)

	// グローバルトレーサープロバイダーを設定
	otel.SetTracerProvider(tp)

	// X-Ray伝搬を設定（AWSサービスとの互換性のため）
	otel.SetTextMapPropagator(awsxray.Propagator{})

	// アプリケーション用のトレーサーを作成
	Tracer = tp.Tracer(serviceName)

	// AWS_XRAY_CONTEXT_MISSING環境変数を設定（X-Ray互換性のため）
	os.Setenv("AWS_XRAY_CONTEXT_MISSING", "LOG_ERROR")

	return tp, nil
}
