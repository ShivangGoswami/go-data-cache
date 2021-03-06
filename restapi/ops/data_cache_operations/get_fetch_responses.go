// Code generated by go-swagger; DO NOT EDIT.

package data_cache_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "datacache/models"
)

// GetFetchOKCode is the HTTP code returned for type GetFetchOK
const GetFetchOKCode int = 200

/*GetFetchOK key value data

swagger:response getFetchOK
*/
type GetFetchOK struct {

	/*
	  In: Body
	*/
	Payload models.CacheCollection `json:"body,omitempty"`
}

// NewGetFetchOK creates GetFetchOK with default headers values
func NewGetFetchOK() *GetFetchOK {

	return &GetFetchOK{}
}

// WithPayload adds the payload to the get fetch o k response
func (o *GetFetchOK) WithPayload(payload models.CacheCollection) *GetFetchOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fetch o k response
func (o *GetFetchOK) SetPayload(payload models.CacheCollection) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFetchOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = models.CacheCollection{}
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*GetFetchDefault Generic error response

swagger:response getFetchDefault
*/
type GetFetchDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetFetchDefault creates GetFetchDefault with default headers values
func NewGetFetchDefault(code int) *GetFetchDefault {
	if code <= 0 {
		code = 500
	}

	return &GetFetchDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the get fetch default response
func (o *GetFetchDefault) WithStatusCode(code int) *GetFetchDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the get fetch default response
func (o *GetFetchDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the get fetch default response
func (o *GetFetchDefault) WithPayload(payload *models.Error) *GetFetchDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get fetch default response
func (o *GetFetchDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetFetchDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
