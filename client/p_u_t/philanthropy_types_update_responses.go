// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// PhilanthropyTypesUpdateReader is a Reader for the PhilanthropyTypesUpdate structure.
type PhilanthropyTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhilanthropyTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhilanthropyTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/PhilanthropyTypes/{id}] PhilanthropyTypes_Update", response, response.Code())
	}
}

// NewPhilanthropyTypesUpdateOK creates a PhilanthropyTypesUpdateOK with default headers values
func NewPhilanthropyTypesUpdateOK() *PhilanthropyTypesUpdateOK {
	return &PhilanthropyTypesUpdateOK{}
}

/*
PhilanthropyTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PhilanthropyTypesUpdateOK struct {
	Payload *models.PhilanthropyType
}

// IsSuccess returns true when this philanthropy types update o k response has a 2xx status code
func (o *PhilanthropyTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this philanthropy types update o k response has a 3xx status code
func (o *PhilanthropyTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this philanthropy types update o k response has a 4xx status code
func (o *PhilanthropyTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this philanthropy types update o k response has a 5xx status code
func (o *PhilanthropyTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this philanthropy types update o k response a status code equal to that given
func (o *PhilanthropyTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the philanthropy types update o k response
func (o *PhilanthropyTypesUpdateOK) Code() int {
	return 200
}

func (o *PhilanthropyTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/PhilanthropyTypes/{id}][%d] philanthropyTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *PhilanthropyTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/PhilanthropyTypes/{id}][%d] philanthropyTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *PhilanthropyTypesUpdateOK) GetPayload() *models.PhilanthropyType {
	return o.Payload
}

func (o *PhilanthropyTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhilanthropyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}