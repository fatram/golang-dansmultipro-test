package position

import (
	"net/http"

	"github.com/fatram/golang-dansmultipro-test/domain/model"
	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
	service "github.com/fatram/golang-dansmultipro-test/service/position"
	"github.com/labstack/echo/v4"
)

type GetPositionControllerImpl struct {
	Service service.PositionService
}

func (ctr *GetPositionControllerImpl) Get(c echo.Context) (err error) {
	_ = model.UserFromJWTContext(c.Get("user"))
	ctx := c.Request().Context()
	id := c.Param("id")
	data, err := ctr.Service.Get(ctx, id)
	if err != nil {
		c.Logger().Errorf("Error on GetPositionControllerImpl.Get: %s", err.Error())
		return err
	}
	return c.JSON(http.StatusOK, data)
}

func newGetPositionController(logger genlog.Logger) *GetPositionControllerImpl {
	return &GetPositionControllerImpl{
		Service: *service.LoadPositionService(logger),
	}
}

type GetAllPositionControllerImpl struct {
	Service service.PositionService
	Filter  model.PositionFilter
}

func (ctr *GetAllPositionControllerImpl) GetAll(c echo.Context) (err error) {
	_ = model.UserFromJWTContext(c.Get("user"))
	filter := ctr.Filter

	if err := c.Bind(&filter); err != nil {
		c.Logger().Errorf("Error on GetAllPositionControllerImpl.GetAll: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}

	if filter.PageNumber != nil {
		if *filter.PageNumber < 1 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid page").SetInternal(err)
		}
	}
	if filter.PageSize != nil {
		if *filter.PageSize < 1 {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid limit").SetInternal(err)
		}
	}

	data, _, err := ctr.Service.GetAll(c.Request().Context(), filter)
	if err != nil {
		c.Logger().Errorf("Error on GetAllPositionControllerImpl.GetAllPosition: %s", err.Error())
		return err
	}
	dataInterface := make([]interface{}, len(data))
	for i, item := range data {
		dataInterface[i] = item
	}
	return c.JSON(http.StatusOK, dataInterface)
}

func newGetAllPositionController(logger genlog.Logger) *GetAllPositionControllerImpl {
	return &GetAllPositionControllerImpl{
		Service: *service.LoadPositionService(logger),
		Filter:  *new(model.PositionFilter),
	}
}
