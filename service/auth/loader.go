package auth

import (
	"sync"

	"github.com/fatram/golang-dansmultipro-test/config"
	"github.com/fatram/golang-dansmultipro-test/domain/repository/mysql"
	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
	"github.com/fatram/golang-dansmultipro-test/pkg/jwt"
)

var (
	authService     *AuthService
	onceAuthService sync.Once
)

func LoadAuthService(logger genlog.Logger) *AuthService {
	onceAuthService.Do(func() {
		authService = &AuthService{
			logger:         logger,
			userRepository: mysql.LoadUserRepository(logger),
			jwtHandler:     jwt.NewJWTHandler(config.Configuration().GetPublicKey(), config.Configuration().GetPrivateKey()),
		}
	})
	return authService
}
