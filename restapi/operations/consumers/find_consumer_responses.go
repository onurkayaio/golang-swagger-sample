// Code generated by go-swagger; DO NOT EDIT.

package consumers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "go-swagger-sample/models"
)

// FindConsumerOKCode is the HTTP code returned for type FindConsumerOK
const FindConsumerOKCode int = 200

/*FindConsumerOK OK

swagger:response findConsumerOK
*/
type FindConsumerOK struct {

	/*
	  In: Body
	*/
	Payload *models.Consumer `json:"body,omitempty"`
}

// NewFindConsumerOK creates FindConsumerOK with default headers values
func NewFindConsumerOK() *FindConsumerOK {

	return &FindConsumerOK{}
}

// WithPayload adds the payload to the find consumer o k response
func (o *FindConsumerOK) WithPayload(payload *models.Consumer) *FindConsumerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find consumer o k response
func (o *FindConsumerOK) SetPayload(payload *models.Consumer) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindConsumerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*FindConsumerDefault generic error response

swagger:response findConsumerDefault
*/
type FindConsumerDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewFindConsumerDefault creates FindConsumerDefault with default headers values
func NewFindConsumerDefault(code int) *FindConsumerDefault {
	if code <= 0 {
		code = 500
	}

	return &FindConsumerDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the find consumer default response
func (o *FindConsumerDefault) WithStatusCode(code int) *FindConsumerDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the find consumer default response
func (o *FindConsumerDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the find consumer default response
func (o *FindConsumerDefault) WithPayload(payload *models.Error) *FindConsumerDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the find consumer default response
func (o *FindConsumerDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *FindConsumerDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}