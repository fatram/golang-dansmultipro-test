package middleware

import (
	"log"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func EchoJWTRSA(key []byte) echo.MiddlewareFunc {
	signKey, err := jwt.ParseRSAPublicKeyFromPEM(key)
	if err != nil {
		log.Fatalf("Failed to parse RSA public key2: %v", err)
	}
	return middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    signKey,
		SigningMethod: "RS256",
		AuthScheme:    "Bearer",
	})
}
