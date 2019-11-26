// Code generated by go-swagger; DO NOT EDIT.

package player

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ilyakaznacheev/roster/internal/models"
)

// PostRostersIDAddPlayerCreatedCode is the HTTP code returned for type PostRostersIDAddPlayerCreated
const PostRostersIDAddPlayerCreatedCode int = 201

/*PostRostersIDAddPlayerCreated created

swagger:response postRostersIdAddPlayerCreated
*/
type PostRostersIDAddPlayerCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Player `json:"body,omitempty"`
}

// NewPostRostersIDAddPlayerCreated creates PostRostersIDAddPlayerCreated with default headers values
func NewPostRostersIDAddPlayerCreated() *PostRostersIDAddPlayerCreated {

	return &PostRostersIDAddPlayerCreated{}
}

// WithPayload adds the payload to the post rosters Id add player created response
func (o *PostRostersIDAddPlayerCreated) WithPayload(payload *models.Player) *PostRostersIDAddPlayerCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rosters Id add player created response
func (o *PostRostersIDAddPlayerCreated) SetPayload(payload *models.Player) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRostersIDAddPlayerCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRostersIDAddPlayerForbiddenCode is the HTTP code returned for type PostRostersIDAddPlayerForbidden
const PostRostersIDAddPlayerForbiddenCode int = 403

/*PostRostersIDAddPlayerForbidden forbidden

swagger:response postRostersIdAddPlayerForbidden
*/
type PostRostersIDAddPlayerForbidden struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRostersIDAddPlayerForbidden creates PostRostersIDAddPlayerForbidden with default headers values
func NewPostRostersIDAddPlayerForbidden() *PostRostersIDAddPlayerForbidden {

	return &PostRostersIDAddPlayerForbidden{}
}

// WithPayload adds the payload to the post rosters Id add player forbidden response
func (o *PostRostersIDAddPlayerForbidden) WithPayload(payload *models.Error) *PostRostersIDAddPlayerForbidden {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rosters Id add player forbidden response
func (o *PostRostersIDAddPlayerForbidden) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRostersIDAddPlayerForbidden) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(403)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRostersIDAddPlayerNotFoundCode is the HTTP code returned for type PostRostersIDAddPlayerNotFound
const PostRostersIDAddPlayerNotFoundCode int = 404

/*PostRostersIDAddPlayerNotFound not found

swagger:response postRostersIdAddPlayerNotFound
*/
type PostRostersIDAddPlayerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRostersIDAddPlayerNotFound creates PostRostersIDAddPlayerNotFound with default headers values
func NewPostRostersIDAddPlayerNotFound() *PostRostersIDAddPlayerNotFound {

	return &PostRostersIDAddPlayerNotFound{}
}

// WithPayload adds the payload to the post rosters Id add player not found response
func (o *PostRostersIDAddPlayerNotFound) WithPayload(payload *models.Error) *PostRostersIDAddPlayerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rosters Id add player not found response
func (o *PostRostersIDAddPlayerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRostersIDAddPlayerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PostRostersIDAddPlayerInternalServerErrorCode is the HTTP code returned for type PostRostersIDAddPlayerInternalServerError
const PostRostersIDAddPlayerInternalServerErrorCode int = 500

/*PostRostersIDAddPlayerInternalServerError internal server error

swagger:response postRostersIdAddPlayerInternalServerError
*/
type PostRostersIDAddPlayerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostRostersIDAddPlayerInternalServerError creates PostRostersIDAddPlayerInternalServerError with default headers values
func NewPostRostersIDAddPlayerInternalServerError() *PostRostersIDAddPlayerInternalServerError {

	return &PostRostersIDAddPlayerInternalServerError{}
}

// WithPayload adds the payload to the post rosters Id add player internal server error response
func (o *PostRostersIDAddPlayerInternalServerError) WithPayload(payload *models.Error) *PostRostersIDAddPlayerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post rosters Id add player internal server error response
func (o *PostRostersIDAddPlayerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostRostersIDAddPlayerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
