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

// ElectronicAddressTypesGetReader is a Reader for the ElectronicAddressTypesGet structure.
type ElectronicAddressTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ElectronicAddressTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewElectronicAddressTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ElectronicAddressTypes/{id}] ElectronicAddressTypes_Get", response, response.Code())
	}
}

// NewElectronicAddressTypesGetOK creates a ElectronicAddressTypesGetOK with default headers values
func NewElectronicAddressTypesGetOK() *ElectronicAddressTypesGetOK {
	return &ElectronicAddressTypesGetOK{}
}

/*
ElectronicAddressTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ElectronicAddressTypesGetOK struct {
	Payload *models.ElectronicAddressType
}

// IsSuccess returns true when this electronic address types get o k response has a 2xx status code
func (o *ElectronicAddressTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this electronic address types get o k response has a 3xx status code
func (o *ElectronicAddressTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this electronic address types get o k response has a 4xx status code
func (o *ElectronicAddressTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this electronic address types get o k response has a 5xx status code
func (o *ElectronicAddressTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this electronic address types get o k response a status code equal to that given
func (o *ElectronicAddressTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the electronic address types get o k response
func (o *ElectronicAddressTypesGetOK) Code() int {
	return 200
}

func (o *ElectronicAddressTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ElectronicAddressTypes/{id}][%d] electronicAddressTypesGetOK  %+v", 200, o.Payload)
}

func (o *ElectronicAddressTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ElectronicAddressTypes/{id}][%d] electronicAddressTypesGetOK  %+v", 200, o.Payload)
}

func (o *ElectronicAddressTypesGetOK) GetPayload() *models.ElectronicAddressType {
	return o.Payload
}

func (o *ElectronicAddressTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ElectronicAddressType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
