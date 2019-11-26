package handlers

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/ilyakaznacheev/roster/internal/models"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/roster"
)

type DatabaseService interface {
}

type RosterHandler struct {
	DB DatabaseService
}

// NewRosterHandler creates a new web API handler
func NewRosterHandler(db DatabaseService) *RosterHandler {
	return &RosterHandler{
		DB: db,
	}
}

func (h *RosterHandler) GetRosterAll(params roster.GetRostersParams) middleware.Responder {

	respData := models.AllRosters{}

	return roster.NewGetRostersOK().WithPayload(respData)
}

func (h *RosterHandler) GetRosterOne(params roster.GetRostersIDParams) middleware.Responder {

	respData := &models.Roster{}

	return roster.NewGetRostersIDOK().WithPayload(respData)
}
