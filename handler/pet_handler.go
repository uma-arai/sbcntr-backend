package handlers

import (
	"net/http"

	"github.com/horsewin/echo-playground-v2/domain/model"
	"github.com/labstack/echo/v4"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"

	"github.com/horsewin/echo-playground-v2/domain/repository"
	"github.com/horsewin/echo-playground-v2/interface/database"
	"github.com/horsewin/echo-playground-v2/usecase"
)

// PetHandler ...
type PetHandler struct {
	Interactor usecase.PetInteractor
}

// NewPetHandler ...
func NewPetHandler(sqlHandler database.SQLHandler) *PetHandler {
	return &PetHandler{
		Interactor: usecase.PetInteractor{
			PetRepository: &repository.PetRepository{
				SQLHandler: sqlHandler,
			},
			ReservationRepository: &repository.ReservationRepository{
				SQLHandler: sqlHandler,
			},
			FavoriteRepository: &repository.FavoriteRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetPets ...
func (handler *PetHandler) GetPets() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// スパンを作成
		ctx := c.Request().Context()
		tracer := otel.Tracer("pet-handler")
		ctx, span := tracer.Start(ctx, "GetPets",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		filter := new(model.PetFilter)
		if err := c.Bind(filter); err != nil {
			span.RecordError(err)
			return err
		}

		// スパンに属性を追加
		span.SetAttributes(
			attribute.String("filter.id", filter.ID),
			attribute.String("filter.name", filter.Name),
			attribute.String("filter.gender", filter.Gender),
			attribute.Float64("filter.price", filter.Price),
		)

		// Pass the context with span to the interactor
		res, err := handler.Interactor.GetPets(ctx, filter)
		if err != nil {
			span.RecordError(err)
			return err
		}

		// 結果の属性を追加
		span.SetAttributes(
			attribute.Int("result_count", len(res)),
		)

		// resの中身をJSONにして返却
		resJSON := model.APIResponse{
			Data: res,
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// UpdateLike ...
func (handler *PetHandler) UpdateLike() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// スパンを作成
		ctx := c.Request().Context()
		tracer := otel.Tracer("pet-handler")
		ctx, span := tracer.Start(ctx, "UpdateLike",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		// パスパラメータ "id" の値を取得
		id := c.Param("id")
		if id == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No id param found",
			})
		}

		// Bindでリクエストの中身をiに詰める
		input := new(struct {
			UserId string `json:"user_id"`
			Value  bool   `json:"value"`
		})
		if err = c.Bind(input); err != nil {
			span.RecordError(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// スパンに属性を追加
		span.SetAttributes(
			attribute.String("pet_id", id),
			attribute.String("user_id", input.UserId),
			attribute.Bool("value", input.Value),
		)

		// UseCaseの実行
		err = handler.Interactor.UpdateLikeCount(ctx, &model.InputUpdateLikeRequest{
			PetId:  id,
			UserId: input.UserId,
			Value:  input.Value,
		})

		if err != nil {
			span.RecordError(err)
			return err
		}

		return c.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: "OK",
		})
	}
}

// Reservation ...
func (handler *PetHandler) Reservation() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		// スパンを作成
		ctx := c.Request().Context()
		tracer := otel.Tracer("pet-handler")
		ctx, span := tracer.Start(ctx, "Reservation",
			trace.WithSpanKind(trace.SpanKindInternal),
		)
		defer span.End()

		// パスパラメータ "id" の値を取得
		petId := c.Param("id")
		if petId == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No id param found",
			})
		}

		// Bindでリクエストの中身をinputにつめる
		input := new(struct {
			UserId          string `json:"user_id"`
			Email           string `json:"email"`
			FullName        string `json:"full_name"`
			ReservationDate string `json:"reservation_date"`
		})
		if err = c.Bind(input); err != nil {
			span.RecordError(err)
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		// スパンに属性を追加
		span.SetAttributes(
			attribute.String("pet_id", petId),
			attribute.String("user_id", input.UserId),
			attribute.String("reservation_date", input.ReservationDate),
		)

		// UseCaseの実行
		err = handler.Interactor.CreateReservation(ctx, &model.Reservation{
			PetId:           petId,
			UserId:          input.UserId,
			Email:           input.Email,
			FullName:        input.FullName,
			ReservationDate: input.ReservationDate,
		})

		if err != nil {
			span.RecordError(err)
			return err
		}

		return c.JSON(http.StatusOK, model.Response{
			Code:    http.StatusOK,
			Message: "Created",
		})
	}
}
