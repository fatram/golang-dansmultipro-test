package mysql

import (
	"context"
	"database/sql"
	"sync"

	"github.com/fatram/golang-dansmultipro-test/domain/entity"
	"github.com/fatram/golang-dansmultipro-test/domain/repository"
	"github.com/fatram/golang-dansmultipro-test/internal/connector"
	"github.com/fatram/golang-dansmultipro-test/pkg/genlog"
	"github.com/pkg/errors"
)

type userRepositoryImpl struct {
	logger genlog.Logger
	db     *sql.DB
}

var (
	userRepository     *userRepositoryImpl
	userRepositoryOnce sync.Once
)

func LoadUserRepository(logger genlog.Logger) repository.UserRepository {
	userRepositoryOnce.Do(func() {
		userRepository = &userRepositoryImpl{
			logger: logger,
			db:     connector.LoadMysqlDatabase(),
		}
	})
	return userRepository
}

func (r *userRepositoryImpl) GetByEmailPassword(ctx context.Context, email, password string) (data *entity.User, err error) {
	query := `
		SELECT
			id, email, password, username, fullname
		FROM
			user
		WHERE email=? AND password=?
	`
	row := r.db.QueryRowContext(ctx, query, email, password)
	if err != nil {
		return nil, errors.Wrap(err, "error in userRepositoryImpl, GetByEmailPassword")
	}
	data = &entity.User{}
	err = row.Scan(&data.ID, &data.Email, &data.Password, &data.Username, &data.Fullname)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error in userRepositoryImpl, GetByEmailPassword")
	}
	return data, nil
}

func (r userRepositoryImpl) Create(ctx context.Context, data entity.User) (id string, err error) {
	query := `
		INSERT INTO user (id, email, password, username, fullname)
		VALUES (?, ?, ?, ?, ?)
	`
	_, err = r.db.ExecContext(ctx, query, data.ID, data.Email, data.Password, data.Username, data.Fullname)
	if err != nil {
		return "", errors.Wrap(err, "error in userRepositoryImpl, Create")
	}
	return data.ID, nil
}
