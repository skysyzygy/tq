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

// SpecialActivityTypesGetReader is a Reader for the SpecialActivityTypesGet structure.
type SpecialActivityTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecialActivityTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpecialActivityTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/SpecialActivityTypes/{id}] SpecialActivityTypes_Get", response, response.Code())
	}
}

// NewSpecialActivityTypesGetOK creates a SpecialActivityTypesGetOK with default headers values
func NewSpecialActivityTypesGetOK() *SpecialActivityTypesGetOK {
	return &SpecialActivityTypesGetOK{}
}

/*
SpecialActivityTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type SpecialActivityTypesGetOK struct {
	Payload *models.SpecialActivityType
}

// IsSuccess returns true when this special activity types get o k response has a 2xx status code
func (o *SpecialActivityTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this special activity types get o k response has a 3xx status code
func (o *SpecialActivityTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this special activity types get o k response has a 4xx status code
func (o *SpecialActivityTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this special activity types get o k response has a 5xx status code
func (o *SpecialActivityTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this special activity types get o k response a status code equal to that given
func (o *SpecialActivityTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the special activity types get o k response
func (o *SpecialActivityTypesGetOK) Code() int {
	return 200
}

func (o *SpecialActivityTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/SpecialActivityTypes/{id}][%d] specialActivityTypesGetOK  %+v", 200, o.Payload)
}

func (o *SpecialActivityTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/SpecialActivityTypes/{id}][%d] specialActivityTypesGetOK  %+v", 200, o.Payload)
}

func (o *SpecialActivityTypesGetOK) GetPayload() *models.SpecialActivityType {
	return o.Payload
}

func (o *SpecialActivityTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpecialActivityType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
