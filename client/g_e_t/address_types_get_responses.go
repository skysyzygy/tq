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

// AddressTypesGetReader is a Reader for the AddressTypesGet structure.
type AddressTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddressTypesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddressTypesGetOK creates a AddressTypesGetOK with default headers values
func NewAddressTypesGetOK() *AddressTypesGetOK {
	return &AddressTypesGetOK{}
}

/*
AddressTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type AddressTypesGetOK struct {
	Payload *models.AddressType
}

// IsSuccess returns true when this address types get o k response has a 2xx status code
func (o *AddressTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this address types get o k response has a 3xx status code
func (o *AddressTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this address types get o k response has a 4xx status code
func (o *AddressTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this address types get o k response has a 5xx status code
func (o *AddressTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this address types get o k response a status code equal to that given
func (o *AddressTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the address types get o k response
func (o *AddressTypesGetOK) Code() int {
	return 200
}

func (o *AddressTypesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AddressTypes/{id}][%d] addressTypesGetOK %s", 200, payload)
}

func (o *AddressTypesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AddressTypes/{id}][%d] addressTypesGetOK %s", 200, payload)
}

func (o *AddressTypesGetOK) GetPayload() *models.AddressType {
	return o.Payload
}

func (o *AddressTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressTypesGetDefault creates a AddressTypesGetDefault with default headers values
func NewAddressTypesGetDefault(code int) *AddressTypesGetDefault {
	return &AddressTypesGetDefault{
		_statusCode: code,
	}
}

/*
AddressTypesGetDefault describes a response with status code -1, with default header values.

Error
*/
type AddressTypesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this address types get default response has a 2xx status code
func (o *AddressTypesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this address types get default response has a 3xx status code
func (o *AddressTypesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this address types get default response has a 4xx status code
func (o *AddressTypesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this address types get default response has a 5xx status code
func (o *AddressTypesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this address types get default response a status code equal to that given
func (o *AddressTypesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the address types get default response
func (o *AddressTypesGetDefault) Code() int {
	return o._statusCode
}

func (o *AddressTypesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AddressTypes/{id}][%d] AddressTypes_Get default %s", o._statusCode, payload)
}

func (o *AddressTypesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/AddressTypes/{id}][%d] AddressTypes_Get default %s", o._statusCode, payload)
}

func (o *AddressTypesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddressTypesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
