package auth

import (
	"github.com/labstack/echo/v4"
)

type AuthController struct {
	echo *echo.Echo
	*LoginControllerImpl
	*CreateUserControllerImpl
}

func NewAuthController(echo *echo.Echo) *AuthController {
	return &AuthController{
		echo:                     echo,
		LoginControllerImpl:      newLoginControllerImpl(echo.Logger),
		CreateUserControllerImpl: newCreateUserController(echo.Logger),
	}
}

func (ctr *AuthController) Start() {
	r := ctr.echo.Group("/")
	r.POST("auth/login", ctr.Login)
	r.POST("auth/registration", ctr.Create)
}
