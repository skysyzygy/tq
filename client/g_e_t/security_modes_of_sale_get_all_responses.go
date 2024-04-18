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

// SecurityModesOfSaleGetAllReader is a Reader for the SecurityModesOfSaleGetAll structure.
type SecurityModesOfSaleGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityModesOfSaleGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityModesOfSaleGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Security/ModesOfSale] SecurityModesOfSale_GetAll", response, response.Code())
	}
}

// NewSecurityModesOfSaleGetAllOK creates a SecurityModesOfSaleGetAllOK with default headers values
func NewSecurityModesOfSaleGetAllOK() *SecurityModesOfSaleGetAllOK {
	return &SecurityModesOfSaleGetAllOK{}
}

/*
SecurityModesOfSaleGetAllOK describes a response with status code 200, with default header values.

OK
*/
type SecurityModesOfSaleGetAllOK struct {
	Payload []*models.ModeOfSaleUserGroup
}

// IsSuccess returns true when this security modes of sale get all o k response has a 2xx status code
func (o *SecurityModesOfSaleGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security modes of sale get all o k response has a 3xx status code
func (o *SecurityModesOfSaleGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security modes of sale get all o k response has a 4xx status code
func (o *SecurityModesOfSaleGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security modes of sale get all o k response has a 5xx status code
func (o *SecurityModesOfSaleGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security modes of sale get all o k response a status code equal to that given
func (o *SecurityModesOfSaleGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the security modes of sale get all o k response
func (o *SecurityModesOfSaleGetAllOK) Code() int {
	return 200
}

func (o *SecurityModesOfSaleGetAllOK) Error() string {
	return fmt.Sprintf("[GET /Security/ModesOfSale][%d] securityModesOfSaleGetAllOK  %+v", 200, o.Payload)
}

func (o *SecurityModesOfSaleGetAllOK) String() string {
	return fmt.Sprintf("[GET /Security/ModesOfSale][%d] securityModesOfSaleGetAllOK  %+v", 200, o.Payload)
}

func (o *SecurityModesOfSaleGetAllOK) GetPayload() []*models.ModeOfSaleUserGroup {
	return o.Payload
}

func (o *SecurityModesOfSaleGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}