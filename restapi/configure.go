// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	mws "github.com/kaisawind/mongodb-proxy/middlewares"
	"github.com/kaisawind/mongodb-proxy/restapi/operations"
	"github.com/kaisawind/mongodb-proxy/restapi/operations/simple_json"
	"github.com/kaisawind/mongodb-proxy/server"
)

//go:generate swagger generate server --target .. --name api --spec ../swagger/swagger.yaml --model-package v1 --principal v1.Principal

func configureFlags(api *operations.API) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.API) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.SimpleJSONAnnotationQueryHandler = simple_json.AnnotationQueryHandlerFunc(func(params simple_json.AnnotationQueryParams) middleware.Responder {
		return server.AnnotationQuery(params)
	})
	api.SimpleJSONMetricFindQueryHandler = simple_json.MetricFindQueryHandlerFunc(func(params simple_json.MetricFindQueryParams) middleware.Responder {
		return server.MetricFindQuery(params)
	})
	api.SimpleJSONQueryHandler = simple_json.QueryHandlerFunc(func(params simple_json.QueryParams) middleware.Responder {
		return server.Query(params)
	})
	api.SimpleJSONTestDatasourceHandler = simple_json.TestDatasourceHandlerFunc(func(params simple_json.TestDatasourceParams) middleware.Responder {
		return server.TestDatasource(params)
	})

	api.ServerShutdown = func() {}

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
	return mws.Limiter(handler)
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return mws.HandlePanic(mws.LogViaLogrus(mws.Cross(handler)))
}
