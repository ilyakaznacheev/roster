package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/go-openapi/runtime/middleware"
	"github.com/stretchr/testify/assert"

	apiModels "github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/player"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/roster"
	"github.com/ilyakaznacheev/roster/internal/database"
	dbModels "github.com/ilyakaznacheev/roster/internal/database/models"
	"github.com/ilyakaznacheev/roster/internal/handlers/mocks"
)

var (
	ErrTestGeneric  = errors.New("test")
	ErrTestNotFound = &database.NotFoundError{Text: "not found"}
)

func TestRosterHandler_GetRosterAll(t *testing.T) {
	p1 := dbModels.Player{
		ID:        12345,
		FirstName: "test1",
		LastName:  "test11",
		Alias:     "ttt111",
	}
	p2 := dbModels.Player{
		ID:        67890,
		FirstName: "test2",
		LastName:  "test22",
		Alias:     "ttt222",
	}

	type mockParams struct {
		res []dbModels.Roster
		err error
	}

	tests := []struct {
		name       string
		params     roster.GetRostersParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name:   "ok",
			params: roster.GetRostersParams{},
			mockParams: mockParams{
				res: []dbModels.Roster{
					{
						ID: 777,
						Players: dbModels.RosterPlayers{
							Active:  []dbModels.Player{p1},
							Benched: []dbModels.Player{p2},
						},
					},
				},
				err: nil,
			},
			wantStatus: 200,
			wantBody:   `[{"id":777,"players":{"active":[{"alias":"ttt111","first_name":"test1","id":12345,"last_name":"test11"}],"benched":[{"alias":"ttt222","first_name":"test2","id":67890,"last_name":"test22"}]}}]`,
		},

		{
			name:   "error",
			params: roster.GetRostersParams{},
			mockParams: mockParams{
				res: nil,
				err: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500, "message":"test"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseService)
			db.On("GetAllRosters").Return(tt.mockParams.res, tt.mockParams.err)

			h := &RosterHandler{
				DB: db,
			}

			validateResponse(t, h.GetRosterAll(tt.params), tt.wantBody, tt.wantStatus)
		})
	}
}

func TestRosterHandler_GetRosterOne(t *testing.T) {
	p1 := dbModels.Player{
		ID:        12345,
		FirstName: "test1",
		LastName:  "test11",
		Alias:     "ttt111",
	}
	p2 := dbModels.Player{
		ID:        67890,
		FirstName: "test2",
		LastName:  "test22",
		Alias:     "ttt222",
	}

	type mockParams struct {
		res *dbModels.Roster
		err error
	}

	tests := []struct {
		name       string
		params     roster.GetRostersIDParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name: "ok",
			params: roster.GetRostersIDParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: &dbModels.Roster{
					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err: nil,
			},
			wantStatus: 200,
			wantBody:   `{"id":777,"players":{"active":[{"alias":"ttt111","first_name":"test1","id":12345,"last_name":"test11"}],"benched":[{"alias":"ttt222","first_name":"test2","id":67890,"last_name":"test22"}]}}`,
		},

		{
			name: "error 404",
			params: roster.GetRostersIDParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: nil,
				err: ErrTestNotFound,
			},
			wantStatus: 404,
			wantBody:   `{"code":404, "message":"not found"}`,
		},

		{
			name: "error 500",
			params: roster.GetRostersIDParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: nil,
				err: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500, "message":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseService)
			db.On("GetRoster", int64(777)).Return(tt.mockParams.res, tt.mockParams.err)

			h := &RosterHandler{
				DB: db,
			}

			validateResponse(t, h.GetRosterOne(tt.params), tt.wantBody, tt.wantStatus)
		})
	}
}

func TestRosterHandler_GetRosterActive(t *testing.T) {
	p1 := dbModels.Player{
		ID:        12345,
		FirstName: "test1",
		LastName:  "test11",
		Alias:     "ttt111",
	}
	p2 := dbModels.Player{
		ID:        67890,
		FirstName: "test2",
		LastName:  "test22",
		Alias:     "ttt222",
	}

	type mockParams struct {
		res *dbModels.Roster
		err error
	}

	tests := []struct {
		name       string
		params     roster.GetRostersIDActiveParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name: "ok",
			params: roster.GetRostersIDActiveParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: &dbModels.Roster{
					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err: nil,
			},
			wantStatus: 200,
			wantBody:   `{"id":777,"players":{"active":[{"alias":"ttt111","first_name":"test1","id":12345,"last_name":"test11"}]}}`,
		},

		{
			name: "error 404",
			params: roster.GetRostersIDActiveParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: nil,
				err: ErrTestNotFound,
			},
			wantStatus: 404,
			wantBody:   `{"code":404, "message":"not found"}`,
		},

		{
			name: "error 500",
			params: roster.GetRostersIDActiveParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: nil,
				err: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500, "message":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseService)
			db.On("GetRoster", int64(777)).Return(tt.mockParams.res, tt.mockParams.err)

			h := &RosterHandler{
				DB: db,
			}

			validateResponse(t, h.GetRosterActive(tt.params), tt.wantBody, tt.wantStatus)
		})
	}
}

