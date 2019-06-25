// Code generated by go-swagger; DO NOT EDIT.

package simple_json

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// TestDatasourceHandlerFunc turns a function with the right signature into a test datasource handler
type TestDatasourceHandlerFunc func(TestDatasourceParams) middleware.Responder

// Handle executing the request and returning a response
func (fn TestDatasourceHandlerFunc) Handle(params TestDatasourceParams) middleware.Responder {
	return fn(params)
}

// TestDatasourceHandler interface for that can handle valid test datasource params
type TestDatasourceHandler interface {
	Handle(TestDatasourceParams) middleware.Responder
}

// NewTestDatasource creates a new http.Handler for the test datasource operation
func NewTestDatasource(ctx *middleware.Context, handler TestDatasourceHandler) *TestDatasource {
	return &TestDatasource{Context: ctx, Handler: handler}
}

/*TestDatasource swagger:route GET / SimpleJSON testDatasource

test connection

should return 200 ok. Used for "Test connection" on the datasource config page.

*/
type TestDatasource struct {
	Context *middleware.Context
	Handler TestDatasourceHandler
}

func (o *TestDatasource) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewTestDatasourceParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}