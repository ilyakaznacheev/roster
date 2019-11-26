// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostRostersIDAddPlayerHandlerFunc turns a function with the right signature into a post rosters ID add player handler
type PostRostersIDAddPlayerHandlerFunc func(PostRostersIDAddPlayerParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn PostRostersIDAddPlayerHandlerFunc) Handle(params PostRostersIDAddPlayerParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// PostRostersIDAddPlayerHandler interface for that can handle valid post rosters ID add player params
type PostRostersIDAddPlayerHandler interface {
	Handle(PostRostersIDAddPlayerParams, interface{}) middleware.Responder
}

// NewPostRostersIDAddPlayer creates a new http.Handler for the post rosters ID add player operation
func NewPostRostersIDAddPlayer(ctx *middleware.Context, handler PostRostersIDAddPlayerHandler) *PostRostersIDAddPlayer {
	return &PostRostersIDAddPlayer{Context: ctx, Handler: handler}
}

/*PostRostersIDAddPlayer swagger:route POST /rosters/{id}/add_player player postRostersIdAddPlayer

Add a new player

Adds a new player (to a benched group)

*/
type PostRostersIDAddPlayer struct {
	Context *middleware.Context
	Handler PostRostersIDAddPlayerHandler
}

func (o *PostRostersIDAddPlayer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostRostersIDAddPlayerParams()

	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		r = aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
