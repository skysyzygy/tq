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

// AccountTypesGetReader is a Reader for the AccountTypesGet structure.
type AccountTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AccountTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAccountTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAccountTypesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAccountTypesGetOK creates a AccountTypesGetOK with default headers values
func NewAccountTypesGetOK() *AccountTypesGetOK {
	return &AccountTypesGetOK{}
}

/*
AccountTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type AccountTypesGetOK struct {
	Payload *models.AccountType
}

// IsSuccess returns true when this account types get o k response has a 2xx status code
func (o *AccountTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this account types get o k response has a 3xx status code
func (o *AccountTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this account types get o k response has a 4xx status code
func (o *AccountTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this account types get o k response has a 5xx status code
func (o *AccountTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this account types get o k response a status code equal to that given
func (o *AccountTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the account types get o k response
func (o *AccountTypesGetOK) Code() int {
	return 200
}

func (o *AccountTypesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AccountTypes/{id}][%d] accountTypesGetOK %s", 200, payload)
}

func (o *AccountTypesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AccountTypes/{id}][%d] accountTypesGetOK %s", 200, payload)
}

func (o *AccountTypesGetOK) GetPayload() *models.AccountType {
	return o.Payload
}

func (o *AccountTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AccountType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAccountTypesGetDefault creates a AccountTypesGetDefault with default headers values
func NewAccountTypesGetDefault(code int) *AccountTypesGetDefault {
	return &AccountTypesGetDefault{
		_statusCode: code,
	}
}

/*
AccountTypesGetDefault describes a response with status code -1, with default header values.

Error
*/
type AccountTypesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this account types get default response has a 2xx status code
func (o *AccountTypesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this account types get default response has a 3xx status code
func (o *AccountTypesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this account types get default response has a 4xx status code
func (o *AccountTypesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this account types get default response has a 5xx status code
func (o *AccountTypesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this account types get default response a status code equal to that given
func (o *AccountTypesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the account types get default response
func (o *AccountTypesGetDefault) Code() int {
	return o._statusCode
}

func (o *AccountTypesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AccountTypes/{id}][%d] AccountTypes_Get default %s", o._statusCode, payload)
}

func (o *AccountTypesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AccountTypes/{id}][%d] AccountTypes_Get default %s", o._statusCode, payload)
}

func (o *AccountTypesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AccountTypesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