func TestRosterHandler_GetRosterBenched(t *testing.T) {
	p1 := dbModels.Player{
		ID:        12345,
		FirstName: "test1",
		LastName:  "test11",
		Alias:     "ttt111",
	}
	p2 := dbModels.Player{
		ID:        67890,
		FirstName: "test2",
		LastName:  "test22",
		Alias:     "ttt222",
	}

	type mockParams struct {
		res *dbModels.Roster
		err error
	}

	tests := []struct {
		name       string
		params     roster.GetRostersIDBenchedParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name: "ok",
			params: roster.GetRostersIDBenchedParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: &dbModels.Roster{
					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err: nil,
			},
			wantStatus: 200,
			wantBody:   `{"id":777,"players":{"benched":[{"alias":"ttt222","first_name":"test2","id":67890,"last_name":"test22"}]}}`,
		},

		{
			name: "error 404",
			params: roster.GetRostersIDBenchedParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: nil,
				err: ErrTestNotFound,
			},
			wantStatus: 404,
			wantBody:   `{"code":404, "message":"not found"}`,
		},

		{
			name: "error 500",
			params: roster.GetRostersIDBenchedParams{
				ID: 777,
			},
			mockParams: mockParams{
				res: nil,
				err: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500, "message":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseService)
			db.On("GetRoster", int64(777)).Return(tt.mockParams.res, tt.mockParams.err)

			h := &RosterHandler{
				DB: db,
			}

			validateResponse(t, h.GetRosterBenched(tt.params), tt.wantBody, tt.wantStatus)
		})
	}
}

func TestRosterHandler_AddPayer(t *testing.T) {
	pr := apiModels.PlayerRequest{
		FirstName: strToPtr("test1"),
		LastName:  strToPtr("test11"),
		Alias:     strToPtr("ttt111"),
	}
	pn := dbModels.Player{
		ID:        12345,
		FirstName: "test1",
		LastName:  "test11",
		Alias:     "ttt111",
	}

	type mockParams struct {
		id  int64
		req dbModels.Player
		err error
	}

	tests := []struct {
		name       string
		params     player.PostRostersIDAddPlayerParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name: "ok",
			params: player.PostRostersIDAddPlayerParams{
				ID:   12345,
				Body: &pr,
			},
			mockParams: mockParams{
				id:  12345,
				req: pn,
				err: nil,
			},
			wantStatus: 201,
			wantBody:   `{"alias":"ttt111","first_name":"test1","id":12345,"last_name":"test11"}`,
		},

		{
			name: "error 404",
			params: player.PostRostersIDAddPlayerParams{
				ID:   12345,
				Body: &pr,
			},
			mockParams: mockParams{
				id:  12345,
				req: pn,
				err: ErrTestNotFound,
			},
			wantStatus: 404,
			wantBody:   `{"code":404, "message":"not found"}`,
		},

		{
			name: "error 500",
			params: player.PostRostersIDAddPlayerParams{
				ID:   12345,
				Body: &pr,
			},
			mockParams: mockParams{
				id:  12345,
				req: pn,
				err: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500, "message":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseService)
			db.On("PushPlayer", tt.mockParams.id, tt.mockParams.req).Return(tt.mockParams.err)

			h := &RosterHandler{
				DB:    db,
				IDGen: getIDGenFunc(12345),
			}

			validateResponse(t, h.AddPayer(tt.params, nil), tt.wantBody, tt.wantStatus)
		})
	}
}

