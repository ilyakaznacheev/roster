package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/go-openapi/runtime/middleware"

	apiModels "github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/player"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/roster"
	"github.com/ilyakaznacheev/roster/internal/database"
	dbModels "github.com/ilyakaznacheev/roster/internal/database/models"
)

const activePlayersNumber = 5

// DatabaseRosterService is a abstract database layer interface
type DatabaseRosterService interface {
	GetAllRosters() ([]dbModels.Roster, error)
	GetRoster(id int64) (*dbModels.Roster, error)
	PushPlayer(id int64, p dbModels.Player) error
	UpdateRoster(dbModels.Roster) error
}

type IDGeneratorFunc func() int64

func (g IDGeneratorFunc) Generate() int64 {
	return g()
}

// RosterHandler is a handler for roster API requests
type RosterHandler struct {
	DB    DatabaseRosterService
	IDGen IDGeneratorFunc
}

// NewRosterHandler creates a new web API handler
func NewRosterHandler(db DatabaseRosterService) *RosterHandler {
	return &RosterHandler{
		DB:    db,
		IDGen: dbModels.GenerateID,
	}
}

// GetRosterAll returns a list of all rosters
func (h *RosterHandler) GetRosterAll(params roster.GetRostersParams) middleware.Responder {
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
				Role:      string(r.Players.Active[idx].Role),
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

// GetRosterOne returns a certain roster
func (h *RosterHandler) GetRosterOne(params roster.GetRostersIDParams) middleware.Responder {
	r, err := h.DB.GetRoster(params.ID)
	if errors.As(err, &database.ErrNotFound) {
		return roster.NewGetRostersIDNotFound().WithPayload(
			errorResp(http.StatusNotFound, err.Error()))
	} else if err != nil {
		return roster.NewGetRostersIDInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	respData := &apiModels.Roster{
		ID: &r.ID,
		Players: &apiModels.RosterPlayers{
			Active:  make([]*apiModels.Player, 0, len(r.Players.Active)),
			Benched: make([]*apiModels.Player, 0, len(r.Players.Benched)),
		},
	}
	for idx := range r.Players.Active {
		respData.Players.Active = append(respData.Players.Active, &apiModels.Player{
			ID:        &r.Players.Active[idx].ID,
			FirstName: &r.Players.Active[idx].FirstName,
			LastName:  &r.Players.Active[idx].LastName,
			Alias:     &r.Players.Active[idx].Alias,
			Role:      string(r.Players.Active[idx].Role),
		})
	}

	for idx := range r.Players.Benched {
		respData.Players.Benched = append(respData.Players.Benched, &apiModels.Player{
			ID:        &r.Players.Benched[idx].ID,
			FirstName: &r.Players.Benched[idx].FirstName,
			LastName:  &r.Players.Benched[idx].LastName,
			Alias:     &r.Players.Benched[idx].Alias,
			Role:      string(r.Players.Benched[idx].Role),
		})
	}

	return roster.NewGetRostersIDOK().WithPayload(respData)
}

// GetRosterActive returns a certain roster with active players
func (h *RosterHandler) GetRosterActive(params roster.GetRostersIDActiveParams) middleware.Responder {
	r, err := h.DB.GetRoster(params.ID)
	if errors.As(err, &database.ErrNotFound) {
		return roster.NewGetRostersIDActiveNotFound().WithPayload(
			errorResp(http.StatusNotFound, err.Error()))
	} else if err != nil {
		return roster.NewGetRostersIDActiveInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	respData := &apiModels.Roster{
		ID: &r.ID,
		Players: &apiModels.RosterPlayers{
			Active: make([]*apiModels.Player, 0, len(r.Players.Active)),
		},
	}
	for idx := range r.Players.Active {
		respData.Players.Active = append(respData.Players.Active, &apiModels.Player{
			ID:        &r.Players.Active[idx].ID,
			FirstName: &r.Players.Active[idx].FirstName,
			LastName:  &r.Players.Active[idx].LastName,
			Alias:     &r.Players.Active[idx].Alias,
			Role:      string(r.Players.Active[idx].Role),
		})
	}

	return roster.NewGetRostersIDActiveOK().WithPayload(respData)
}

// GetRosterBenched returns a certain roster with benched players
func (h *RosterHandler) GetRosterBenched(params roster.GetRostersIDBenchedParams) middleware.Responder {
	r, err := h.DB.GetRoster(params.ID)
	if errors.As(err, &database.ErrNotFound) {
		return roster.NewGetRostersIDBenchedNotFound().WithPayload(
			errorResp(http.StatusNotFound, err.Error()))
	} else if err != nil {
		return roster.NewGetRostersIDBenchedInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	respData := &apiModels.Roster{
		ID: &r.ID,
		Players: &apiModels.RosterPlayers{
			Benched: make([]*apiModels.Player, 0, len(r.Players.Benched)),
		},
	}
	for idx := range r.Players.Benched {
		respData.Players.Benched = append(respData.Players.Benched, &apiModels.Player{
			ID:        &r.Players.Benched[idx].ID,
			FirstName: &r.Players.Benched[idx].FirstName,
			LastName:  &r.Players.Benched[idx].LastName,
			Alias:     &r.Players.Benched[idx].Alias,
			Role:      string(r.Players.Benched[idx].Role),
		})
	}

	return roster.NewGetRostersIDBenchedOK().WithPayload(respData)
}

// AddPayer adds a player to the roster
func (h *RosterHandler) AddPayer(params player.PostRostersIDAddPlayerParams, _ interface{}) middleware.Responder {
	p := dbModels.Player{
		ID:        h.IDGen.Generate(),
		FirstName: *params.Body.FirstName,
		LastName:  *params.Body.LastName,
		Alias:     *params.Body.Alias,
		Role:      dbModels.PlayerRole(*params.Body.Role),
	}

	err := h.DB.PushPlayer(params.ID, p)
	if errors.As(err, &database.ErrNotFound) {
		return player.NewPostRostersIDRearrangeNotFound().WithPayload(
			errorResp(http.StatusNotFound, err.Error()))
	} else if err != nil {
		return player.NewPostRostersIDRearrangeInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	respData := &apiModels.Player{
		ID:        &p.ID,
		FirstName: &p.FirstName,
		LastName:  &p.LastName,
		Alias:     &p.Alias,
		Role:      string(p.Role),
	}

	return player.NewPostRostersIDAddPlayerCreated().WithPayload(respData)
}

// RearrangeRoster rearranges players in a certain roster
func (h *RosterHandler) RearrangeRoster(params player.PostRostersIDRearrangeParams, _ interface{}) middleware.Responder {
	// prepare request keys
	activeKeys := make(map[int64]struct{}, len(params.Body.ToActive))
	for _, id := range params.Body.ToActive {
		activeKeys[id] = struct{}{}
	}

	benchedKeys := make(map[int64]struct{}, len(params.Body.ToBenched))
	for _, id := range params.Body.ToBenched {
		benchedKeys[id] = struct{}{}
	}

	// validations

	// check the number of keys. It can be only 5 active players,
	// so the requester has to add the same number of players as remove
	if len(activeKeys) != len(benchedKeys) {
		return player.NewPostRostersIDRearrangeBadRequest().WithPayload(
			errorResp(http.StatusBadRequest, "wrong number of players"))
	}

	// check if key sets don't cross
	for id := range activeKeys {
		if _, ok := benchedKeys[id]; ok {
			errorResp(http.StatusBadRequest, fmt.Sprintf("players id %d is duplicated", id))
		}
	}
	for id := range benchedKeys {
		if _, ok := activeKeys[id]; ok {
			errorResp(http.StatusBadRequest, fmt.Sprintf("players id %d is duplicated", id))
		}
	}

	// read data from the database
	r, err := h.DB.GetRoster(params.ID)
	if errors.As(err, &database.ErrNotFound) {
		return player.NewPostRostersIDRearrangeNotFound().WithPayload(
			errorResp(http.StatusNotFound, err.Error()))
	} else if err != nil {
		return player.NewPostRostersIDRearrangeInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	if len(r.Players.Active) != activePlayersNumber {
		return player.NewPostRostersIDRearrangeBadRequest().WithPayload(
			errorResp(http.StatusBadRequest, fmt.Sprintf("bad roster: the initial amount of active players is %d", len(r.Players.Active))))
	}

	// process the change
	newActive := make([]dbModels.Player, 0, len(r.Players.Active))
	newBenched := make([]dbModels.Player, 0, len(r.Players.Benched))

	for _, p := range r.Players.Active {
		if _, ok := benchedKeys[p.ID]; ok {
			newBenched = append(newBenched, p)
			delete(benchedKeys, p.ID)
		} else {
			newActive = append(newActive, p)
		}
	}

	for _, p := range r.Players.Benched {
		if _, ok := activeKeys[p.ID]; ok {
			newActive = append(newActive, p)
			delete(activeKeys, p.ID)
		} else {
			newBenched = append(newBenched, p)
		}
	}

	// check if all players were found
	if len(activeKeys) != 0 || len(benchedKeys) != 0 {
		return player.NewPostRostersIDRearrangeBadRequest().WithPayload(
			errorResp(http.StatusBadRequest, "wrong player ids"))
	}

	r.Players.Active = newActive
	r.Players.Benched = newBenched

	err = h.DB.UpdateRoster(*r)
	if errors.As(err, &database.ErrNotFound) {
		return player.NewPostRostersIDRearrangeInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, "data was changed in concurrent process\nplease repeat the request"))
	} else if err != nil {
		return player.NewPostRostersIDRearrangeInternalServerError().WithPayload(
			errorResp(http.StatusInternalServerError, err.Error()))
	}

	// prepare response
	respData := &apiModels.Roster{
		ID: &r.ID,
		Players: &apiModels.RosterPlayers{
			Active:  make([]*apiModels.Player, 0, len(r.Players.Active)),
			Benched: make([]*apiModels.Player, 0, len(r.Players.Benched)),
		},
	}
	for idx := range r.Players.Active {
		respData.Players.Active = append(respData.Players.Active, &apiModels.Player{
			ID:        &r.Players.Active[idx].ID,
			FirstName: &r.Players.Active[idx].FirstName,
			LastName:  &r.Players.Active[idx].LastName,
			Alias:     &r.Players.Active[idx].Alias,
			Role:      string(r.Players.Active[idx].Role),
		})
	}

	for idx := range r.Players.Benched {
		respData.Players.Benched = append(respData.Players.Benched, &apiModels.Player{
			ID:        &r.Players.Benched[idx].ID,
			FirstName: &r.Players.Benched[idx].FirstName,
			LastName:  &r.Players.Benched[idx].LastName,
			Alias:     &r.Players.Benched[idx].Alias,
			Role:      string(r.Players.Benched[idx].Role),
		})
	}

	return player.NewPostRostersIDRearrangeOK().WithPayload(respData)
}

// errorResp response error helper
func errorResp(status int64, message string) *apiModels.Error {
	log.Printf("error %d: %s", status, message)
	return &apiModels.Error{
		Code:    &status,
		Message: &message,
	}
}
