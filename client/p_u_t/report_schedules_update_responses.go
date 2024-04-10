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

// ReportSchedulesUpdateReader is a Reader for the ReportSchedulesUpdate structure.
type ReportSchedulesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportSchedulesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportSchedulesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /Reporting/ReportSchedules/{id}] ReportSchedules_Update", response, response.Code())
	}
}

// NewReportSchedulesUpdateOK creates a ReportSchedulesUpdateOK with default headers values
func NewReportSchedulesUpdateOK() *ReportSchedulesUpdateOK {
	return &ReportSchedulesUpdateOK{}
}

/*
ReportSchedulesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ReportSchedulesUpdateOK struct {
	Payload *models.ReportSchedule
}

// IsSuccess returns true when this report schedules update o k response has a 2xx status code
func (o *ReportSchedulesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report schedules update o k response has a 3xx status code
func (o *ReportSchedulesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report schedules update o k response has a 4xx status code
func (o *ReportSchedulesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report schedules update o k response has a 5xx status code
func (o *ReportSchedulesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report schedules update o k response a status code equal to that given
func (o *ReportSchedulesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the report schedules update o k response
func (o *ReportSchedulesUpdateOK) Code() int {
	return 200
}

func (o *ReportSchedulesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /Reporting/ReportSchedules/{id}][%d] reportSchedulesUpdateOK  %+v", 200, o.Payload)
}

func (o *ReportSchedulesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /Reporting/ReportSchedules/{id}][%d] reportSchedulesUpdateOK  %+v", 200, o.Payload)
}

func (o *ReportSchedulesUpdateOK) GetPayload() *models.ReportSchedule {
	return o.Payload
}

func (o *ReportSchedulesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportSchedule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
