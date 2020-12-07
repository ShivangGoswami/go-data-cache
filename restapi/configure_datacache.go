// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"

	pkgapi "datacache/api"
	"datacache/customerror"
	"datacache/models"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"datacache/restapi/ops"
	"datacache/restapi/ops/data_cache_operations"
)

//go:generate swagger generate server --target ..\..\datacache --name Datacache --spec ..\api.yml --api-package ops

func configureFlags(api *ops.DatacacheAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *ops.DatacacheAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	svc := pkgapi.NewService()

	api.DataCacheOperationsGetFetchHandler = data_cache_operations.GetFetchHandlerFunc(func(params data_cache_operations.GetFetchParams) middleware.Responder {
		log.Printf("Store Request Metadata:%#v\n", params.HTTPRequest)
		return data_cache_operations.NewGetFetchOK()
	})

	api.DataCacheOperationsPostStoreHandler = data_cache_operations.PostStoreHandlerFunc(func(params data_cache_operations.PostStoreParams) middleware.Responder {
		log.Printf("Store Request Metadata:%#v\n", params.HTTPRequest)
		log.Println("Store Request Received--->", "key:", *params.Params.Key, "value:", params.Params.Value, "expiration:", *params.Params.Expiration)
		err := svc.StoreData(params)
		if err != nil {
			if err == customerror.MongoStoreException {
				return data_cache_operations.NewPostStoreDefault(500).WithPayload(&models.Error{500, &[]string{"Error while writing value to database"}[0]})
			} else if err == customerror.TimeParseError {
				return data_cache_operations.NewPostStoreDefault(422).WithPayload(&models.Error{422, &[]string{"Error while Parsing expiry duration"}[0]})
			}
		}
		return data_cache_operations.NewPostStoreCreated()
	})

	api.ServerShutdown = func() {
		svc.DestroyService()
		log.Println("System: Exit Log")
		svc.Memory.Range(func(key, value interface{}) bool {
			if temp1, ok := value.(*models.Cache); ok {
				if temp2, ok := svc.Timestamp.Load(key); ok {
					fmt.Println(key, "-->", "Value:", temp1.Value, "Expiration", *temp1.Expiration, "Access TimeStamp:", temp2)
				}
			}
			return true
		})
	}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
