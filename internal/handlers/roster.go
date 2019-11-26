package handlers

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	apiModels "github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/player"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/roster"
	dbModels "github.com/ilyakaznacheev/roster/internal/database/models"
)

type DatabaseService interface {
	GetAllRosters() ([]dbModels.Roster, error)
	GetRoster(id int) (*dbModels.Roster, error)
	UpdateRoster(dbModels.Roster) error
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

	h.DB.St

	rosters, err := h.DB.GetAllRosters()
	if err != nil {

		return roster.NewGetRostersInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	respData := make(apiModels.AllRosters, 0, len(rosters))

	for idx, r := range rosters {
		respRoster := &apiModels.Roster{
			ID: &rosters[idx].ID,
			Players: &apiModels.RosterPlayers{
				Active:  make([]*apiModels.Player, 0, len(r.Players.Active)),
				Benched: make([]*apiModels.Player, 0, len(r.Players.Benched)),
			},
		}
		for idx := range r.Players.Active {
			respRoster.Players.Active = append(respRoster.Players.Active, &apiModels.Player{
				ID:        &r.Players.Active[idx].ID,
				FirstName: &r.Players.Active[idx].FirstName,
				LastName:  &r.Players.Active[idx].LastName,
				Alias:     &r.Players.Active[idx].Alias,
			})
		}

		for idx := range r.Players.Benched {
			respRoster.Players.Benched = append(respRoster.Players.Benched, &apiModels.Player{
				ID:        &r.Players.Benched[idx].ID,
				FirstName: &r.Players.Benched[idx].FirstName,
				LastName:  &r.Players.Benched[idx].LastName,
				Alias:     &r.Players.Benched[idx].Alias,
			})
		}

		respData = append(respData, respRoster)
	}

	return roster.NewGetRostersOK().WithPayload(respData)
}

func (h *RosterHandler) GetRosterOne(params roster.GetRostersIDParams) middleware.Responder {

	respData := &apiModels.Roster{}

	return roster.NewGetRostersIDOK().WithPayload(respData)
}

func (h *RosterHandler) GetRosterActive(params roster.GetRostersIDActiveParams) middleware.Responder {

	respData := &apiModels.Roster{}

	return roster.NewGetRostersIDActiveOK().WithPayload(respData)
}

func (h *RosterHandler) GetRosterBenched(params roster.GetRostersIDBenchedParams) middleware.Responder {

	respData := &apiModels.Roster{}

	return roster.NewGetRostersIDBenchedOK().WithPayload(respData)
}

func (h *RosterHandler) AddPayer(params player.PostRostersIDAddPlayerParams, _ interface{}) middleware.Responder {

	respData := &apiModels.Player{}

	return player.NewPostRostersIDAddPlayerCreated().WithPayload(respData)
}

func (h *RosterHandler) RearrangeRoster(params player.PostRostersIDRearrangeParams, _ interface{}) middleware.Responder {

	respData := &apiModels.Roster{}

	return player.NewPostRostersIDRearrangeOK().WithPayload(respData)
}

func errorResp(status int64, message string) *apiModels.Error {
	return &apiModels.Error{
		Code:    &status,
		Message: &message,
	}
}
