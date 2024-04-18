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

// ScheduleTypesGetReader is a Reader for the ScheduleTypesGet structure.
type ScheduleTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ScheduleTypes/{id}] ScheduleTypes_Get", response, response.Code())
	}
}

// NewScheduleTypesGetOK creates a ScheduleTypesGetOK with default headers values
func NewScheduleTypesGetOK() *ScheduleTypesGetOK {
	return &ScheduleTypesGetOK{}
}

/*
ScheduleTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ScheduleTypesGetOK struct {
	Payload *models.ScheduleType
}

// IsSuccess returns true when this schedule types get o k response has a 2xx status code
func (o *ScheduleTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule types get o k response has a 3xx status code
func (o *ScheduleTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule types get o k response has a 4xx status code
func (o *ScheduleTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule types get o k response has a 5xx status code
func (o *ScheduleTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule types get o k response a status code equal to that given
func (o *ScheduleTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schedule types get o k response
func (o *ScheduleTypesGetOK) Code() int {
	return 200
}

func (o *ScheduleTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ScheduleTypes/{id}][%d] scheduleTypesGetOK  %+v", 200, o.Payload)
}

func (o *ScheduleTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ScheduleTypes/{id}][%d] scheduleTypesGetOK  %+v", 200, o.Payload)
}

func (o *ScheduleTypesGetOK) GetPayload() *models.ScheduleType {
	return o.Payload
}

func (o *ScheduleTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}