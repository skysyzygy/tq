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

// ElectronicAddressesUpdateReader is a Reader for the ElectronicAddressesUpdate structure.
type ElectronicAddressesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ElectronicAddressesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewElectronicAddressesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewElectronicAddressesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewElectronicAddressesUpdateOK creates a ElectronicAddressesUpdateOK with default headers values
func NewElectronicAddressesUpdateOK() *ElectronicAddressesUpdateOK {
	return &ElectronicAddressesUpdateOK{}
}

/*
ElectronicAddressesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ElectronicAddressesUpdateOK struct {
	Payload *models.ElectronicAddress
}

// IsSuccess returns true when this electronic addresses update o k response has a 2xx status code
func (o *ElectronicAddressesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this electronic addresses update o k response has a 3xx status code
func (o *ElectronicAddressesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this electronic addresses update o k response has a 4xx status code
func (o *ElectronicAddressesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this electronic addresses update o k response has a 5xx status code
func (o *ElectronicAddressesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this electronic addresses update o k response a status code equal to that given
func (o *ElectronicAddressesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the electronic addresses update o k response
func (o *ElectronicAddressesUpdateOK) Code() int {
	return 200
}

func (o *ElectronicAddressesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/ElectronicAddresses/{electronicAddressId}][%d] electronicAddressesUpdateOK %s", 200, payload)
}

func (o *ElectronicAddressesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/ElectronicAddresses/{electronicAddressId}][%d] electronicAddressesUpdateOK %s", 200, payload)
}

func (o *ElectronicAddressesUpdateOK) GetPayload() *models.ElectronicAddress {
	return o.Payload
}

func (o *ElectronicAddressesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElectronicAddress)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewElectronicAddressesUpdateDefault creates a ElectronicAddressesUpdateDefault with default headers values
func NewElectronicAddressesUpdateDefault(code int) *ElectronicAddressesUpdateDefault {
	return &ElectronicAddressesUpdateDefault{
		_statusCode: code,
	}
}

/*
ElectronicAddressesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ElectronicAddressesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this electronic addresses update default response has a 2xx status code
func (o *ElectronicAddressesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this electronic addresses update default response has a 3xx status code
func (o *ElectronicAddressesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this electronic addresses update default response has a 4xx status code
func (o *ElectronicAddressesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this electronic addresses update default response has a 5xx status code
func (o *ElectronicAddressesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this electronic addresses update default response a status code equal to that given
func (o *ElectronicAddressesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the electronic addresses update default response
func (o *ElectronicAddressesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ElectronicAddressesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/ElectronicAddresses/{electronicAddressId}][%d] ElectronicAddresses_Update default %s", o._statusCode, payload)
}

func (o *ElectronicAddressesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/ElectronicAddresses/{electronicAddressId}][%d] ElectronicAddresses_Update default %s", o._statusCode, payload)
}

func (o *ElectronicAddressesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ElectronicAddressesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
