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

// ConstituentProtectionTypesGetReader is a Reader for the ConstituentProtectionTypesGet structure.
type ConstituentProtectionTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentProtectionTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentProtectionTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ConstituentProtectionTypes/{id}] ConstituentProtectionTypes_Get", response, response.Code())
	}
}

// NewConstituentProtectionTypesGetOK creates a ConstituentProtectionTypesGetOK with default headers values
func NewConstituentProtectionTypesGetOK() *ConstituentProtectionTypesGetOK {
	return &ConstituentProtectionTypesGetOK{}
}

/*
ConstituentProtectionTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentProtectionTypesGetOK struct {
	Payload *models.ConstituentProtectionType
}

// IsSuccess returns true when this constituent protection types get o k response has a 2xx status code
func (o *ConstituentProtectionTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent protection types get o k response has a 3xx status code
func (o *ConstituentProtectionTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent protection types get o k response has a 4xx status code
func (o *ConstituentProtectionTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent protection types get o k response has a 5xx status code
func (o *ConstituentProtectionTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent protection types get o k response a status code equal to that given
func (o *ConstituentProtectionTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituent protection types get o k response
func (o *ConstituentProtectionTypesGetOK) Code() int {
	return 200
}

func (o *ConstituentProtectionTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ConstituentProtectionTypes/{id}][%d] constituentProtectionTypesGetOK  %+v", 200, o.Payload)
}

func (o *ConstituentProtectionTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ConstituentProtectionTypes/{id}][%d] constituentProtectionTypesGetOK  %+v", 200, o.Payload)
}

func (o *ConstituentProtectionTypesGetOK) GetPayload() *models.ConstituentProtectionType {
	return o.Payload
}

func (o *ConstituentProtectionTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentProtectionType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
