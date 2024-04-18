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

// ElectronicAddressTypesGetAllReader is a Reader for the ElectronicAddressTypesGetAll structure.
type ElectronicAddressTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ElectronicAddressTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewElectronicAddressTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ElectronicAddressTypes] ElectronicAddressTypes_GetAll", response, response.Code())
	}
}

// NewElectronicAddressTypesGetAllOK creates a ElectronicAddressTypesGetAllOK with default headers values
func NewElectronicAddressTypesGetAllOK() *ElectronicAddressTypesGetAllOK {
	return &ElectronicAddressTypesGetAllOK{}
}

/*
ElectronicAddressTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ElectronicAddressTypesGetAllOK struct {
	Payload []*models.ElectronicAddressType
}

// IsSuccess returns true when this electronic address types get all o k response has a 2xx status code
func (o *ElectronicAddressTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this electronic address types get all o k response has a 3xx status code
func (o *ElectronicAddressTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this electronic address types get all o k response has a 4xx status code
func (o *ElectronicAddressTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this electronic address types get all o k response has a 5xx status code
func (o *ElectronicAddressTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this electronic address types get all o k response a status code equal to that given
func (o *ElectronicAddressTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the electronic address types get all o k response
func (o *ElectronicAddressTypesGetAllOK) Code() int {
	return 200
}

func (o *ElectronicAddressTypesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ElectronicAddressTypes][%d] electronicAddressTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *ElectronicAddressTypesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ElectronicAddressTypes][%d] electronicAddressTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *ElectronicAddressTypesGetAllOK) GetPayload() []*models.ElectronicAddressType {
	return o.Payload
}

func (o *ElectronicAddressTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}