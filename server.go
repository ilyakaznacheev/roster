package roster

import (
	"github.com/go-openapi/loads"

	"github.com/ilyakaznacheev/roster/internal/config"
	"github.com/ilyakaznacheev/roster/internal/database"
	"github.com/ilyakaznacheev/roster/internal/handlers"
	"github.com/ilyakaznacheev/roster/internal/restapi"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/auth"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/roster"
)

// Run starts the server
func Run(cfg config.Application) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	// setup external services and dependencies
	mgoDB, err := database.NewMongoHandler(cfg.Database.MongoURI)
	if err != nil {
		return err
	}
	rh := handlers.NewRosterHandler(mgoDB)
	ah := handlers.NewAuthHandler()

	api := operations.NewRosterAPI(swaggerSpec)

	// routing
	api.RosterGetRostersHandler = roster.GetRostersHandlerFunc(rh.GetRosterAll)
	api.RosterGetRostersIDHandler = roster.GetRostersIDHandlerFunc(rh.GetRosterOne)
	api.AuthPostLoginHandler = auth.PostLoginHandlerFunc(ah.HandleLogin)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = cfg.Server.Port
	server.Host = cfg.Server.Host

	// server.ConfigureFlags()
	// server.ConfigureAPI()

	return server.Serve()
}
