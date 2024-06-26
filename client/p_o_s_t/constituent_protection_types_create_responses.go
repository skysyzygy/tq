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

// ConstituentProtectionTypesCreateReader is a Reader for the ConstituentProtectionTypesCreate structure.
type ConstituentProtectionTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentProtectionTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentProtectionTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewConstituentProtectionTypesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewConstituentProtectionTypesCreateOK creates a ConstituentProtectionTypesCreateOK with default headers values
func NewConstituentProtectionTypesCreateOK() *ConstituentProtectionTypesCreateOK {
	return &ConstituentProtectionTypesCreateOK{}
}

/*
ConstituentProtectionTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentProtectionTypesCreateOK struct {
	Payload *models.ConstituentProtectionType
}

// IsSuccess returns true when this constituent protection types create o k response has a 2xx status code
func (o *ConstituentProtectionTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent protection types create o k response has a 3xx status code
func (o *ConstituentProtectionTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent protection types create o k response has a 4xx status code
func (o *ConstituentProtectionTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent protection types create o k response has a 5xx status code
func (o *ConstituentProtectionTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent protection types create o k response a status code equal to that given
func (o *ConstituentProtectionTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituent protection types create o k response
func (o *ConstituentProtectionTypesCreateOK) Code() int {
	return 200
}

func (o *ConstituentProtectionTypesCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ConstituentProtectionTypes][%d] constituentProtectionTypesCreateOK %s", 200, payload)
}

func (o *ConstituentProtectionTypesCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ConstituentProtectionTypes][%d] constituentProtectionTypesCreateOK %s", 200, payload)
}

func (o *ConstituentProtectionTypesCreateOK) GetPayload() *models.ConstituentProtectionType {
	return o.Payload
}

func (o *ConstituentProtectionTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentProtectionType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewConstituentProtectionTypesCreateDefault creates a ConstituentProtectionTypesCreateDefault with default headers values
func NewConstituentProtectionTypesCreateDefault(code int) *ConstituentProtectionTypesCreateDefault {
	return &ConstituentProtectionTypesCreateDefault{
		_statusCode: code,
	}
}

/*
ConstituentProtectionTypesCreateDefault describes a response with status code -1, with default header values.

Error
*/
type ConstituentProtectionTypesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this constituent protection types create default response has a 2xx status code
func (o *ConstituentProtectionTypesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this constituent protection types create default response has a 3xx status code
func (o *ConstituentProtectionTypesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this constituent protection types create default response has a 4xx status code
func (o *ConstituentProtectionTypesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this constituent protection types create default response has a 5xx status code
func (o *ConstituentProtectionTypesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this constituent protection types create default response a status code equal to that given
func (o *ConstituentProtectionTypesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the constituent protection types create default response
func (o *ConstituentProtectionTypesCreateDefault) Code() int {
	return o._statusCode
}

func (o *ConstituentProtectionTypesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ConstituentProtectionTypes][%d] ConstituentProtectionTypes_Create default %s", o._statusCode, payload)
}

func (o *ConstituentProtectionTypesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ConstituentProtectionTypes][%d] ConstituentProtectionTypes_Create default %s", o._statusCode, payload)
}

func (o *ConstituentProtectionTypesCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ConstituentProtectionTypesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
