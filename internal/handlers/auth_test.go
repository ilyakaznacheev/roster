package handlers_test

import (
	"testing"

	"github.com/stretchr/testify/mock"

	apiModels "github.com/ilyakaznacheev/roster/internal/api/models"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/auth"
	dbModels "github.com/ilyakaznacheev/roster/internal/database/models"
	"github.com/ilyakaznacheev/roster/internal/handlers"
	"github.com/ilyakaznacheev/roster/internal/handlers/mocks"
)

func TestAuthHandler_HandleRegistration(t *testing.T) {
	type mockParams struct {
		err error
	}

	tests := []struct {
		name       string
		params     auth.PostRegisterParams
		mockParams mockParams
		wantStatus int
		wantBody   string
	}{
		{
			name: "ok",
			params: auth.PostRegisterParams{
				Body: &apiModels.AuthRequest{
					Login:    strToPtr("test-login"),
					Password: strToPtr("test-pass"),
				},
			},
			mockParams: mockParams{
				err: nil,
			},
			wantStatus: 201,
			wantBody:   "",
		},

		{
			name: "error exists",
			params: auth.PostRegisterParams{
				Body: &apiModels.AuthRequest{
					Login:    strToPtr("test-login"),
					Password: strToPtr("test-pass"),
				},
			},
			mockParams: mockParams{
				err: ErrTestExists,
			},
			wantStatus: 409,
			wantBody:   `{"code":409, "message":"already exists"}`,
		},

		{
			name: "error server",
			params: auth.PostRegisterParams{
				Body: &apiModels.AuthRequest{
					Login:    strToPtr("test-login"),
					Password: strToPtr("test-pass"),
				},
			},
			mockParams: mockParams{
				err: ErrTestGeneric,
			},
			wantStatus: 500,
			wantBody:   `{"code":500, "message":"test"}`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseAuthService)
			db.On("AddUser", mock.Anything).Return(tt.mockParams.err)

			h := &handlers.AuthHandler{
				DB: db,
			}

			validateResponse(t, h.HandleRegistration(tt.params), tt.wantBody, tt.wantStatus)
		})
	}
}

func TestAuthHandler_Authenticate(t *testing.T) {
	type mockParams struct {
		login string
		res   *dbModels.Credentials
		err   error
	}
	tests := []struct {
		name       string
		params     auth.PostLoginParams
		mockParams mockParams
		wantStatus int
	}{
		{
			name: "ok",
			params: auth.PostLoginParams{
				Body: &apiModels.AuthRequest{
					Login:    strToPtr("test-login"),
					Password: strToPtr("test-pass"),
				},
			},
			mockParams: mockParams{
				login: "test-login",
				res: &dbModels.Credentials{
					Login:    "test-login",
					PassHash: "$2y$12$fI3697OVT/mYWctZLnVV.uSq.zPWua2xopLRpSRRVVCIoYYvBRJxi",
				},
				err: nil,
			},
			wantStatus: 200,
		},

		{
			name: "error password",
			params: auth.PostLoginParams{
				Body: &apiModels.AuthRequest{
					Login:    strToPtr("test-login"),
					Password: strToPtr("test-pass"),
				},
			},
			mockParams: mockParams{
				login: "test-login",
				res: &dbModels.Credentials{
					Login:    "test-login",
					PassHash: "qwerty",
				},
				err: nil,
			},
			wantStatus: 403,
		},

		{
			name: "error not found",
			params: auth.PostLoginParams{
				Body: &apiModels.AuthRequest{
					Login:    strToPtr("test-login"),
					Password: strToPtr("test-pass"),
				},
			},
			mockParams: mockParams{
				login: "test-login",
				res:   nil,
				err:   ErrTestNotFound,
			},
			wantStatus: 403,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := new(mocks.DatabaseAuthService)
			db.On("GetUser", tt.mockParams.login).Return(tt.mockParams.res, tt.mockParams.err)

			h := &handlers.AuthHandler{
				DB: db,
			}

			validateResponse(t, h.HandleLogin(tt.params), "", tt.wantStatus)
		})
	}
}
