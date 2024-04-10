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

// PhoneTypesGetReader is a Reader for the PhoneTypesGet structure.
type PhoneTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhoneTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhoneTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PhoneTypes/{id}] PhoneTypes_Get", response, response.Code())
	}
}

// NewPhoneTypesGetOK creates a PhoneTypesGetOK with default headers values
func NewPhoneTypesGetOK() *PhoneTypesGetOK {
	return &PhoneTypesGetOK{}
}

/*
PhoneTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type PhoneTypesGetOK struct {
	Payload *models.PhoneType
}

// IsSuccess returns true when this phone types get o k response has a 2xx status code
func (o *PhoneTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phone types get o k response has a 3xx status code
func (o *PhoneTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phone types get o k response has a 4xx status code
func (o *PhoneTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phone types get o k response has a 5xx status code
func (o *PhoneTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phone types get o k response a status code equal to that given
func (o *PhoneTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phone types get o k response
func (o *PhoneTypesGetOK) Code() int {
	return 200
}

func (o *PhoneTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PhoneTypes/{id}][%d] phoneTypesGetOK  %+v", 200, o.Payload)
}

func (o *PhoneTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PhoneTypes/{id}][%d] phoneTypesGetOK  %+v", 200, o.Payload)
}

func (o *PhoneTypesGetOK) GetPayload() *models.PhoneType {
	return o.Payload
}

func (o *PhoneTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhoneType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
