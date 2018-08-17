// Code generated by go-swagger; DO NOT EDIT.

package consumers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// FindConsumerHandlerFunc turns a function with the right signature into a find consumer handler
type FindConsumerHandlerFunc func(FindConsumerParams) middleware.Responder

// Handle executing the request and returning a response
func (fn FindConsumerHandlerFunc) Handle(params FindConsumerParams) middleware.Responder {
	return fn(params)
}

// FindConsumerHandler interface for that can handle valid find consumer params
type FindConsumerHandler interface {
	Handle(FindConsumerParams) middleware.Responder
}

// NewFindConsumer creates a new http.Handler for the find consumer operation
func NewFindConsumer(ctx *middleware.Context, handler FindConsumerHandler) *FindConsumer {
	return &FindConsumer{Context: ctx, Handler: handler}
}

/*FindConsumer swagger:route GET /consumers/{id} consumers findConsumer

FindConsumer find consumer API

*/
type FindConsumer struct {
	Context *middleware.Context
	Handler FindConsumerHandler
}

func (o *FindConsumer) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewFindConsumerParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}