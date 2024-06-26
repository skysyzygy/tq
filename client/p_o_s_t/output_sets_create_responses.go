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

// OutputSetsCreateReader is a Reader for the OutputSetsCreate structure.
type OutputSetsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *OutputSetsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewOutputSetsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewOutputSetsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewOutputSetsCreateOK creates a OutputSetsCreateOK with default headers values
func NewOutputSetsCreateOK() *OutputSetsCreateOK {
	return &OutputSetsCreateOK{}
}

/*
OutputSetsCreateOK describes a response with status code 200, with default header values.

OK
*/
type OutputSetsCreateOK struct {
	Payload *models.OutputSet
}

// IsSuccess returns true when this output sets create o k response has a 2xx status code
func (o *OutputSetsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this output sets create o k response has a 3xx status code
func (o *OutputSetsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this output sets create o k response has a 4xx status code
func (o *OutputSetsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this output sets create o k response has a 5xx status code
func (o *OutputSetsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this output sets create o k response a status code equal to that given
func (o *OutputSetsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the output sets create o k response
func (o *OutputSetsCreateOK) Code() int {
	return 200
}

func (o *OutputSetsCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/OutputSets][%d] outputSetsCreateOK %s", 200, payload)
}

func (o *OutputSetsCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/OutputSets][%d] outputSetsCreateOK %s", 200, payload)
}

func (o *OutputSetsCreateOK) GetPayload() *models.OutputSet {
	return o.Payload
}

func (o *OutputSetsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OutputSet)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewOutputSetsCreateDefault creates a OutputSetsCreateDefault with default headers values
func NewOutputSetsCreateDefault(code int) *OutputSetsCreateDefault {
	return &OutputSetsCreateDefault{
		_statusCode: code,
	}
}

/*
OutputSetsCreateDefault describes a response with status code -1, with default header values.

Error
*/
type OutputSetsCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this output sets create default response has a 2xx status code
func (o *OutputSetsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this output sets create default response has a 3xx status code
func (o *OutputSetsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this output sets create default response has a 4xx status code
func (o *OutputSetsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this output sets create default response has a 5xx status code
func (o *OutputSetsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this output sets create default response a status code equal to that given
func (o *OutputSetsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the output sets create default response
func (o *OutputSetsCreateDefault) Code() int {
	return o._statusCode
}

func (o *OutputSetsCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/OutputSets][%d] OutputSets_Create default %s", o._statusCode, payload)
}

func (o *OutputSetsCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Reporting/OutputSets][%d] OutputSets_Create default %s", o._statusCode, payload)
}

func (o *OutputSetsCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *OutputSetsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
