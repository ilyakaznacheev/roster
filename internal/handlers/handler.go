package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ilyakaznacheev/roster/internal/models"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/roster"
)

type RosterHandler struct {
}

func (h *RosterHandler) GetRoster(params roster.GetRosterParams) middleware.Responder {

	respData := models.AllRosters{}

	return roster.NewGetRosterOK().WithPayload(respData)
}
