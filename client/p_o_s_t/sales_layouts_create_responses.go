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

// SalesLayoutsCreateReader is a Reader for the SalesLayoutsCreate structure.
type SalesLayoutsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalesLayoutsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalesLayoutsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSalesLayoutsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSalesLayoutsCreateOK creates a SalesLayoutsCreateOK with default headers values
func NewSalesLayoutsCreateOK() *SalesLayoutsCreateOK {
	return &SalesLayoutsCreateOK{}
}

/*
SalesLayoutsCreateOK describes a response with status code 200, with default header values.

OK
*/
type SalesLayoutsCreateOK struct {
	Payload *models.SalesLayout
}

// IsSuccess returns true when this sales layouts create o k response has a 2xx status code
func (o *SalesLayoutsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sales layouts create o k response has a 3xx status code
func (o *SalesLayoutsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sales layouts create o k response has a 4xx status code
func (o *SalesLayoutsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sales layouts create o k response has a 5xx status code
func (o *SalesLayoutsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sales layouts create o k response a status code equal to that given
func (o *SalesLayoutsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sales layouts create o k response
func (o *SalesLayoutsCreateOK) Code() int {
	return 200
}

func (o *SalesLayoutsCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/SalesLayouts/Setup][%d] salesLayoutsCreateOK %s", 200, payload)
}

func (o *SalesLayoutsCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/SalesLayouts/Setup][%d] salesLayoutsCreateOK %s", 200, payload)
}

func (o *SalesLayoutsCreateOK) GetPayload() *models.SalesLayout {
	return o.Payload
}

func (o *SalesLayoutsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SalesLayout)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSalesLayoutsCreateDefault creates a SalesLayoutsCreateDefault with default headers values
func NewSalesLayoutsCreateDefault(code int) *SalesLayoutsCreateDefault {
	return &SalesLayoutsCreateDefault{
		_statusCode: code,
	}
}

/*
SalesLayoutsCreateDefault describes a response with status code -1, with default header values.

Error
*/
type SalesLayoutsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this sales layouts create default response has a 2xx status code
func (o *SalesLayoutsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this sales layouts create default response has a 3xx status code
func (o *SalesLayoutsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this sales layouts create default response has a 4xx status code
func (o *SalesLayoutsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this sales layouts create default response has a 5xx status code
func (o *SalesLayoutsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this sales layouts create default response a status code equal to that given
func (o *SalesLayoutsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the sales layouts create default response
func (o *SalesLayoutsCreateDefault) Code() int {
	return o._statusCode
}

func (o *SalesLayoutsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/SalesLayouts/Setup][%d] SalesLayouts_Create default %s", o._statusCode, payload)
}

func (o *SalesLayoutsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/SalesLayouts/Setup][%d] SalesLayouts_Create default %s", o._statusCode, payload)
}

func (o *SalesLayoutsCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SalesLayoutsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
