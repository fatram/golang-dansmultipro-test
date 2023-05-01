package repository

import (
	"context"

	"github.com/fatram/golang-dansmultipro-test/domain/entity"
)

type UserRepository interface {
	GetByEmailPassword(ctx context.Context, email, password string) (data *entity.User, err error)
	Create(ctx context.Context, data entity.User) (id string, err error)
}
