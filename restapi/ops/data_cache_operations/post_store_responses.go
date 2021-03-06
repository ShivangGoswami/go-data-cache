// Code generated by go-swagger; DO NOT EDIT.

package data_cache_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "datacache/models"
)

// PostStoreCreatedCode is the HTTP code returned for type PostStoreCreated
const PostStoreCreatedCode int = 201

/*PostStoreCreated Cache Data stored Successfully

swagger:response postStoreCreated
*/
type PostStoreCreated struct {
}

// NewPostStoreCreated creates PostStoreCreated with default headers values
func NewPostStoreCreated() *PostStoreCreated {

	return &PostStoreCreated{}
}

// WriteResponse to the client
func (o *PostStoreCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(201)
}

/*PostStoreDefault Generic error response

swagger:response postStoreDefault
*/
type PostStoreDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewPostStoreDefault creates PostStoreDefault with default headers values
func NewPostStoreDefault(code int) *PostStoreDefault {
	if code <= 0 {
		code = 500
	}

	return &PostStoreDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the post store default response
func (o *PostStoreDefault) WithStatusCode(code int) *PostStoreDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the post store default response
func (o *PostStoreDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the post store default response
func (o *PostStoreDefault) WithPayload(payload *models.Error) *PostStoreDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post store default response
func (o *PostStoreDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostStoreDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
