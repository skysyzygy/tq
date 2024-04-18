// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// ConstituentProtectionTypesCreateReader is a Reader for the ConstituentProtectionTypesCreate structure.
type ConstituentProtectionTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentProtectionTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentProtectionTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/ConstituentProtectionTypes] ConstituentProtectionTypes_Create", response, response.Code())
	}
}

// NewConstituentProtectionTypesCreateOK creates a ConstituentProtectionTypesCreateOK with default headers values
func NewConstituentProtectionTypesCreateOK() *ConstituentProtectionTypesCreateOK {
	return &ConstituentProtectionTypesCreateOK{}
}

/*
ConstituentProtectionTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentProtectionTypesCreateOK struct {
	Payload *models.ConstituentProtectionType
}

// IsSuccess returns true when this constituent protection types create o k response has a 2xx status code
func (o *ConstituentProtectionTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent protection types create o k response has a 3xx status code
func (o *ConstituentProtectionTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent protection types create o k response has a 4xx status code
func (o *ConstituentProtectionTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent protection types create o k response has a 5xx status code
func (o *ConstituentProtectionTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent protection types create o k response a status code equal to that given
func (o *ConstituentProtectionTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituent protection types create o k response
func (o *ConstituentProtectionTypesCreateOK) Code() int {
	return 200
}

func (o *ConstituentProtectionTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/ConstituentProtectionTypes][%d] constituentProtectionTypesCreateOK  %+v", 200, o.Payload)
}

func (o *ConstituentProtectionTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/ConstituentProtectionTypes][%d] constituentProtectionTypesCreateOK  %+v", 200, o.Payload)
}

func (o *ConstituentProtectionTypesCreateOK) GetPayload() *models.ConstituentProtectionType {
	return o.Payload
}

func (o *ConstituentProtectionTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituentProtectionType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}