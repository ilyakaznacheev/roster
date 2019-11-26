package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/ilyakaznacheev/roster/internal/models"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/auth"
)

type AuthHandler struct {
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{}
}

func (h *AuthHandler) HandleLogin(params auth.PostLoginParams) middleware.Responder {

	respData := models.AuthToken{}

	return auth.NewPostLoginOK().WithPayload(&respData)
}
