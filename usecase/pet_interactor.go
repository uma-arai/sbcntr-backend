package usecase

import (
	"context"
	"math/rand"
	"time"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/horsewin/echo-playground-v2/domain/model/errors"
	"github.com/horsewin/echo-playground-v2/domain/repository"
	"github.com/horsewin/echo-playground-v2/utils"
	"github.com/jinzhu/copier"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// PetInteractor ...
type PetInteractor struct {
	PetRepository         repository.PetRepositoryInterface
	ReservationRepository repository.ReservationRepositoryInterface
	FavoriteRepository    repository.FavoriteRepositoryInterface
}

// GetPets ...
func (interactor *PetInteractor) GetPets(ctx context.Context, filter *model.PetFilter) (pets []model.Pet, err error) {
	// スパンを作成
	tracer := otel.Tracer("pet-interactor")
	ctx, span := tracer.Start(ctx, "PetInteractor.GetPets",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	if filter != nil {
		span.SetAttributes(
			attribute.String("filter.id", filter.ID),
			attribute.String("filter.name", filter.Name),
			attribute.String("filter.gender", filter.Gender),
			attribute.Float64("filter.price", filter.Price),
		)
	}

	// 性能テスト用：3回に1回だけCPU負荷を発生させる
	if rand.Intn(3) == 0 {
		induceCpuLoad()
	}

	// 性能テスト用：3回に1回だけレイテンシを発生させる
	if rand.Intn(3) == 0 {
		induceLatency()
	}

	// Repository層からデータを取得
	_app, err := interactor.PetRepository.Find(ctx, filter)
	if err != nil {
		return pets, errors.NewBusinessError("10001E", err)
	}

	// ドメインモデルに変換
	for _, p := range _app.Data {
		// 予約数を取得
		// Note: 意図的にN+1問題を起こしている箇所。OpenTelemetryで確認するため。
		reservationCount, countErr := interactor.ReservationRepository.GetCountByPetID(ctx, p.ID)
		if countErr != nil {
			utils.LogError("Failed to get reservation count: %v", countErr)
			// エラーが発生しても処理は継続
			reservationCount = 0
		}

		pets = append(pets, model.Pet{
			ID:       p.ID,
			Name:     p.Name,
			Breed:    p.Breed,
			Gender:   p.Gender,
			Price:    p.Price,
			ImageURL: p.ImageURL,
			Likes:    p.Likes,
			Shop: model.Shop{
				Name:     p.ShopName,
				Location: p.ShopLocation,
			},
			BirthDate:        p.BirthDate,
			ReferenceNumber:  p.ReferenceNumber,
			Tags:             p.Tags,
			ReservationCount: reservationCount,
		})
	}

	// 結果の属性を追加
	span.SetAttributes(
		attribute.Int("result_count", len(pets)),
	)

	return pets, nil
}

// UpdateLikeCount ...
func (interactor *PetInteractor) UpdateLikeCount(ctx context.Context, input *model.InputUpdateLikeRequest) (err error) {
	// スパンを作成
	tracer := otel.Tracer("pet-interactor")
	ctx, span := tracer.Start(ctx, "PetInteractor.UpdateLikeCount",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	if input != nil {
		span.SetAttributes(
			attribute.String("input.pet_id", input.PetId),
			attribute.String("input.user_id", input.UserId),
			attribute.Bool("input.value", input.Value),
		)
	}

	// like状態を取得
	favMap, err := interactor.FavoriteRepository.FindByUserId(ctx, input.UserId)
	if err != nil {
		return errors.NewBusinessError("10001E", err)
	}
	if favMap[input.PetId].Value == input.Value {
		return errors.NewBusinessError("00001I", nil)
	}

	// 現在のペットモデルを取得
	petData, err := interactor.PetRepository.Find(ctx, &model.PetFilter{ID: input.PetId})
	if err != nil {
		return errors.NewBusinessError("10001E", err)
	}
	var pet model.Pet
	for _, p := range petData.Data {
		if p.ID == input.PetId {
			err = copier.Copy(&pet, &p)
			if err != nil {
				return errors.NewBusinessError("10002E", err)
			}

			// マッピングしきれない構造は手動でコピー
			pet.Shop = model.Shop{
				Name:     p.ShopName,
				Location: p.ShopLocation,
			}
		}
	}

	// Like数を更新
	if input.Value {
		pet.Likes = pet.Likes + 1
	} else {
		pet.Likes = pet.Likes - 1
	}

	// TODO: トランザクションを使って更新処理を行う
	err = interactor.PetRepository.Update(ctx, &pet)
	if err != nil {
		return errors.NewBusinessError("10003E", err)
	}

	if input.Value {
		err = interactor.FavoriteRepository.Create(ctx, &model.Favorite{
			PetId:  input.PetId,
			UserId: input.UserId,
			Value:  input.Value,
		})
		if err != nil {
			return errors.NewBusinessError("10004E", err)
		}
	} else {
		err = interactor.FavoriteRepository.Delete(ctx, &model.Favorite{
			PetId:  input.PetId,
			UserId: input.UserId,
		})
		if err != nil {
			return errors.NewBusinessError("10005E", err)
		}
	}

	return
}

// CreateReservation ...
func (interactor *PetInteractor) CreateReservation(ctx context.Context, input *model.Reservation) (err error) {
	// スパンを作成
	tracer := otel.Tracer("pet-interactor")
	ctx, span := tracer.Start(ctx, "PetInteractor.CreateReservation",
		trace.WithSpanKind(trace.SpanKindInternal),
	)
	defer span.End()

	// 属性を追加
	if input != nil {
		span.SetAttributes(
			attribute.String("input.pet_id", input.PetId),
			attribute.String("input.user_id", input.UserId),
			attribute.String("input.reservation_date", input.ReservationDate),
		)
	}

	err = interactor.ReservationRepository.Create(ctx, input)
	if err != nil {
		return errors.NewBusinessError("10003E", err)
	}
	return
}

// induceCpuLoad ... 意図的にテスト用のCPU負荷を発生させる関数
func induceCpuLoad() {
	t := time.NewTimer(3 * time.Second)

	go func() {
		//nolint:staticcheck // 意図的に無限ループを作成してCPU負荷をシミュレートするためのコード
		for {
		}
	}()
	<-t.C
	t.Stop()
}

// induceLatency ... 意図的にテスト用のレイテンシを発生させる関数
func induceLatency() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	milliseconds := r.Intn(500) + 500
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}
