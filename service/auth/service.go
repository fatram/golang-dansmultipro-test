package auth

import (
	"context"
	"net/http"

	"github.com/fatram/golang-dansmultipro-test/config"
	"github.com/fatram/golang-dansmultipro-test/domain/entity"
	"github.com/fatram/golang-dansmultipro-test/domain/model"
	"github.com/fatram/golang-dansmultipro-test/domain/repository"
	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
	"github.com/fatram/golang-dansmultipro-test/pkg/jwt"
	"github.com/labstack/echo/v4"
)

type AuthService struct {
	logger         genlog.Logger
	userRepository repository.UserRepository
	jwtHandler     *jwt.JWTHandler
}

func (s *AuthService) Create(ctx context.Context, data interface{}) (string, error) {
	userModel, ok := data.(*model.UserCreate)
	if !ok {
		s.logger.Errorf("data tidak sesuai")
		return "", echo.NewHTTPError(http.StatusBadRequest, "data tidak sesuai")
	}

	user := entity.User{
		Email:    userModel.Email,
		Password: userModel.Password,
		Username: userModel.Username,
		Fullname: userModel.Fullname,
	}
	user.SetID()
	return s.userRepository.Create(ctx, user)
}

func (s *AuthService) Login(ctx context.Context, data interface{}) (string, error) {
	creds, ok := data.(*model.UserLogin)
	if !ok {
		s.logger.Errorf("data tidak sesuai")
		return "", echo.NewHTTPError(http.StatusBadRequest, "data tidak sesuai")
	}
	if creds.Password == "" || creds.Email == "" {
		s.logger.Errorf("data kosong")
		return "", echo.NewHTTPError(http.StatusBadRequest, "data kosong")
	}

	user, err := s.userRepository.GetByEmailPassword(ctx, creds.Email, creds.Password)
	if err != nil {
		s.logger.Errorf(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if user == nil {
		s.logger.Errorf("user tidak ditemukan")
		return "", echo.NewHTTPError(http.StatusUnauthorized, "user tidak ditemukan")
	}

	token, err := s.jwtHandler.GenerateForAuthSession(*user, config.Configuration().AccessTokenTTL)
	if err != nil {
		s.logger.Errorf(err.Error())
		return "", echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return token, nil
}
