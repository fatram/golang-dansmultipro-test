package auth

import (
	"net/http"

	"github.com/fatram/golang-dansmultipro-test/domain/model"
	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
	service "github.com/fatram/golang-dansmultipro-test/service/auth"
	"github.com/labstack/echo/v4"
)

type LoginControllerImpl struct {
	Service  service.AuthService
	Bind     model.BindFunc
	Validate model.ValideFunc
}

func newLoginControllerImpl(logger genlog.Logger) *LoginControllerImpl {
	return &LoginControllerImpl{
		Service:  *service.LoadAuthService(logger),
		Bind:     model.BindUserLogin,
		Validate: model.ValidateUserLogin,
	}
}

func (ctr *LoginControllerImpl) Login(c echo.Context) (err error) {
	data, err := ctr.Bind(c)
	if err != nil {
		c.Logger().Errorf("Error on LoginControllerImpl.Login: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	if err := ctr.Validate(c, data); err != nil {
		c.Logger().Errorf("Validation error on LoginControllerImpl.Login: %s", err.Error())
		return err
	}

	token, err := ctr.Service.Login(c.Request().Context(), data)
	if err != nil {
		c.Logger().Errorf("Error on LoginControllerImpl.Login: %s", err.Error())
		return err
	}
	return c.JSON(http.StatusOK, model.TokenResponse{Token: token})
}

type CreateUserControllerImpl struct {
	Service  service.AuthService
	Bind     model.BindFunc
	Validate model.ValideFunc
}

func (ctr *CreateUserControllerImpl) Create(c echo.Context) (err error) {
	_ = model.UserFromJWTContext(c.Get("user"))
	data, err := ctr.Bind(c)
	if err != nil {
		c.Logger().Errorf("Error on CreateControllerImpl.Create: %s", err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, "invalid request body").SetInternal(err)
	}
	if err := ctr.Validate(c, data); err != nil {
		c.Logger().Errorf("Validation error on CreateControllerImpl.Create: %s", err.Error())
		return err
	}

	id, err := ctr.Service.Create(c.Request().Context(), data)
	if err != nil {
		c.Logger().Errorf("Error on CreateControllerImpl.Create: %s", err.Error())
		return err
	}
	response := model.CreateResponse{
		ID: id,
	}
	return c.JSON(http.StatusCreated, response)
}

func newCreateUserController(logger genlog.Logger) *CreateUserControllerImpl {
	return &CreateUserControllerImpl{
		Service:  *service.LoadAuthService(logger),
		Bind:     model.BindUserCreate,
		Validate: model.ValidateUserCreate,
	}
}
