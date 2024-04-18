// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SchedulePatternTypesDeleteReader is a Reader for the SchedulePatternTypesDelete structure.
type SchedulePatternTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchedulePatternTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSchedulePatternTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/SchedulePatternTypes/{id}] SchedulePatternTypes_Delete", response, response.Code())
	}
}

// NewSchedulePatternTypesDeleteNoContent creates a SchedulePatternTypesDeleteNoContent with default headers values
func NewSchedulePatternTypesDeleteNoContent() *SchedulePatternTypesDeleteNoContent {
	return &SchedulePatternTypesDeleteNoContent{}
}

/*
SchedulePatternTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type SchedulePatternTypesDeleteNoContent struct {
}

// IsSuccess returns true when this schedule pattern types delete no content response has a 2xx status code
func (o *SchedulePatternTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule pattern types delete no content response has a 3xx status code
func (o *SchedulePatternTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule pattern types delete no content response has a 4xx status code
func (o *SchedulePatternTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule pattern types delete no content response has a 5xx status code
func (o *SchedulePatternTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule pattern types delete no content response a status code equal to that given
func (o *SchedulePatternTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the schedule pattern types delete no content response
func (o *SchedulePatternTypesDeleteNoContent) Code() int {
	return 204
}

func (o *SchedulePatternTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/SchedulePatternTypes/{id}][%d] schedulePatternTypesDeleteNoContent ", 204)
}

func (o *SchedulePatternTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/SchedulePatternTypes/{id}][%d] schedulePatternTypesDeleteNoContent ", 204)
}

func (o *SchedulePatternTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}