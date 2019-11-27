package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-openapi/runtime/middleware"

	"github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/auth"
)

// AuthHandler is an authentication API request handler
type AuthHandler struct {
	authKey []byte
}

// NewAuthHandler creates a new authentication API request handler
func NewAuthHandler(authKey string) *AuthHandler {
	return &AuthHandler{
		authKey: []byte(authKey),
	}
}

// HandleLogin checks user credentials and issues a JWT
func (h *AuthHandler) HandleLogin(params auth.PostLoginParams) middleware.Responder {

	if !h.isValidUser(*params.Body.Login, *params.Body.Password) {
		return auth.NewPostLoginForbidden().WithPayload(
			errorResp(http.StatusForbidden, "Forbidden"))
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

	respData := models.AuthToken{
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

// isValidUser checks user validity/registration status
func (h *AuthHandler) isValidUser(login, password string) bool {
	// Just an example. Here should be a request to login database or service
	return true
}
