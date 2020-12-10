package handlers

import (
	"github.com/iselegant/cnappdemo/domain/repository"
	"github.com/iselegant/cnappdemo/interface/database"
	"github.com/iselegant/cnappdemo/usecase"
	"github.com/iselegant/cnappdemo/utils"
	"github.com/labstack/echo"
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

// GetAppInfo ...
func (handler *AppHandler) GetAppInfo() echo.HandlerFunc {

	return func(c echo.Context) (err error) {

		id := c.QueryParam("id")
		if id == "" {
			return c.JSON(400, "Invalid parameter id.")
		}

		//err = utils.HeaderCheck(c, echo.HeaderContentType, echo.MIMEApplicationJSON)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		//err = utils.ClientIDCheck(c)
		//if err != nil {
		//	return utils.GetErrorMassage(c, "en", err)
		//}

		resJSON, err := handler.Interactor.GetAppInfo(id)
		if err != nil {
			return utils.GetErrorMassage(c, "en", err)
		}

		return c.JSON(200, resJSON)
	}
}
