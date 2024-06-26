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

// ContactPermissionsForTransactionReader is a Reader for the ContactPermissionsForTransaction structure.
type ContactPermissionsForTransactionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPermissionsForTransactionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPermissionsForTransactionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContactPermissionsForTransactionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContactPermissionsForTransactionOK creates a ContactPermissionsForTransactionOK with default headers values
func NewContactPermissionsForTransactionOK() *ContactPermissionsForTransactionOK {
	return &ContactPermissionsForTransactionOK{}
}

/*
ContactPermissionsForTransactionOK describes a response with status code 200, with default header values.

OK
*/
type ContactPermissionsForTransactionOK struct {
	Payload []*models.ContactPermission
}

// IsSuccess returns true when this contact permissions for transaction o k response has a 2xx status code
func (o *ContactPermissionsForTransactionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact permissions for transaction o k response has a 3xx status code
func (o *ContactPermissionsForTransactionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact permissions for transaction o k response has a 4xx status code
func (o *ContactPermissionsForTransactionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact permissions for transaction o k response has a 5xx status code
func (o *ContactPermissionsForTransactionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact permissions for transaction o k response a status code equal to that given
func (o *ContactPermissionsForTransactionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact permissions for transaction o k response
func (o *ContactPermissionsForTransactionOK) Code() int {
	return 200
}

func (o *ContactPermissionsForTransactionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/ContactPermissions/ForTransaction][%d] contactPermissionsForTransactionOK %s", 200, payload)
}

func (o *ContactPermissionsForTransactionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/ContactPermissions/ForTransaction][%d] contactPermissionsForTransactionOK %s", 200, payload)
}

func (o *ContactPermissionsForTransactionOK) GetPayload() []*models.ContactPermission {
	return o.Payload
}

func (o *ContactPermissionsForTransactionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContactPermissionsForTransactionDefault creates a ContactPermissionsForTransactionDefault with default headers values
func NewContactPermissionsForTransactionDefault(code int) *ContactPermissionsForTransactionDefault {
	return &ContactPermissionsForTransactionDefault{
		_statusCode: code,
	}
}

/*
ContactPermissionsForTransactionDefault describes a response with status code -1, with default header values.

Error
*/
type ContactPermissionsForTransactionDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this contact permissions for transaction default response has a 2xx status code
func (o *ContactPermissionsForTransactionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contact permissions for transaction default response has a 3xx status code
func (o *ContactPermissionsForTransactionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contact permissions for transaction default response has a 4xx status code
func (o *ContactPermissionsForTransactionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contact permissions for transaction default response has a 5xx status code
func (o *ContactPermissionsForTransactionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contact permissions for transaction default response a status code equal to that given
func (o *ContactPermissionsForTransactionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contact permissions for transaction default response
func (o *ContactPermissionsForTransactionDefault) Code() int {
	return o._statusCode
}

func (o *ContactPermissionsForTransactionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/ContactPermissions/ForTransaction][%d] ContactPermissions_ForTransaction default %s", o._statusCode, payload)
}

func (o *ContactPermissionsForTransactionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /CRM/ContactPermissions/ForTransaction][%d] ContactPermissions_ForTransaction default %s", o._statusCode, payload)
}

func (o *ContactPermissionsForTransactionDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ContactPermissionsForTransactionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
