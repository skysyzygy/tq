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

// StepTypesUpdateReader is a Reader for the StepTypesUpdate structure.
type StepTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StepTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStepTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/StepTypes/{id}] StepTypes_Update", response, response.Code())
	}
}

// NewStepTypesUpdateOK creates a StepTypesUpdateOK with default headers values
func NewStepTypesUpdateOK() *StepTypesUpdateOK {
	return &StepTypesUpdateOK{}
}

/*
StepTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type StepTypesUpdateOK struct {
	Payload *models.StepType
}

// IsSuccess returns true when this step types update o k response has a 2xx status code
func (o *StepTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this step types update o k response has a 3xx status code
func (o *StepTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this step types update o k response has a 4xx status code
func (o *StepTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this step types update o k response has a 5xx status code
func (o *StepTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this step types update o k response a status code equal to that given
func (o *StepTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the step types update o k response
func (o *StepTypesUpdateOK) Code() int {
	return 200
}

func (o *StepTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/StepTypes/{id}][%d] stepTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *StepTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/StepTypes/{id}][%d] stepTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *StepTypesUpdateOK) GetPayload() *models.StepType {
	return o.Payload
}

func (o *StepTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.StepType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
