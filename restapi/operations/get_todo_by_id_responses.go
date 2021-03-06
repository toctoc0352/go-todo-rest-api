// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "github.com/toctoc0352/go-todo-rest-api/models"
)

// GetTodoByIDOKCode is the HTTP code returned for type GetTodoByIDOK
const GetTodoByIDOKCode int = 200

/*GetTodoByIDOK Success

swagger:response getTodoByIdOK
*/
type GetTodoByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.Todo `json:"body,omitempty"`
}

// NewGetTodoByIDOK creates GetTodoByIDOK with default headers values
func NewGetTodoByIDOK() *GetTodoByIDOK {

	return &GetTodoByIDOK{}
}

// WithPayload adds the payload to the get todo by Id o k response
func (o *GetTodoByIDOK) WithPayload(payload *models.Todo) *GetTodoByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get todo by Id o k response
func (o *GetTodoByIDOK) SetPayload(payload *models.Todo) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetTodoByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