func TestRosterHandler_RearrangeRoster(t *testing.T) {
	p1 := dbModels.Player{
		ID:        12345,
		FirstName: "test1",
		LastName:  "test11",
		Alias:     "ttt111",
	}
	p2 := dbModels.Player{
		ID:        67890,
		FirstName: "test2",
		LastName:  "test22",
		Alias:     "ttt222",
	}

	type mockParams struct {
		// get roster
		id   int64
		res  *dbModels.Roster
		err1 error
		// update roster
		// func input
		req dbModels.Roster
		// func output
		err2 error
	}

	tests := []struct {
		name       string
		params     player.PostRostersIDRearrangeParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name: "ok",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{67890},
					ToBenched: []int64{12345},
				},
			},
			mockParams: mockParams{
				id: 777,
				res: &dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err1: nil,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: nil,
			},
			wantStatus: 200,
			wantBody:   `{"id":777,"players":{"active":[{"alias":"ttt222","first_name":"test2","id":67890,"last_name":"test22"}],"benched":[{"alias":"ttt111","first_name":"test1","id":12345,"last_name":"test11"}]}}`,
		},

		{
			name: "wrong number of request keys: active",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{67890, 54321},
					ToBenched: []int64{12345},
				},
			},
			mockParams: mockParams{
				id: 777,
				res: &dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err1: nil,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: nil,
			},
			wantStatus: 400,
			wantBody:   `{"code":400,"message":"wrong number of players"}`,
		},

		{
			name: "wrong number of request keys: benched",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{67890},
					ToBenched: []int64{12345, 54321},
				},
			},
			mockParams: mockParams{
				id: 777,
				res: &dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err1: nil,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: nil,
			},
			wantStatus: 400,
			wantBody:   `{"code":400,"message":"wrong number of players"}`,
		},

		{
			name: "wrong keys",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{54321},
					ToBenched: []int64{88888},
				},
			},
			mockParams: mockParams{
				id: 777,
				res: &dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err1: nil,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: nil,
			},
			wantStatus: 400,
			wantBody:   `{"code":400,"message":"wrong player ids"}`,
		},

		{
			name: "error not found",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{12345},
					ToBenched: []int64{67890},
				},
			},
			mockParams: mockParams{
				id:   777,
				res:  nil,
				err1: ErrTestNotFound,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: nil,
			},
			wantStatus: 404,
			wantBody:   `{"code":404,"message":"not found"}`,
		},

		{
			name: "read db error",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{67890},
					ToBenched: []int64{12345},
				},
			},
			mockParams: mockParams{
				id:   777,
				res:  nil,
				err1: ErrTestGeneric,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: nil,
			},
			wantStatus: 500,
			wantBody:   `{"code":500,"message":"test"}`,
		},

		{
			name: "concurrent change error",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{67890},
					ToBenched: []int64{12345},
				},
			},
			mockParams: mockParams{
				id: 777,
				res: &dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err1: nil,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: ErrTestNotFound,
			},
			wantStatus: 500,
			wantBody:   `{"code":500,"message":"data was changed in concurrent process\nplease repeat the request"}`,
		},

		{
			name: "update db error",
			params: player.PostRostersIDRearrangeParams{
				ID: 777,
				Body: &apiModels.RearrangeRequest{
					ToActive:  []int64{67890},
					ToBenched: []int64{12345},
				},
			},
			mockParams: mockParams{
				id: 777,
				res: &dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p1},
						Benched: []dbModels.Player{p2},
					},
				},
				err1: nil,
				req: dbModels.Roster{

					ID: 777,
					Players: dbModels.RosterPlayers{
						Active:  []dbModels.Player{p2},
						Benched: []dbModels.Player{p1},
					},
				},
				err2: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500,"message":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseService)
			db.On("GetRoster", tt.mockParams.id).Return(tt.mockParams.res, tt.mockParams.err1)
			db.On("UpdateRoster", tt.mockParams.req).Return(tt.mockParams.err2)

			h := &RosterHandler{
				DB: db,
			}

			validateResponse(t, h.RearrangeRoster(tt.params, nil), tt.wantBody, tt.wantStatus)
		})
	}
}

type mockProducer struct {
}

func (s *mockProducer) Produce(w io.Writer, p interface{}) error {
	return json.NewEncoder(w).Encode(p)
}

func validateResponse(t *testing.T, r middleware.Responder, body string, code int) {
	rc := httptest.NewRecorder()
	r.WriteResponse(rc, &mockProducer{})

	assert.JSONEqf(t, body, rc.Body.String(), "wrong response body: %s", rc.Body.String())
	assert.Equal(t, code, rc.Code, "wrong response status code")

}

func strToPtr(s string) *string {
	return &s
}

func getIDGenFunc(id int64) func() int64 {
	return func() int64 { return id }
}
