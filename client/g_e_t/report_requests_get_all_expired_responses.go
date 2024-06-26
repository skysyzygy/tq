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

// ReportRequestsGetAllExpiredReader is a Reader for the ReportRequestsGetAllExpired structure.
type ReportRequestsGetAllExpiredReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportRequestsGetAllExpiredReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportRequestsGetAllExpiredOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReportRequestsGetAllExpiredDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReportRequestsGetAllExpiredOK creates a ReportRequestsGetAllExpiredOK with default headers values
func NewReportRequestsGetAllExpiredOK() *ReportRequestsGetAllExpiredOK {
	return &ReportRequestsGetAllExpiredOK{}
}

/*
ReportRequestsGetAllExpiredOK describes a response with status code 200, with default header values.

OK
*/
type ReportRequestsGetAllExpiredOK struct {
	Payload []*models.ExpiredReportRequest
}

// IsSuccess returns true when this report requests get all expired o k response has a 2xx status code
func (o *ReportRequestsGetAllExpiredOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this report requests get all expired o k response has a 3xx status code
func (o *ReportRequestsGetAllExpiredOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this report requests get all expired o k response has a 4xx status code
func (o *ReportRequestsGetAllExpiredOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this report requests get all expired o k response has a 5xx status code
func (o *ReportRequestsGetAllExpiredOK) IsServerError() bool {
	return false
}

// IsCode returns true when this report requests get all expired o k response a status code equal to that given
func (o *ReportRequestsGetAllExpiredOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the report requests get all expired o k response
func (o *ReportRequestsGetAllExpiredOK) Code() int {
	return 200
}

func (o *ReportRequestsGetAllExpiredOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests/Expired][%d] reportRequestsGetAllExpiredOK %s", 200, payload)
}

func (o *ReportRequestsGetAllExpiredOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests/Expired][%d] reportRequestsGetAllExpiredOK %s", 200, payload)
}

func (o *ReportRequestsGetAllExpiredOK) GetPayload() []*models.ExpiredReportRequest {
	return o.Payload
}

func (o *ReportRequestsGetAllExpiredOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportRequestsGetAllExpiredDefault creates a ReportRequestsGetAllExpiredDefault with default headers values
func NewReportRequestsGetAllExpiredDefault(code int) *ReportRequestsGetAllExpiredDefault {
	return &ReportRequestsGetAllExpiredDefault{
		_statusCode: code,
	}
}

/*
ReportRequestsGetAllExpiredDefault describes a response with status code -1, with default header values.

Error
*/
type ReportRequestsGetAllExpiredDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this report requests get all expired default response has a 2xx status code
func (o *ReportRequestsGetAllExpiredDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this report requests get all expired default response has a 3xx status code
func (o *ReportRequestsGetAllExpiredDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this report requests get all expired default response has a 4xx status code
func (o *ReportRequestsGetAllExpiredDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this report requests get all expired default response has a 5xx status code
func (o *ReportRequestsGetAllExpiredDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this report requests get all expired default response a status code equal to that given
func (o *ReportRequestsGetAllExpiredDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the report requests get all expired default response
func (o *ReportRequestsGetAllExpiredDefault) Code() int {
	return o._statusCode
}

func (o *ReportRequestsGetAllExpiredDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests/Expired][%d] ReportRequests_GetAllExpired default %s", o._statusCode, payload)
}

func (o *ReportRequestsGetAllExpiredDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/ReportRequests/Expired][%d] ReportRequests_GetAllExpired default %s", o._statusCode, payload)
}

func (o *ReportRequestsGetAllExpiredDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReportRequestsGetAllExpiredDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
