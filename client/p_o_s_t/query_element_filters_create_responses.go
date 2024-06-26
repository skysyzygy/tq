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

// QueryElementFiltersCreateReader is a Reader for the QueryElementFiltersCreate structure.
type QueryElementFiltersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryElementFiltersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryElementFiltersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQueryElementFiltersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQueryElementFiltersCreateOK creates a QueryElementFiltersCreateOK with default headers values
func NewQueryElementFiltersCreateOK() *QueryElementFiltersCreateOK {
	return &QueryElementFiltersCreateOK{}
}

/*
QueryElementFiltersCreateOK describes a response with status code 200, with default header values.

OK
*/
type QueryElementFiltersCreateOK struct {
	Payload *models.QueryElementFilter
}

// IsSuccess returns true when this query element filters create o k response has a 2xx status code
func (o *QueryElementFiltersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this query element filters create o k response has a 3xx status code
func (o *QueryElementFiltersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this query element filters create o k response has a 4xx status code
func (o *QueryElementFiltersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this query element filters create o k response has a 5xx status code
func (o *QueryElementFiltersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this query element filters create o k response a status code equal to that given
func (o *QueryElementFiltersCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the query element filters create o k response
func (o *QueryElementFiltersCreateOK) Code() int {
	return 200
}

func (o *QueryElementFiltersCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElementFilters][%d] queryElementFiltersCreateOK %s", 200, payload)
}

func (o *QueryElementFiltersCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElementFilters][%d] queryElementFiltersCreateOK %s", 200, payload)
}

func (o *QueryElementFiltersCreateOK) GetPayload() *models.QueryElementFilter {
	return o.Payload
}

func (o *QueryElementFiltersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.QueryElementFilter)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryElementFiltersCreateDefault creates a QueryElementFiltersCreateDefault with default headers values
func NewQueryElementFiltersCreateDefault(code int) *QueryElementFiltersCreateDefault {
	return &QueryElementFiltersCreateDefault{
		_statusCode: code,
	}
}

/*
QueryElementFiltersCreateDefault describes a response with status code -1, with default header values.

Error
*/
type QueryElementFiltersCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this query element filters create default response has a 2xx status code
func (o *QueryElementFiltersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this query element filters create default response has a 3xx status code
func (o *QueryElementFiltersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this query element filters create default response has a 4xx status code
func (o *QueryElementFiltersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this query element filters create default response has a 5xx status code
func (o *QueryElementFiltersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this query element filters create default response a status code equal to that given
func (o *QueryElementFiltersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the query element filters create default response
func (o *QueryElementFiltersCreateDefault) Code() int {
	return o._statusCode
}

func (o *QueryElementFiltersCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElementFilters][%d] QueryElementFilters_Create default %s", o._statusCode, payload)
}

func (o *QueryElementFiltersCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/QueryElementFilters][%d] QueryElementFilters_Create default %s", o._statusCode, payload)
}

func (o *QueryElementFiltersCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QueryElementFiltersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
