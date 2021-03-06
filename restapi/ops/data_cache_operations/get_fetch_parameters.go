// Code generated by go-swagger; DO NOT EDIT.

package data_cache_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetFetchParams creates a new GetFetchParams object
// no default values defined in spec.
func NewGetFetchParams() GetFetchParams {

	return GetFetchParams{}
}

// GetFetchParams contains all the bound params for the get fetch operation
// typically these are obtained from a http.Request
//
// swagger:parameters GetFetch
type GetFetchParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  Minimum: 1
	  In: query
	*/
	Index *int64
	/*
	  In: query
	*/
	Key *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewGetFetchParams() beforehand.
func (o *GetFetchParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qIndex, qhkIndex, _ := qs.GetOK("index")
	if err := o.bindIndex(qIndex, qhkIndex, route.Formats); err != nil {
		res = append(res, err)
	}

	qKey, qhkKey, _ := qs.GetOK("key")
	if err := o.bindKey(qKey, qhkKey, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindIndex binds and validates parameter Index from query.
func (o *GetFetchParams) bindIndex(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertInt64(raw)
	if err != nil {
		return errors.InvalidType("index", "query", "int64", raw)
	}
	o.Index = &value

	if err := o.validateIndex(formats); err != nil {
		return err
	}

	return nil
}

// validateIndex carries on validations for parameter Index
func (o *GetFetchParams) validateIndex(formats strfmt.Registry) error {

	if err := validate.MinimumInt("index", "query", int64(*o.Index), 1, false); err != nil {
		return err
	}

	return nil
}

// bindKey binds and validates parameter Key from query.
func (o *GetFetchParams) bindKey(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: false
	// AllowEmptyValue: false
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Key = &raw

	return nil
}
