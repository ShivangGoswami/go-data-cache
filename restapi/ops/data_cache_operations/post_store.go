// Code generated by go-swagger; DO NOT EDIT.

package data_cache_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PostStoreHandlerFunc turns a function with the right signature into a post store handler
type PostStoreHandlerFunc func(PostStoreParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostStoreHandlerFunc) Handle(params PostStoreParams) middleware.Responder {
	return fn(params)
}

// PostStoreHandler interface for that can handle valid post store params
type PostStoreHandler interface {
	Handle(PostStoreParams) middleware.Responder
}

// NewPostStore creates a new http.Handler for the post store operation
func NewPostStore(ctx *middleware.Context, handler PostStoreHandler) *PostStore {
	return &PostStore{Context: ctx, Handler: handler}
}

/*PostStore swagger:route POST /store DataCacheOperations postStore

Store Data

Store Data in Cache


*/
type PostStore struct {
	Context *middleware.Context
	Handler PostStoreHandler
}

func (o *PostStore) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPostStoreParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}