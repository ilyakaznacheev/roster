package roster

import (
	"log"

	"github.com/go-openapi/loads"

	"github.com/ilyakaznacheev/roster/internal/api/restapi"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/auth"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/player"
	"github.com/ilyakaznacheev/roster/internal/api/restapi/operations/roster"
	"github.com/ilyakaznacheev/roster/internal/config"
	"github.com/ilyakaznacheev/roster/internal/database"
	"github.com/ilyakaznacheev/roster/internal/handlers"
)

// Run starts the server
func Run(cfg config.Application) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	// setup external services and dependencies
	mgoDB, err := database.NewMongoHandler(cfg.Database)
	if err != nil {
		return err
	}
	rh := handlers.NewRosterHandler(mgoDB)
	ah := handlers.NewAuthHandler(cfg.Server.AuthKey, mgoDB)

	api := operations.NewRosterAPI(swaggerSpec)

	// routing
	api.RosterGetRostersHandler = roster.GetRostersHandlerFunc(rh.GetRosterAll)
	api.RosterGetRostersIDHandler = roster.GetRostersIDHandlerFunc(rh.GetRosterOne)
	api.RosterGetRostersIDActiveHandler = roster.GetRostersIDActiveHandlerFunc(rh.GetRosterActive)
	api.RosterGetRostersIDBenchedHandler = roster.GetRostersIDBenchedHandlerFunc(rh.GetRosterBenched)
	api.PlayerPostRostersIDAddPlayerHandler = player.PostRostersIDAddPlayerHandlerFunc(rh.AddPayer)
	api.PlayerPostRostersIDRearrangeHandler = player.PostRostersIDRearrangeHandlerFunc(rh.RearrangeRoster)
	api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(ah.HandleLogin)
	api.AuthPostRegisterHandler = auth.PostRegisterHandlerFunc(ah.HandleRegistration)

	api.BearerAuth = ah.Authenticate
	api.Logger = log.Printf

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = cfg.Server.Port
	server.Host = cfg.Server.Host

	// server.ConfigureFlags()

	server.ConfigureAPI()

	return server.Serve()
}
