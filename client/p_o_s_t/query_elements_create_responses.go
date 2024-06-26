// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// QueryElementsCreateReader is a Reader for the QueryElementsCreate structure.
type QueryElementsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryElementsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryElementsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQueryElementsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryElementsCreateOK creates a QueryElementsCreateOK with default headers values
func NewQueryElementsCreateOK() *QueryElementsCreateOK {
	return &QueryElementsCreateOK{}
}

/*
QueryElementsCreateOK describes a response with status code 200, with default header values.

OK
*/
type QueryElementsCreateOK struct {
	Payload *models.QueryElement
}

// IsSuccess returns true when this query elements create o k response has a 2xx status code
func (o *QueryElementsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query elements create o k response has a 3xx status code
func (o *QueryElementsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query elements create o k response has a 4xx status code
func (o *QueryElementsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query elements create o k response has a 5xx status code
func (o *QueryElementsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query elements create o k response a status code equal to that given
func (o *QueryElementsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query elements create o k response
func (o *QueryElementsCreateOK) Code() int {
	return 200
}

func (o *QueryElementsCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElements][%d] queryElementsCreateOK %s", 200, payload)
}

func (o *QueryElementsCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElements][%d] queryElementsCreateOK %s", 200, payload)
}

func (o *QueryElementsCreateOK) GetPayload() *models.QueryElement {
	return o.Payload
}

func (o *QueryElementsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryElement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryElementsCreateDefault creates a QueryElementsCreateDefault with default headers values
func NewQueryElementsCreateDefault(code int) *QueryElementsCreateDefault {
	return &QueryElementsCreateDefault{
		_statusCode: code,
	}
}

/*
QueryElementsCreateDefault describes a response with status code -1, with default header values.

Error
*/
type QueryElementsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this query elements create default response has a 2xx status code
func (o *QueryElementsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query elements create default response has a 3xx status code
func (o *QueryElementsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query elements create default response has a 4xx status code
func (o *QueryElementsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query elements create default response has a 5xx status code
func (o *QueryElementsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query elements create default response a status code equal to that given
func (o *QueryElementsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query elements create default response
func (o *QueryElementsCreateDefault) Code() int {
	return o._statusCode
}

func (o *QueryElementsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElements][%d] QueryElements_Create default %s", o._statusCode, payload)
}

func (o *QueryElementsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElements][%d] QueryElements_Create default %s", o._statusCode, payload)
}

func (o *QueryElementsCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryElementsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
