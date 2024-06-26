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

// AddressesUpdateReader is a Reader for the AddressesUpdate structure.
type AddressesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddressesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddressesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAddressesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddressesUpdateOK creates a AddressesUpdateOK with default headers values
func NewAddressesUpdateOK() *AddressesUpdateOK {
	return &AddressesUpdateOK{}
}

/*
AddressesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type AddressesUpdateOK struct {
	Payload *models.Address
}

// IsSuccess returns true when this addresses update o k response has a 2xx status code
func (o *AddressesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this addresses update o k response has a 3xx status code
func (o *AddressesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this addresses update o k response has a 4xx status code
func (o *AddressesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this addresses update o k response has a 5xx status code
func (o *AddressesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this addresses update o k response a status code equal to that given
func (o *AddressesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the addresses update o k response
func (o *AddressesUpdateOK) Code() int {
	return 200
}

func (o *AddressesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/Addresses/{addressId}][%d] addressesUpdateOK %s", 200, payload)
}

func (o *AddressesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/Addresses/{addressId}][%d] addressesUpdateOK %s", 200, payload)
}

func (o *AddressesUpdateOK) GetPayload() *models.Address {
	return o.Payload
}

func (o *AddressesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Address)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddressesUpdateDefault creates a AddressesUpdateDefault with default headers values
func NewAddressesUpdateDefault(code int) *AddressesUpdateDefault {
	return &AddressesUpdateDefault{
		_statusCode: code,
	}
}

/*
AddressesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type AddressesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this addresses update default response has a 2xx status code
func (o *AddressesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this addresses update default response has a 3xx status code
func (o *AddressesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this addresses update default response has a 4xx status code
func (o *AddressesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this addresses update default response has a 5xx status code
func (o *AddressesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this addresses update default response a status code equal to that given
func (o *AddressesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the addresses update default response
func (o *AddressesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *AddressesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/Addresses/{addressId}][%d] Addresses_Update default %s", o._statusCode, payload)
}

func (o *AddressesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/Addresses/{addressId}][%d] Addresses_Update default %s", o._statusCode, payload)
}

func (o *AddressesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AddressesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
