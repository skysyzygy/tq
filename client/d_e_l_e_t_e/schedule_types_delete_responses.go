// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ScheduleTypesDeleteReader is a Reader for the ScheduleTypesDelete structure.
type ScheduleTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewScheduleTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/ScheduleTypes/{id}] ScheduleTypes_Delete", response, response.Code())
	}
}

// NewScheduleTypesDeleteNoContent creates a ScheduleTypesDeleteNoContent with default headers values
func NewScheduleTypesDeleteNoContent() *ScheduleTypesDeleteNoContent {
	return &ScheduleTypesDeleteNoContent{}
}

/*
ScheduleTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ScheduleTypesDeleteNoContent struct {
}

// IsSuccess returns true when this schedule types delete no content response has a 2xx status code
func (o *ScheduleTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule types delete no content response has a 3xx status code
func (o *ScheduleTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule types delete no content response has a 4xx status code
func (o *ScheduleTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule types delete no content response has a 5xx status code
func (o *ScheduleTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule types delete no content response a status code equal to that given
func (o *ScheduleTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the schedule types delete no content response
func (o *ScheduleTypesDeleteNoContent) Code() int {
	return 204
}

func (o *ScheduleTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ScheduleTypes/{id}][%d] scheduleTypesDeleteNoContent ", 204)
}

func (o *ScheduleTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ScheduleTypes/{id}][%d] scheduleTypesDeleteNoContent ", 204)
}

func (o *ScheduleTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}