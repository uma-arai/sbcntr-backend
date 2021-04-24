package handlers

import (
	"net/http"

	"github.com/labstack/echo"
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
		var fav bool
		favorite := c.QueryParam("favorite")
		if favorite == "" {
			fav = false
		} else {
			fav = true
		}

		//err = utils.HeaderCheck(c, echo.HeaderContentType, echo.MIMEApplicationJSON)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		//err = utils.ClientIDCheck(c)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		resJSON, err := handler.Interactor.GetItems(fav)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}

// CreateItem ...
func (handler *AppHandler) CreateItem() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		var fav bool
		favorite := c.QueryParam("favorite")
		if favorite == "" {
			fav = false
		} else {
			fav = true
		}

		resJSON, err := handler.Interactor.GetItems(fav)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(http.StatusOK, resJSON)
	}
}
