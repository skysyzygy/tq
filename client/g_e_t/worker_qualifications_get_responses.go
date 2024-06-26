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

// WorkerQualificationsGetReader is a Reader for the WorkerQualificationsGet structure.
type WorkerQualificationsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkerQualificationsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkerQualificationsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWorkerQualificationsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkerQualificationsGetOK creates a WorkerQualificationsGetOK with default headers values
func NewWorkerQualificationsGetOK() *WorkerQualificationsGetOK {
	return &WorkerQualificationsGetOK{}
}

/*
WorkerQualificationsGetOK describes a response with status code 200, with default header values.

OK
*/
type WorkerQualificationsGetOK struct {
	Payload *models.WorkerQualification
}

// IsSuccess returns true when this worker qualifications get o k response has a 2xx status code
func (o *WorkerQualificationsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this worker qualifications get o k response has a 3xx status code
func (o *WorkerQualificationsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this worker qualifications get o k response has a 4xx status code
func (o *WorkerQualificationsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this worker qualifications get o k response has a 5xx status code
func (o *WorkerQualificationsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this worker qualifications get o k response a status code equal to that given
func (o *WorkerQualificationsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the worker qualifications get o k response
func (o *WorkerQualificationsGetOK) Code() int {
	return 200
}

func (o *WorkerQualificationsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/WorkerQualifications/{workerQualificationId}][%d] workerQualificationsGetOK %s", 200, payload)
}

func (o *WorkerQualificationsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/WorkerQualifications/{workerQualificationId}][%d] workerQualificationsGetOK %s", 200, payload)
}

func (o *WorkerQualificationsGetOK) GetPayload() *models.WorkerQualification {
	return o.Payload
}

func (o *WorkerQualificationsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkerQualification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkerQualificationsGetDefault creates a WorkerQualificationsGetDefault with default headers values
func NewWorkerQualificationsGetDefault(code int) *WorkerQualificationsGetDefault {
	return &WorkerQualificationsGetDefault{
		_statusCode: code,
	}
}

/*
WorkerQualificationsGetDefault describes a response with status code -1, with default header values.

Error
*/
type WorkerQualificationsGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this worker qualifications get default response has a 2xx status code
func (o *WorkerQualificationsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this worker qualifications get default response has a 3xx status code
func (o *WorkerQualificationsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this worker qualifications get default response has a 4xx status code
func (o *WorkerQualificationsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this worker qualifications get default response has a 5xx status code
func (o *WorkerQualificationsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this worker qualifications get default response a status code equal to that given
func (o *WorkerQualificationsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the worker qualifications get default response
func (o *WorkerQualificationsGetDefault) Code() int {
	return o._statusCode
}

func (o *WorkerQualificationsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/WorkerQualifications/{workerQualificationId}][%d] WorkerQualifications_Get default %s", o._statusCode, payload)
}

func (o *WorkerQualificationsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/WorkerQualifications/{workerQualificationId}][%d] WorkerQualifications_Get default %s", o._statusCode, payload)
}

func (o *WorkerQualificationsGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *WorkerQualificationsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
