// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ElectronicAddressesIsAttachedToLoginReader is a Reader for the ElectronicAddressesIsAttachedToLogin structure.
type ElectronicAddressesIsAttachedToLoginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ElectronicAddressesIsAttachedToLoginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewElectronicAddressesIsAttachedToLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/ElectronicAddresses/{electronicAddressId}/IsAttachedToLogin] ElectronicAddresses_IsAttachedToLogin", response, response.Code())
	}
}

// NewElectronicAddressesIsAttachedToLoginOK creates a ElectronicAddressesIsAttachedToLoginOK with default headers values
func NewElectronicAddressesIsAttachedToLoginOK() *ElectronicAddressesIsAttachedToLoginOK {
	return &ElectronicAddressesIsAttachedToLoginOK{}
}

/*
ElectronicAddressesIsAttachedToLoginOK describes a response with status code 200, with default header values.

OK
*/
type ElectronicAddressesIsAttachedToLoginOK struct {
	Payload *models.AttachedToLogin
}

// IsSuccess returns true when this electronic addresses is attached to login o k response has a 2xx status code
func (o *ElectronicAddressesIsAttachedToLoginOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this electronic addresses is attached to login o k response has a 3xx status code
func (o *ElectronicAddressesIsAttachedToLoginOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this electronic addresses is attached to login o k response has a 4xx status code
func (o *ElectronicAddressesIsAttachedToLoginOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this electronic addresses is attached to login o k response has a 5xx status code
func (o *ElectronicAddressesIsAttachedToLoginOK) IsServerError() bool {
	return false
}

// IsCode returns true when this electronic addresses is attached to login o k response a status code equal to that given
func (o *ElectronicAddressesIsAttachedToLoginOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the electronic addresses is attached to login o k response
func (o *ElectronicAddressesIsAttachedToLoginOK) Code() int {
	return 200
}

func (o *ElectronicAddressesIsAttachedToLoginOK) Error() string {
	return fmt.Sprintf("[GET /CRM/ElectronicAddresses/{electronicAddressId}/IsAttachedToLogin][%d] electronicAddressesIsAttachedToLoginOK  %+v", 200, o.Payload)
}

func (o *ElectronicAddressesIsAttachedToLoginOK) String() string {
	return fmt.Sprintf("[GET /CRM/ElectronicAddresses/{electronicAddressId}/IsAttachedToLogin][%d] electronicAddressesIsAttachedToLoginOK  %+v", 200, o.Payload)
}

func (o *ElectronicAddressesIsAttachedToLoginOK) GetPayload() *models.AttachedToLogin {
	return o.Payload
}

func (o *ElectronicAddressesIsAttachedToLoginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AttachedToLogin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
