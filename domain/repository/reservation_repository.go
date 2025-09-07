package repository

import (
	"context"
	"time"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/horsewin/echo-playground-v2/interface/database"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// ReservationRepositoryInterface ...
type ReservationRepositoryInterface interface {
	Create(ctx context.Context, input *model.Reservation) (err error)
	GetCountByPetID(ctx context.Context, petID string) (count int64, err error)
}

// ReservationRepository ...
type ReservationRepository struct {
	database.SQLHandler
}

const ReservationTable = "reservations"

// Create ...
func (repo *ReservationRepository) Create(ctx context.Context, input *model.Reservation) (err error) {
	// スパンを作成
	tracer := otel.Tracer("reservation-repository")
	ctx, span := tracer.Start(ctx, "ReservationRepository.Create",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("pet_id", input.PetId),
		attribute.String("user_id", input.UserId),
		attribute.String("reservation_date", input.ReservationDate),
	)

	// ドメインモデルをmapに変換
	rsvDatetime, err := time.Parse("20060102", input.ReservationDate)
	if err != nil {
		return
	}

	in := map[string]interface{}{
		"pet_id":    input.PetId,
		"user_id":   input.UserId,
		"email":     input.Email,
		"user_name": input.FullName,
		// yyyymmdd形式をdatetimeに変換
		"reservation_date_time": rsvDatetime,
		"status":                "pending", // デフォルトステータスを設定
	}

	// リポジトリモデルをDBに保存
	err = repo.SQLHandler.Create(ctx, in, ReservationTable)

	if err != nil {
		span.RecordError(err)
	}

	return
}

// GetCountByPetID ...
func (repo *ReservationRepository) GetCountByPetID(ctx context.Context, petID string) (count int64, err error) {
	// スパンを作成
	tracer := otel.Tracer("reservation-repository")
	_, span := tracer.Start(ctx, "ReservationRepository.GetCountByPetID",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("pet_id", petID),
	)

	var tmpCount int
	whereClause := "pet_id = :pet_id"
	whereArgs := map[string]interface{}{"pet_id": petID}
	err = repo.SQLHandler.Count(ctx, &tmpCount, ReservationTable, whereClause, whereArgs)
	count = int64(tmpCount)

	if err != nil {
		span.RecordError(err)
	}

	return
}
