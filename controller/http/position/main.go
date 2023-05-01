package position

import (
	"github.com/fatram/golang-dansmultipro-test/config"
	"github.com/fatram/golang-dansmultipro-test/internal/middleware"
	"github.com/labstack/echo/v4"
)

type PositionController struct {
	echo *echo.Echo
	*GetAllPositionControllerImpl
	*GetPositionControllerImpl
}

func NewPositionController(echo *echo.Echo) *PositionController {
	return &PositionController{
		echo:                         echo,
		GetAllPositionControllerImpl: newGetAllPositionController(echo.Logger),
		GetPositionControllerImpl:    newGetPositionController(echo.Logger),
	}
}

func (ctr PositionController) Start() {
	loginMiddleware := middleware.EchoJWTRSA(
		config.Configuration().GetPublicKey(),
	)

	r := ctr.echo.Group("/")
	r.GET("positions", ctr.GetAll, loginMiddleware)
	r.GET("positions/:id", ctr.Get, loginMiddleware)
}
