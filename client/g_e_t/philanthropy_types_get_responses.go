// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// PhilanthropyTypesGetReader is a Reader for the PhilanthropyTypesGet structure.
type PhilanthropyTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhilanthropyTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhilanthropyTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPhilanthropyTypesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPhilanthropyTypesGetOK creates a PhilanthropyTypesGetOK with default headers values
func NewPhilanthropyTypesGetOK() *PhilanthropyTypesGetOK {
	return &PhilanthropyTypesGetOK{}
}

/*
PhilanthropyTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type PhilanthropyTypesGetOK struct {
	Payload *models.PhilanthropyType
}

// IsSuccess returns true when this philanthropy types get o k response has a 2xx status code
func (o *PhilanthropyTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this philanthropy types get o k response has a 3xx status code
func (o *PhilanthropyTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this philanthropy types get o k response has a 4xx status code
func (o *PhilanthropyTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this philanthropy types get o k response has a 5xx status code
func (o *PhilanthropyTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this philanthropy types get o k response a status code equal to that given
func (o *PhilanthropyTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the philanthropy types get o k response
func (o *PhilanthropyTypesGetOK) Code() int {
	return 200
}

func (o *PhilanthropyTypesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PhilanthropyTypes/{id}][%d] philanthropyTypesGetOK %s", 200, payload)
}

func (o *PhilanthropyTypesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PhilanthropyTypes/{id}][%d] philanthropyTypesGetOK %s", 200, payload)
}

func (o *PhilanthropyTypesGetOK) GetPayload() *models.PhilanthropyType {
	return o.Payload
}

func (o *PhilanthropyTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhilanthropyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPhilanthropyTypesGetDefault creates a PhilanthropyTypesGetDefault with default headers values
func NewPhilanthropyTypesGetDefault(code int) *PhilanthropyTypesGetDefault {
	return &PhilanthropyTypesGetDefault{
		_statusCode: code,
	}
}

/*
PhilanthropyTypesGetDefault describes a response with status code -1, with default header values.

Error
*/
type PhilanthropyTypesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this philanthropy types get default response has a 2xx status code
func (o *PhilanthropyTypesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this philanthropy types get default response has a 3xx status code
func (o *PhilanthropyTypesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this philanthropy types get default response has a 4xx status code
func (o *PhilanthropyTypesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this philanthropy types get default response has a 5xx status code
func (o *PhilanthropyTypesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this philanthropy types get default response a status code equal to that given
func (o *PhilanthropyTypesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the philanthropy types get default response
func (o *PhilanthropyTypesGetDefault) Code() int {
	return o._statusCode
}

func (o *PhilanthropyTypesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PhilanthropyTypes/{id}][%d] PhilanthropyTypes_Get default %s", o._statusCode, payload)
}

func (o *PhilanthropyTypesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PhilanthropyTypes/{id}][%d] PhilanthropyTypes_Get default %s", o._statusCode, payload)
}

func (o *PhilanthropyTypesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PhilanthropyTypesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
