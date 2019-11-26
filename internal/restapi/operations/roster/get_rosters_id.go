// Code generated by go-swagger; DO NOT EDIT.

package roster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetRostersIDHandlerFunc turns a function with the right signature into a get rosters ID handler
type GetRostersIDHandlerFunc func(GetRostersIDParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetRostersIDHandlerFunc) Handle(params GetRostersIDParams) middleware.Responder {
	return fn(params)
}

// GetRostersIDHandler interface for that can handle valid get rosters ID params
type GetRostersIDHandler interface {
	Handle(GetRostersIDParams) middleware.Responder
}

// NewGetRostersID creates a new http.Handler for the get rosters ID operation
func NewGetRostersID(ctx *middleware.Context, handler GetRostersIDHandler) *GetRostersID {
	return &GetRostersID{Context: ctx, Handler: handler}
}

/*GetRostersID swagger:route GET /rosters/{id} roster getRostersId

Get a roster

Returns a roster with all playes

*/
type GetRostersID struct {
	Context *middleware.Context
	Handler GetRostersIDHandler
}

func (o *GetRostersID) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetRostersIDParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
