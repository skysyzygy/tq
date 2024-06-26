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

// ReportsGetReader is a Reader for the ReportsGet structure.
type ReportsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReportsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReportsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewReportsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewReportsGetOK creates a ReportsGetOK with default headers values
func NewReportsGetOK() *ReportsGetOK {
	return &ReportsGetOK{}
}

/*
ReportsGetOK describes a response with status code 200, with default header values.

OK
*/
type ReportsGetOK struct {
	Payload *models.Report
}

// IsSuccess returns true when this reports get o k response has a 2xx status code
func (o *ReportsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reports get o k response has a 3xx status code
func (o *ReportsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reports get o k response has a 4xx status code
func (o *ReportsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reports get o k response has a 5xx status code
func (o *ReportsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reports get o k response a status code equal to that given
func (o *ReportsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reports get o k response
func (o *ReportsGetOK) Code() int {
	return 200
}

func (o *ReportsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/Reports/{reportId}][%d] reportsGetOK %s", 200, payload)
}

func (o *ReportsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/Reports/{reportId}][%d] reportsGetOK %s", 200, payload)
}

func (o *ReportsGetOK) GetPayload() *models.Report {
	return o.Payload
}

func (o *ReportsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Report)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReportsGetDefault creates a ReportsGetDefault with default headers values
func NewReportsGetDefault(code int) *ReportsGetDefault {
	return &ReportsGetDefault{
		_statusCode: code,
	}
}

/*
ReportsGetDefault describes a response with status code -1, with default header values.

Error
*/
type ReportsGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this reports get default response has a 2xx status code
func (o *ReportsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this reports get default response has a 3xx status code
func (o *ReportsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this reports get default response has a 4xx status code
func (o *ReportsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this reports get default response has a 5xx status code
func (o *ReportsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this reports get default response a status code equal to that given
func (o *ReportsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the reports get default response
func (o *ReportsGetDefault) Code() int {
	return o._statusCode
}

func (o *ReportsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/Reports/{reportId}][%d] Reports_Get default %s", o._statusCode, payload)
}

func (o *ReportsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Reporting/Reports/{reportId}][%d] Reports_Get default %s", o._statusCode, payload)
}

func (o *ReportsGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ReportsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
