package auth

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/fatram/golang-dansmultipro-test/config"
	"github.com/fatram/golang-dansmultipro-test/domain/entity"
	"github.com/fatram/golang-dansmultipro-test/domain/model"
	"github.com/fatram/golang-dansmultipro-test/domain/repository"
	imysql "github.com/fatram/golang-dansmultipro-test/domain/repository/mysql"
	"github.com/fatram/golang-dansmultipro-test/internal/connector"
	"github.com/fatram/golang-dansmultipro-test/internal/pkg"
	"github.com/fatram/golang-dansmultipro-test/internal/pkg/beautify"
	"github.com/fatram/golang-dansmultipro-test/pkg/jwt"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	e    = pkg.LoadEcho()
	user = entity.User{
		ID:       "111",
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

func TestLogin(t *testing.T) {
	jsonData := `{
		"email": "satu@satu.satu",
		"password": "satu"
	}`
	req := httptest.NewRequest(http.MethodPost, "/user/auth/login", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewAuthController(e)
	if assert.NoError(t, h.Login(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestUserRegistration(t *testing.T) {
	jsonData := `{
		"fullname": "Tiga",
		"password": "tiga",
		"email": "tiga@tiga.tiga",
		"username": "tiga"
	}`
	req := httptest.NewRequest(http.MethodPost, "/auth/registration", strings.NewReader(jsonData))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	h := NewAuthController(e)
	idModel := model.CreateResponse{}
	if assert.NoError(t, h.Create(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		json.Unmarshal(rec.Body.Bytes(), &idModel)
		id := idModel.ID
		assert.NotEmpty(t, id)
		saved, _ := userRepo.GetByEmailPassword(ctx, "tiga@tiga.tiga", "tiga")
		log.Printf("Saved entity: %s", beautify.JSONString(saved))
		assert.NotEmpty(t, saved.ID)
		assert.Equal(t, id, saved.ID)
		assert.Equal(t, "Tiga", saved.Fullname)
		assert.Equal(t, "tiga", saved.Password)
		assert.Equal(t, "tiga@tiga.tiga", saved.Email)
		assert.Equal(t, "tiga", saved.Username)
		log.Printf("Result: %s", rec.Body.String())
	}
}
