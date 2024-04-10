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

// CrediteeTypesGetReader is a Reader for the CrediteeTypesGet structure.
type CrediteeTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CrediteeTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCrediteeTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/CrediteeTypes/{id}] CrediteeTypes_Get", response, response.Code())
	}
}

// NewCrediteeTypesGetOK creates a CrediteeTypesGetOK with default headers values
func NewCrediteeTypesGetOK() *CrediteeTypesGetOK {
	return &CrediteeTypesGetOK{}
}

/*
CrediteeTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type CrediteeTypesGetOK struct {
	Payload *models.CrediteeType
}

// IsSuccess returns true when this creditee types get o k response has a 2xx status code
func (o *CrediteeTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this creditee types get o k response has a 3xx status code
func (o *CrediteeTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this creditee types get o k response has a 4xx status code
func (o *CrediteeTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this creditee types get o k response has a 5xx status code
func (o *CrediteeTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this creditee types get o k response a status code equal to that given
func (o *CrediteeTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the creditee types get o k response
func (o *CrediteeTypesGetOK) Code() int {
	return 200
}

func (o *CrediteeTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/CrediteeTypes/{id}][%d] crediteeTypesGetOK  %+v", 200, o.Payload)
}

func (o *CrediteeTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/CrediteeTypes/{id}][%d] crediteeTypesGetOK  %+v", 200, o.Payload)
}

func (o *CrediteeTypesGetOK) GetPayload() *models.CrediteeType {
	return o.Payload
}

func (o *CrediteeTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CrediteeType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
