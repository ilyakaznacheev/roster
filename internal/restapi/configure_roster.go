// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"github.com/ilyakaznacheev/roster/internal/restapi/operations"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/auth"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/player"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/roster"
)

//go:generate swagger generate server --target ../../internal --name Roster --spec ../../swagger.yml --exclude-main

func configureFlags(api *operations.RosterAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.RosterAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	// Applies when the "Authorization" header is set
	api.BearerAuth = func(token string) (interface{}, error) {
		return nil, errors.NotImplemented("api key auth (Bearer) Authorization from header param [Authorization] has not yet been implemented")
	}

	// Set your custom authorizer if needed. Default one is security.Authorized()
	// Expected interface runtime.Authorizer
	//
	// Example:
	// api.APIAuthorizer = security.Authorized()
	if api.RosterGetRosterHandler == nil {
		api.RosterGetRosterHandler = roster.GetRosterHandlerFunc(func(params roster.GetRosterParams) middleware.Responder {
			return middleware.NotImplemented("operation roster.GetRoster has not yet been implemented")
		})
	}
	if api.RosterGetRosterRosteridHandler == nil {
		api.RosterGetRosterRosteridHandler = roster.GetRosterRosteridHandlerFunc(func(params roster.GetRosterRosteridParams) middleware.Responder {
			return middleware.NotImplemented("operation roster.GetRosterRosterid has not yet been implemented")
		})
	}
	if api.RosterGetRosterRosteridActiveHandler == nil {
		api.RosterGetRosterRosteridActiveHandler = roster.GetRosterRosteridActiveHandlerFunc(func(params roster.GetRosterRosteridActiveParams) middleware.Responder {
			return middleware.NotImplemented("operation roster.GetRosterRosteridActive has not yet been implemented")
		})
	}
	if api.RosterGetRosterRosteridBenchedHandler == nil {
		api.RosterGetRosterRosteridBenchedHandler = roster.GetRosterRosteridBenchedHandlerFunc(func(params roster.GetRosterRosteridBenchedParams) middleware.Responder {
			return middleware.NotImplemented("operation roster.GetRosterRosteridBenched has not yet been implemented")
		})
	}
	if api.AuthPostLoginHandler == nil {
		api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(func(params auth.PostLoginParams) middleware.Responder {
			return middleware.NotImplemented("operation auth.PostLogin has not yet been implemented")
		})
	}
	if api.PlayerPostRosterRosteridAddPlayerHandler == nil {
		api.PlayerPostRosterRosteridAddPlayerHandler = player.PostRosterRosteridAddPlayerHandlerFunc(func(params player.PostRosterRosteridAddPlayerParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation player.PostRosterRosteridAddPlayer has not yet been implemented")
		})
	}
	if api.PlayerPostRosterRosteridRearrangeHandler == nil {
		api.PlayerPostRosterRosteridRearrangeHandler = player.PostRosterRosteridRearrangeHandlerFunc(func(params player.PostRosterRosteridRearrangeParams, principal interface{}) middleware.Responder {
			return middleware.NotImplemented("operation player.PostRosterRosteridRearrange has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
