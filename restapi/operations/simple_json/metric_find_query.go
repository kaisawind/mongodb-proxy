// Code generated by go-swagger; DO NOT EDIT.

package simple_json

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// MetricFindQueryHandlerFunc turns a function with the right signature into a metric find query handler
type MetricFindQueryHandlerFunc func(MetricFindQueryParams) middleware.Responder

// Handle executing the request and returning a response
func (fn MetricFindQueryHandlerFunc) Handle(params MetricFindQueryParams) middleware.Responder {
	return fn(params)
}

// MetricFindQueryHandler interface for that can handle valid metric find query params
type MetricFindQueryHandler interface {
	Handle(MetricFindQueryParams) middleware.Responder
}

// NewMetricFindQuery creates a new http.Handler for the metric find query operation
func NewMetricFindQuery(ctx *middleware.Context, handler MetricFindQueryHandler) *MetricFindQuery {
	return &MetricFindQuery{Context: ctx, Handler: handler}
}

/*MetricFindQuery swagger:route POST /search SimpleJSON metricFindQuery

find metric options

used by the find metric options on the query tab in panels.

*/
type MetricFindQuery struct {
	Context *middleware.Context
	Handler MetricFindQueryHandler
}

func (o *MetricFindQuery) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewMetricFindQueryParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
