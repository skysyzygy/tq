// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// ReportRequestsGetAllReader is a Reader for the ReportRequestsGetAll structure.
type ReportRequestsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportRequestsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportRequestsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReportRequestsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReportRequestsGetAllOK creates a ReportRequestsGetAllOK with default headers values
func NewReportRequestsGetAllOK() *ReportRequestsGetAllOK {
	return &ReportRequestsGetAllOK{}
}

/*
ReportRequestsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type ReportRequestsGetAllOK struct {
	Payload []*models.ReportRequest
}

// IsSuccess returns true when this report requests get all o k response has a 2xx status code
func (o *ReportRequestsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report requests get all o k response has a 3xx status code
func (o *ReportRequestsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report requests get all o k response has a 4xx status code
func (o *ReportRequestsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report requests get all o k response has a 5xx status code
func (o *ReportRequestsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report requests get all o k response a status code equal to that given
func (o *ReportRequestsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the report requests get all o k response
func (o *ReportRequestsGetAllOK) Code() int {
	return 200
}

func (o *ReportRequestsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests][%d] reportRequestsGetAllOK %s", 200, payload)
}

func (o *ReportRequestsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests][%d] reportRequestsGetAllOK %s", 200, payload)
}

func (o *ReportRequestsGetAllOK) GetPayload() []*models.ReportRequest {
	return o.Payload
}

func (o *ReportRequestsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportRequestsGetAllDefault creates a ReportRequestsGetAllDefault with default headers values
func NewReportRequestsGetAllDefault(code int) *ReportRequestsGetAllDefault {
	return &ReportRequestsGetAllDefault{
		_statusCode: code,
	}
}

/*
ReportRequestsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type ReportRequestsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this report requests get all default response has a 2xx status code
func (o *ReportRequestsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this report requests get all default response has a 3xx status code
func (o *ReportRequestsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this report requests get all default response has a 4xx status code
func (o *ReportRequestsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this report requests get all default response has a 5xx status code
func (o *ReportRequestsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this report requests get all default response a status code equal to that given
func (o *ReportRequestsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the report requests get all default response
func (o *ReportRequestsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *ReportRequestsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests][%d] ReportRequests_GetAll default %s", o._statusCode, payload)
}

func (o *ReportRequestsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests][%d] ReportRequests_GetAll default %s", o._statusCode, payload)
}

func (o *ReportRequestsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReportRequestsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
