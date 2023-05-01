package model

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	UserType string
	FullName string
	Sub      string
}

type UserCreate struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Username string `json:"username"`
	Fullname string `json:"fullname"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (u User) GetID() string {
	return u.Sub
}

func (u User) IsUserType(userType string) bool {
	return u.UserType == userType
}

func UserFromJWTContext(userCtx interface{}) User {
	user := User{}
	if token, ok := userCtx.(*jwt.Token); ok {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			if fullName, ok := claims["full_name"]; ok {
				user.FullName = fullName.(string)
			}
			if sub, ok := claims["sub"]; ok {
				user.Sub = sub.(string)
			}
			if userType, ok := claims["user_type"]; ok {
				user.UserType = userType.(string)
			}
		}
	}
	return user
}

func BindUserLogin(binder Binder) (interface{}, error) {
	creds := new(UserLogin)
	if err := binder.Bind(creds); err != nil {
		return nil, err
	}
	return creds, nil
}

func ValidateUserLogin(validator Validator, data interface{}) error {
	creds := data.(*UserLogin)
	return validator.Validate(creds)
}

func BindUserCreate(binder Binder) (data interface{}, err error) {
	position := new(UserCreate)
	if err = binder.Bind(position); err != nil {
		return nil, err
	}
	return position, nil
}

func ValidateUserCreate(validator Validator, data interface{}) error {
	position := data.(*UserCreate)
	return validator.Validate(position)
}
