package auth

import (
	"github.com/labstack/echo/v4"
)

// Login godoc
// @Summary			Login user
// @Description		Masuk dengan akun user
// @Tags         	Auth
// @Accept       	json
// @Accept       	x-www-form-urlencoded
// @Produce      	json
// @Param			body			body	model.UserLogin	true	"body"
// @Success      	200  {object}   model.TokenResponse
// @Failure      	400,401,500  {object}  	pkg.Error
// @Router       	/auth/login [post]
func (ctr *AuthController) Login(c echo.Context) (err error) {
	return ctr.LoginControllerImpl.Login(c)
}

// Register user godoc
// @Summary			Register user
// @Description		Membuat user
// @Tags         	user
// @Accept       	json
// @Accept       	x-www-form-urlencoded
// @Produce      	json
// @Param			body			body	model.UserCreate	true	"body"
// @Success      	201  {object}   model.CreateResponse
// @Failure      	400,500  {object}  	pkg.Error
// @Router       	/auth/registration [post]
func (ctr *AuthController) Create(c echo.Context) error {
	return ctr.CreateUserControllerImpl.Create(c)
}
