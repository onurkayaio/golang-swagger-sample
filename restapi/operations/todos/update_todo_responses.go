// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "go-swagger-sample/models"
)

// UpdateTodoOKCode is the HTTP code returned for type UpdateTodoOK
const UpdateTodoOKCode int = 200

/*UpdateTodoOK OK

swagger:response updateTodoOK
*/
type UpdateTodoOK struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewUpdateTodoOK creates UpdateTodoOK with default headers values
func NewUpdateTodoOK() *UpdateTodoOK {

	return &UpdateTodoOK{}
}

// WithPayload adds the payload to the update todo o k response
func (o *UpdateTodoOK) WithPayload(payload *models.Item) *UpdateTodoOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update todo o k response
func (o *UpdateTodoOK) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTodoOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateTodoDefault error

swagger:response updateTodoDefault
*/
type UpdateTodoDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateTodoDefault creates UpdateTodoDefault with default headers values
func NewUpdateTodoDefault(code int) *UpdateTodoDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateTodoDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update todo default response
func (o *UpdateTodoDefault) WithStatusCode(code int) *UpdateTodoDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update todo default response
func (o *UpdateTodoDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update todo default response
func (o *UpdateTodoDefault) WithPayload(payload *models.Error) *UpdateTodoDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update todo default response
func (o *UpdateTodoDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateTodoDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
