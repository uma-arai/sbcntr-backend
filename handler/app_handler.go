package handlers

import (
	"net/http"
	"time"

	"github.com/uma-arai/sbcntr-backend/domain/model"

	"github.com/labstack/echo/v4"
	"github.com/uma-arai/sbcntr-backend/domain/repository"
	"github.com/uma-arai/sbcntr-backend/interface/database"
	"github.com/uma-arai/sbcntr-backend/usecase"
	"github.com/uma-arai/sbcntr-backend/utils"
)

// AppHandler ...
type AppHandler struct {
	Interactor usecase.AppInteractor
}

// NewAppHandler ...
func NewAppHandler(sqlHandler database.SQLHandler) *AppHandler {
	return &AppHandler{
		Interactor: usecase.AppInteractor{
			AppRepository: &repository.AppRepository{
				SQLHandler: sqlHandler,
			},
		},
	}
}

// GetItems ...
func (handler *AppHandler) GetItems() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		favorite := c.QueryParam("favorite")
		resJSON, err := handler.Interactor.GetItems(favorite)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// CreateItem ...
func (handler *AppHandler) CreateItem() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		i := new(model.Item)
		if err = c.Bind(i); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		input := model.Item{
			Title:     i.Title,
			Name:      i.Name,
			Favorite:  false,
			Img:       i.Img,
			CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
			UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		}

		if input.Name == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No name param found",
			})
		}

		if input.Title == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No title param found",
			})

		}

		if input.Img == "" {
			return c.JSON(http.StatusBadRequest, model.Response{
				Message: "No img param found",
			})
		}

		resJSON, err := handler.Interactor.CreateItem(input)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// UpdateFavoriteAttr ...
func (handler *AppHandler) UpdateFavoriteAttr() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		input := new(model.Item)
		if err = c.Bind(input); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		resJSON, err := handler.Interactor.UpdateFavoriteAttr(*input)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}
