package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"
	"golang.org/x/crypto/bcrypt"

	apiModels "github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/auth"
	"github.com/ilyakaznacheev/roster/internal/database"
	dbModels "github.com/ilyakaznacheev/roster/internal/database/models"
)

// DatabaseService is a abstract database layer interface
type DatabaseAuthService interface {
	AddUser(c dbModels.Credentials) error
	GetUser(login string) (*dbModels.Credentials, error)
}

// AuthHandler is an authentication API request handler
type AuthHandler struct {
	DB      DatabaseAuthService
	authKey []byte
}

// NewAuthHandler creates a new authentication API request handler
func NewAuthHandler(authKey string, db DatabaseAuthService) *AuthHandler {
	return &AuthHandler{
		authKey: []byte(authKey),
		DB:      db,
	}
}

func (h *AuthHandler) HandleRegistration(params auth.PostRegisterParams) middleware.Responder {
	hPass, err := hashPassword(*params.Body.Password)
	if err != nil {
		return auth.NewPostRegisterInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}
	c := dbModels.Credentials{
		Login:    *params.Body.Login,
		PassHash: hPass,
	}
	err = h.DB.AddUser(c)
	if errors.As(err, &database.ErrExists) {
		return auth.NewPostRegisterConflict().WithPayload(
			errorResp(http.StatusConflict, err.Error()))
	} else if err != nil {
		return auth.NewPostRegisterInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	return auth.NewPostRegisterCreated()

}

// HandleLogin checks user credentials and issues a JWT
func (h *AuthHandler) HandleLogin(params auth.PostLoginParams) middleware.Responder {
	c, err := h.DB.GetUser(*params.Body.Login)
	if errors.As(err, &database.ErrNotFound) {
		return auth.NewPostLoginForbidden().WithPayload(
			errorResp(http.StatusForbidden, "forbidden"))
	} else if err != nil {
		return auth.NewPostLoginInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	if !verifyPassword(*params.Body.Password, c.PassHash) {
		return auth.NewPostLoginForbidden().WithPayload(
			errorResp(http.StatusForbidden, "forbidden"))
	}

	claims := make(jwt.MapClaims)

	claims["iss"] = ""
	claims["sub"] = *params.Body.Login
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(h.authKey)
	if err != nil {
		return auth.NewPostLoginInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	log.Printf("issued token: %s", tokenString)

	respData := apiModels.AuthToken{
		Token: &tokenString,
	}

	return auth.NewPostLoginOK().WithPayload(&respData)
}

// Authenticate validates a JWT
func (h *AuthHandler) Authenticate(tokenString string) (interface{}, error) {
	ts := strings.Split(tokenString, ` `)[1]
	token, err := jwt.Parse(ts, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return h.authKey, nil
	})

	return token.Valid, err
}

func hashPassword(p string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func verifyPassword(pass, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	return err == nil
}
