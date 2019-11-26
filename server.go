package roster

import (
	"github.com/go-openapi/loads"

	"github.com/ilyakaznacheev/roster/internal/handlers"
	"github.com/ilyakaznacheev/roster/internal/restapi"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations"
	"github.com/ilyakaznacheev/roster/internal/restapi/operations/roster"
)

// Run starts the server
func Run() error {

	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return err
	}

	rh := &handlers.RosterHandler{}

	api := operations.NewRosterAPI(swaggerSpec)
	api.RosterGetRosterHandler = roster.GetRosterHandlerFunc(rh.GetRoster)

	return nil
}
