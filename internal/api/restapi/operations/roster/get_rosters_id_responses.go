// Code generated by go-swagger; DO NOT EDIT.

package roster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/ilyakaznacheev/roster/internal/api/models"
)

// GetRostersIDOKCode is the HTTP code returned for type GetRostersIDOK
const GetRostersIDOKCode int = 200

/*GetRostersIDOK successful operation

swagger:response getRostersIdOK
*/
type GetRostersIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Roster `json:"body,omitempty"`
}

// NewGetRostersIDOK creates GetRostersIDOK with default headers values
func NewGetRostersIDOK() *GetRostersIDOK {

	return &GetRostersIDOK{}
}

// WithPayload adds the payload to the get rosters Id o k response
func (o *GetRostersIDOK) WithPayload(payload *models.Roster) *GetRostersIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rosters Id o k response
func (o *GetRostersIDOK) SetPayload(payload *models.Roster) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRostersIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRostersIDNotFoundCode is the HTTP code returned for type GetRostersIDNotFound
const GetRostersIDNotFoundCode int = 404

/*GetRostersIDNotFound not found

swagger:response getRostersIdNotFound
*/
type GetRostersIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRostersIDNotFound creates GetRostersIDNotFound with default headers values
func NewGetRostersIDNotFound() *GetRostersIDNotFound {

	return &GetRostersIDNotFound{}
}

// WithPayload adds the payload to the get rosters Id not found response
func (o *GetRostersIDNotFound) WithPayload(payload *models.Error) *GetRostersIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rosters Id not found response
func (o *GetRostersIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRostersIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetRostersIDInternalServerErrorCode is the HTTP code returned for type GetRostersIDInternalServerError
const GetRostersIDInternalServerErrorCode int = 500

/*GetRostersIDInternalServerError internal server error

swagger:response getRostersIdInternalServerError
*/
type GetRostersIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetRostersIDInternalServerError creates GetRostersIDInternalServerError with default headers values
func NewGetRostersIDInternalServerError() *GetRostersIDInternalServerError {

	return &GetRostersIDInternalServerError{}
}

// WithPayload adds the payload to the get rosters Id internal server error response
func (o *GetRostersIDInternalServerError) WithPayload(payload *models.Error) *GetRostersIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get rosters Id internal server error response
func (o *GetRostersIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetRostersIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
