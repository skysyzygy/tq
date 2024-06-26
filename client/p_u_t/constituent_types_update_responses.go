// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// ConstituentTypesUpdateReader is a Reader for the ConstituentTypesUpdate structure.
type ConstituentTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentTypesUpdateOK creates a ConstituentTypesUpdateOK with default headers values
func NewConstituentTypesUpdateOK() *ConstituentTypesUpdateOK {
	return &ConstituentTypesUpdateOK{}
}

/*
ConstituentTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentTypesUpdateOK struct {
	Payload *models.ConstituentType
}

// IsSuccess returns true when this constituent types update o k response has a 2xx status code
func (o *ConstituentTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent types update o k response has a 3xx status code
func (o *ConstituentTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent types update o k response has a 4xx status code
func (o *ConstituentTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent types update o k response has a 5xx status code
func (o *ConstituentTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent types update o k response a status code equal to that given
func (o *ConstituentTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituent types update o k response
func (o *ConstituentTypesUpdateOK) Code() int {
	return 200
}

func (o *ConstituentTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ConstituentTypes/{id}][%d] constituentTypesUpdateOK %s", 200, payload)
}

func (o *ConstituentTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ConstituentTypes/{id}][%d] constituentTypesUpdateOK %s", 200, payload)
}

func (o *ConstituentTypesUpdateOK) GetPayload() *models.ConstituentType {
	return o.Payload
}

func (o *ConstituentTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConstituentTypesUpdateDefault creates a ConstituentTypesUpdateDefault with default headers values
func NewConstituentTypesUpdateDefault(code int) *ConstituentTypesUpdateDefault {
	return &ConstituentTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
ConstituentTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituent types update default response has a 2xx status code
func (o *ConstituentTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituent types update default response has a 3xx status code
func (o *ConstituentTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituent types update default response has a 4xx status code
func (o *ConstituentTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituent types update default response has a 5xx status code
func (o *ConstituentTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituent types update default response a status code equal to that given
func (o *ConstituentTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituent types update default response
func (o *ConstituentTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ConstituentTypes/{id}][%d] ConstituentTypes_Update default %s", o._statusCode, payload)
}

func (o *ConstituentTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ConstituentTypes/{id}][%d] ConstituentTypes_Update default %s", o._statusCode, payload)
}

func (o *ConstituentTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
