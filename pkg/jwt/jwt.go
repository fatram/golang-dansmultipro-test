package jwt

import (
	"crypto/rsa"
	"log"
	"net/http"
	"time"

	"github.com/fatram/golang-dansmultipro-test/domain/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type JWTHandler struct {
	pubKey  *rsa.PublicKey
	privKey *rsa.PrivateKey
}

func NewJWTHandler(publicKeyFile []byte, privateKeyFile []byte) *JWTHandler {
	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicKeyFile)
	if err != nil {
		log.Fatalln("jwt public key parser failed: please check in \"log_data\" directory for further information")
	}
	privKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKeyFile)
	if err != nil {
		log.Fatalln("jwt private key parser failed: please check in \"log_data\" directory for further information")
	}

	return &JWTHandler{pubKey, privKey}
}

func (j JWTHandler) GenerateForAuthSession(payload entity.User, lifespan time.Duration) (token string, err error) {
	now := time.Now()
	expirationTime := now.Add(lifespan)

	// create jwt claims
	claims := jwt.MapClaims{
		"exp":       expirationTime.Unix(),
		"iat":       now.Unix(),
		"nbf":       now.Unix(),
		"sub":       payload.ID,
		"full_name": payload.Fullname,
	}
	jwttoken := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	return jwttoken.SignedString(j.privKey)
}

func (j JWTHandler) GenerateWithAction(payload string, lifespan time.Duration, idSchool string, action string) (token string, err error) {
	now := time.Now()
	expirationTime := now.Add(lifespan)

	// create jwt claims
	claims := jwt.MapClaims{
		"exp":       expirationTime.Unix(),
		"iat":       now.Unix(),
		"nbf":       now.Unix(),
		"sub":       payload,
		"school_id": idSchool,
		"action":    action,
	}
	jwttoken := jwt.NewWithClaims(jwt.GetSigningMethod("RS256"), claims)

	if token, err = jwttoken.SignedString(j.privKey); err != nil {
		err = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return
}

func (j JWTHandler) ValidateActionWithClaim(token string) (isValid bool, claim jwt.MapClaims, err error) {
	jwttoken, err := jwt.Parse(token, func(jwttoken *jwt.Token) (interface{}, error) {
		if _, ok := jwttoken.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, echo.ErrUnauthorized
		}
		return j.pubKey, nil
	})

	if err != nil {
		_ = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		err = nil
	}

	if !jwttoken.Valid {
		err = echo.NewHTTPError(http.StatusUnauthorized, "invalid token")
		return
	}

	claim, ok := jwttoken.Claims.(jwt.MapClaims)
	if !ok {
		err = echo.NewHTTPError(http.StatusUnauthorized, "cannot assert jwt payload to MapClaims")
		return
	}
	isValid = true
	return
}
