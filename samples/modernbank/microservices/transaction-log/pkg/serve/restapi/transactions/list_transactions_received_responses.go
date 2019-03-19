// Code generated by go-swagger; DO NOT EDIT.

package transactions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	model "github.com/tetrateio/training/samples/modernbank/microservices/transaction-log/pkg/model"
)

// ListTransactionsReceivedOKCode is the HTTP code returned for type ListTransactionsReceivedOK
const ListTransactionsReceivedOKCode int = 200

/*ListTransactionsReceivedOK Success!

swagger:response listTransactionsReceivedOK
*/
type ListTransactionsReceivedOK struct {

	/*
	  In: Body
	*/
	Payload []*model.Transaction `json:"body,omitempty"`
}

// NewListTransactionsReceivedOK creates ListTransactionsReceivedOK with default headers values
func NewListTransactionsReceivedOK() *ListTransactionsReceivedOK {

	return &ListTransactionsReceivedOK{}
}

// WithPayload adds the payload to the list transactions received o k response
func (o *ListTransactionsReceivedOK) WithPayload(payload []*model.Transaction) *ListTransactionsReceivedOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list transactions received o k response
func (o *ListTransactionsReceivedOK) SetPayload(payload []*model.Transaction) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListTransactionsReceivedOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make([]*model.Transaction, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// ListTransactionsReceivedNotFoundCode is the HTTP code returned for type ListTransactionsReceivedNotFound
const ListTransactionsReceivedNotFoundCode int = 404

/*ListTransactionsReceivedNotFound No transactions found

swagger:response listTransactionsReceivedNotFound
*/
type ListTransactionsReceivedNotFound struct {
}

// NewListTransactionsReceivedNotFound creates ListTransactionsReceivedNotFound with default headers values
func NewListTransactionsReceivedNotFound() *ListTransactionsReceivedNotFound {

	return &ListTransactionsReceivedNotFound{}
}

// WriteResponse to the client
func (o *ListTransactionsReceivedNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}

// ListTransactionsReceivedInternalServerErrorCode is the HTTP code returned for type ListTransactionsReceivedInternalServerError
const ListTransactionsReceivedInternalServerErrorCode int = 500

/*ListTransactionsReceivedInternalServerError Internal server error

swagger:response listTransactionsReceivedInternalServerError
*/
type ListTransactionsReceivedInternalServerError struct {
}

// NewListTransactionsReceivedInternalServerError creates ListTransactionsReceivedInternalServerError with default headers values
func NewListTransactionsReceivedInternalServerError() *ListTransactionsReceivedInternalServerError {

	return &ListTransactionsReceivedInternalServerError{}
}

// WriteResponse to the client
func (o *ListTransactionsReceivedInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(500)
}
