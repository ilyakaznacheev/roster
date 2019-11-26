package roster

import (
	"github.com/go-openapi/loads"

	"github.com/ilyakaznacheev/roster/internal/config"
	"github.com/ilyakaznacheev/roster/internal/database"
	"github.com/ilyakaznacheev/roster/internal/handlers"
	"github.com/ilyakaznacheev/roster/internal/restapi"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/roster"
)

// Run starts the server
func Run(cfg config.Application) error {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		return err
	}

	mgoDB, err := database.NewMongoHandler(cfg.Database.MongoURI)
	if err != nil {
		return err
	}

	rh := handlers.NewRosterHandler(mgoDB)

	api := operations.NewRosterAPI(swaggerSpec)

	api.RosterGetRostersHandler = roster.GetRostersHandlerFunc(rh.GetRosterAll)
	api.RosterGetRostersIDHandler = roster.GetRostersIDHandlerFunc(rh.GetRosterOne)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = cfg.Server.Port
	server.Host = cfg.Server.Host

	// type Config struct {
	// 	MongoURI string `long:"mongo-uri" description:"MongoDB connection URI" env:"MONGO_URI"`
	// }
	// var c Config

	// if _, err := flags.Parse(server); err != nil {
	// 	return nil
	// }
	// if _, err := flags.Parse(&c); err != nil {
	// 	return nil
	// }

	// server.ConfigureFlags()
	// server.ConfigureAPI()

	return server.Serve()
}
