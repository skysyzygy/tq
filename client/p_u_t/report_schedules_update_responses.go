// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
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
		result := NewReportSchedulesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
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
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/ReportSchedules/{id}][%d] reportSchedulesUpdateOK %s", 200, payload)
}

func (o *ReportSchedulesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/ReportSchedules/{id}][%d] reportSchedulesUpdateOK %s", 200, payload)
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

// NewReportSchedulesUpdateDefault creates a ReportSchedulesUpdateDefault with default headers values
func NewReportSchedulesUpdateDefault(code int) *ReportSchedulesUpdateDefault {
	return &ReportSchedulesUpdateDefault{
		_statusCode: code,
	}
}

/*
ReportSchedulesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ReportSchedulesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this report schedules update default response has a 2xx status code
func (o *ReportSchedulesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this report schedules update default response has a 3xx status code
func (o *ReportSchedulesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this report schedules update default response has a 4xx status code
func (o *ReportSchedulesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this report schedules update default response has a 5xx status code
func (o *ReportSchedulesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this report schedules update default response a status code equal to that given
func (o *ReportSchedulesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the report schedules update default response
func (o *ReportSchedulesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ReportSchedulesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/ReportSchedules/{id}][%d] ReportSchedules_Update default %s", o._statusCode, payload)
}

func (o *ReportSchedulesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Reporting/ReportSchedules/{id}][%d] ReportSchedules_Update default %s", o._statusCode, payload)
}

func (o *ReportSchedulesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReportSchedulesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
