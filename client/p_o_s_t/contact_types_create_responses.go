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

// ContactTypesCreateReader is a Reader for the ContactTypesCreate structure.
type ContactTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContactTypesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContactTypesCreateOK creates a ContactTypesCreateOK with default headers values
func NewContactTypesCreateOK() *ContactTypesCreateOK {
	return &ContactTypesCreateOK{}
}

/*
ContactTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type ContactTypesCreateOK struct {
	Payload *models.ContactType
}

// IsSuccess returns true when this contact types create o k response has a 2xx status code
func (o *ContactTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact types create o k response has a 3xx status code
func (o *ContactTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact types create o k response has a 4xx status code
func (o *ContactTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact types create o k response has a 5xx status code
func (o *ContactTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact types create o k response a status code equal to that given
func (o *ContactTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact types create o k response
func (o *ContactTypesCreateOK) Code() int {
	return 200
}

func (o *ContactTypesCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContactTypes][%d] contactTypesCreateOK %s", 200, payload)
}

func (o *ContactTypesCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContactTypes][%d] contactTypesCreateOK %s", 200, payload)
}

func (o *ContactTypesCreateOK) GetPayload() *models.ContactType {
	return o.Payload
}

func (o *ContactTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContactTypesCreateDefault creates a ContactTypesCreateDefault with default headers values
func NewContactTypesCreateDefault(code int) *ContactTypesCreateDefault {
	return &ContactTypesCreateDefault{
		_statusCode: code,
	}
}

/*
ContactTypesCreateDefault describes a response with status code -1, with default header values.

Error
*/
type ContactTypesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this contact types create default response has a 2xx status code
func (o *ContactTypesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contact types create default response has a 3xx status code
func (o *ContactTypesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contact types create default response has a 4xx status code
func (o *ContactTypesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contact types create default response has a 5xx status code
func (o *ContactTypesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contact types create default response a status code equal to that given
func (o *ContactTypesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contact types create default response
func (o *ContactTypesCreateDefault) Code() int {
	return o._statusCode
}

func (o *ContactTypesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContactTypes][%d] ContactTypes_Create default %s", o._statusCode, payload)
}

func (o *ContactTypesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/ContactTypes][%d] ContactTypes_Create default %s", o._statusCode, payload)
}

func (o *ContactTypesCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ContactTypesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
