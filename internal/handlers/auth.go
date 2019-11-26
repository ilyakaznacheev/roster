package handlers

import (
	"github.com/go-openapi/runtime/middleware"

	"github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/auth"
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
