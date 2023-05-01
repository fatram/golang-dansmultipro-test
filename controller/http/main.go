package http

import (
	"fmt"

	"github.com/fatram/golang-dansmultipro-test/controller/http/auth"
	"github.com/fatram/golang-dansmultipro-test/controller/http/position"
	_ "github.com/fatram/golang-dansmultipro-test/docs"
	"github.com/fatram/golang-dansmultipro-test/internal/pkg"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4" // we use echo version 4 here
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type HttpController interface {
	Start(host string, port int)
}

type httpCtr struct {
	echo *echo.Echo
}

func NewHttpController() HttpController {
	e := echo.New()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Logger.SetLevel(log.DEBUG)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
	e.Logger.SetHeader(pkg.DefaultLogHeader)

	return &httpCtr{e}
}

func (ctr *httpCtr) Start(host string, port int) {
	auth.NewAuthController(ctr.echo).Start()
	position.NewPositionController(ctr.echo).Start()
	ctr.echo.Logger.Fatal(ctr.echo.Start(fmt.Sprintf("%s:%d", host, port)))
}
