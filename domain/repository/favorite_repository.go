package repository

import (
	"context"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/horsewin/echo-playground-v2/interface/database"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

type favorite struct {
	ID     string `db:"id"`
	PetId  string `db:"pet_id"`
	UserId string `db:"user_id"`
}

type favorites struct {
	Data []favorite
}

// FavoriteRepositoryInterface ...
type FavoriteRepositoryInterface interface {
	FindByUserId(ctx context.Context, userId string) (favorites map[string]model.Favorite, err error)
	Create(ctx context.Context, input *model.Favorite) (err error)
	Delete(ctx context.Context, input *model.Favorite) (err error)
}

// FavoriteRepository ...
type FavoriteRepository struct {
	database.SQLHandler
}

const FavoriteTable = "favorites"

// FindByUserId ...
func (f FavoriteRepository) FindByUserId(ctx context.Context, userId string) (favMap map[string]model.Favorite, err error) {
	// スパンを作成
	tracer := otel.Tracer("favorite-repository")
	_, span := tracer.Start(ctx, "FavoriteRepository.FindByUserId",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("user_id", userId),
	)

	// inputをmapに変換
	in := map[string]interface{}{"user_id": userId}

	// リポジトリモデルをDBから取得
	var _favorites favorites
	err = f.SQLHandler.Where(ctx, &_favorites.Data, FavoriteTable, "user_id = :user_id", in)
	if err != nil {
		span.RecordError(err)
		return
	}

	// 結果の属性を追加
	span.SetAttributes(
		attribute.Int("result_count", len(_favorites.Data)),
	)

	// ドメインモデルに変換
	favMap = make(map[string]model.Favorite)
	for _, _favorite := range _favorites.Data {
		favMap[_favorite.PetId] = model.Favorite{
			Id:     _favorite.ID,
			PetId:  _favorite.PetId,
			UserId: _favorite.UserId,
			Value:  true,
		}
	}

	return
}

// Create ...
func (f FavoriteRepository) Create(ctx context.Context, input *model.Favorite) (err error) {
	// スパンを作成
	tracer := otel.Tracer("favorite-repository")
	_, span := tracer.Start(ctx, "FavoriteRepository.Create",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("pet_id", input.PetId),
		attribute.String("user_id", input.UserId),
		attribute.Bool("value", input.Value),
	)

	// ドメインモデルをmapに変換
	in := map[string]interface{}{"pet_id": input.PetId, "user_id": input.UserId}

	// リポジトリモデルをDBに保存
	err = f.SQLHandler.Create(ctx, in, FavoriteTable)

	if err != nil {
		span.RecordError(err)
	}

	return
}

// Delete ...
func (f FavoriteRepository) Delete(ctx context.Context, input *model.Favorite) (err error) {
	// スパンを作成
	tracer := otel.Tracer("favorite-repository")
	_, span := tracer.Start(ctx, "FavoriteRepository.Delete",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	span.SetAttributes(
		attribute.String("pet_id", input.PetId),
		attribute.String("user_id", input.UserId),
	)

	// ドメインモデルをリポジトリモデルに変換
	_favorite := favorite{
		PetId:  input.PetId,
		UserId: input.UserId,
	}

	// リポジトリモデルをmapに変換
	in := map[string]interface{}{"pet_id": _favorite.PetId, "user_id": _favorite.UserId}

	// リポジトリモデルをDBから削除
	err = f.SQLHandler.Delete(ctx, in, FavoriteTable)

	if err != nil {
		span.RecordError(err)
	}

	return
}
