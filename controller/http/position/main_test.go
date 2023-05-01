package position

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/fatram/golang-dansmultipro-test/config"
	"github.com/fatram/golang-dansmultipro-test/domain/entity"
	"github.com/fatram/golang-dansmultipro-test/domain/model"
	"github.com/fatram/golang-dansmultipro-test/domain/repository"
	imysql "github.com/fatram/golang-dansmultipro-test/domain/repository/mysql"
	"github.com/fatram/golang-dansmultipro-test/internal/connector"
	"github.com/fatram/golang-dansmultipro-test/internal/pkg"
	"github.com/fatram/golang-dansmultipro-test/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

var (
	e    = pkg.LoadEcho()
	user = entity.User{
		ID:       "1",
		Username: "satu",
		Fullname: "Satu",
		Email:    "satu@satu.satu",
		Password: "satu",
	}

	userRepo   repository.UserRepository
	jwtHandler *jwt.JWTHandler
	db         *sql.DB
	ctx        = context.Background()
	userID     = ""
	positionID = "ecbe528e-ae60-45ad-9706-76819ae07c85"
	userToken  = ""
)

func setUp() {
	config.ReadConfig("../../../.test.env")
	db = connector.LoadMysqlDatabase()
	e.Validator = &pkg.CustomValidator{Validator: validator.New()}
	userRepo = imysql.LoadUserRepository(e.Logger)
	jwtHandler = jwt.NewJWTHandler(config.Configuration().GetPublicKey(), config.Configuration().GetPrivateKey())
	id, err := userRepo.Create(ctx, user)
	if err != nil {
		log.Fatal(err)
	}
	userID = id
	userToken, err = jwtHandler.GenerateForAuthSession(user, 1*time.Hour)
	if err != nil {
		log.Fatal(err)
	}
}

func tearDown() {
	_, err := db.Exec("DELETE FROM user WHERE email='satu@satu.satu'")
	if err != nil {
		log.Println("failed deleting user data")
	}
	db.Close()
}

func TestMain(m *testing.M) {
	setUp()
	code := m.Run()
	tearDown()
	os.Exit(code)
}

func TestGetAllPosition(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/positions", nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userToken))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewPositionController(e)

	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := []model.Position{}
		json.Unmarshal(rec.Body.Bytes(), &body)
		data := reflect.ValueOf(body)
		assert.Greater(t, data.Len(), 0)
	}
}

func TestGetPosition(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/positions/"+positionID, nil)
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", userToken))
	rec := httptest.NewRecorder()
	h := NewPositionController(e)
	c := e.NewContext(req, rec)

	if assert.NoError(t, h.GetAll(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		body := model.Position{}
		json.Unmarshal(rec.Body.Bytes(), &body)
		data := reflect.ValueOf(body)
		assert.NotEmpty(t, data)
	}
}
