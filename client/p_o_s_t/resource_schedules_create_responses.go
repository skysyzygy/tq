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

// ResourceSchedulesCreateReader is a Reader for the ResourceSchedulesCreate structure.
type ResourceSchedulesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResourceSchedulesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResourceSchedulesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /EventsManagement/ResourceSchedules] ResourceSchedules_Create", response, response.Code())
	}
}

// NewResourceSchedulesCreateOK creates a ResourceSchedulesCreateOK with default headers values
func NewResourceSchedulesCreateOK() *ResourceSchedulesCreateOK {
	return &ResourceSchedulesCreateOK{}
}

/*
ResourceSchedulesCreateOK describes a response with status code 200, with default header values.

OK
*/
type ResourceSchedulesCreateOK struct {
	Payload *models.ResourceSchedule
}

// IsSuccess returns true when this resource schedules create o k response has a 2xx status code
func (o *ResourceSchedulesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this resource schedules create o k response has a 3xx status code
func (o *ResourceSchedulesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this resource schedules create o k response has a 4xx status code
func (o *ResourceSchedulesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this resource schedules create o k response has a 5xx status code
func (o *ResourceSchedulesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this resource schedules create o k response a status code equal to that given
func (o *ResourceSchedulesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the resource schedules create o k response
func (o *ResourceSchedulesCreateOK) Code() int {
	return 200
}

func (o *ResourceSchedulesCreateOK) Error() string {
	return fmt.Sprintf("[POST /EventsManagement/ResourceSchedules][%d] resourceSchedulesCreateOK  %+v", 200, o.Payload)
}

func (o *ResourceSchedulesCreateOK) String() string {
	return fmt.Sprintf("[POST /EventsManagement/ResourceSchedules][%d] resourceSchedulesCreateOK  %+v", 200, o.Payload)
}

func (o *ResourceSchedulesCreateOK) GetPayload() *models.ResourceSchedule {
	return o.Payload
}

func (o *ResourceSchedulesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
