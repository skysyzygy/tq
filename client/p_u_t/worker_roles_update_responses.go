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

// WorkerRolesUpdateReader is a Reader for the WorkerRolesUpdate structure.
type WorkerRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WorkerRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWorkerRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWorkerRolesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWorkerRolesUpdateOK creates a WorkerRolesUpdateOK with default headers values
func NewWorkerRolesUpdateOK() *WorkerRolesUpdateOK {
	return &WorkerRolesUpdateOK{}
}

/*
WorkerRolesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type WorkerRolesUpdateOK struct {
	Payload *models.WorkerRole
}

// IsSuccess returns true when this worker roles update o k response has a 2xx status code
func (o *WorkerRolesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this worker roles update o k response has a 3xx status code
func (o *WorkerRolesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this worker roles update o k response has a 4xx status code
func (o *WorkerRolesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this worker roles update o k response has a 5xx status code
func (o *WorkerRolesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this worker roles update o k response a status code equal to that given
func (o *WorkerRolesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the worker roles update o k response
func (o *WorkerRolesUpdateOK) Code() int {
	return 200
}

func (o *WorkerRolesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/WorkerRoles/{id}][%d] workerRolesUpdateOK %s", 200, payload)
}

func (o *WorkerRolesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/WorkerRoles/{id}][%d] workerRolesUpdateOK %s", 200, payload)
}

func (o *WorkerRolesUpdateOK) GetPayload() *models.WorkerRole {
	return o.Payload
}

func (o *WorkerRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WorkerRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWorkerRolesUpdateDefault creates a WorkerRolesUpdateDefault with default headers values
func NewWorkerRolesUpdateDefault(code int) *WorkerRolesUpdateDefault {
	return &WorkerRolesUpdateDefault{
		_statusCode: code,
	}
}

/*
WorkerRolesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type WorkerRolesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this worker roles update default response has a 2xx status code
func (o *WorkerRolesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this worker roles update default response has a 3xx status code
func (o *WorkerRolesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this worker roles update default response has a 4xx status code
func (o *WorkerRolesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this worker roles update default response has a 5xx status code
func (o *WorkerRolesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this worker roles update default response a status code equal to that given
func (o *WorkerRolesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the worker roles update default response
func (o *WorkerRolesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *WorkerRolesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/WorkerRoles/{id}][%d] WorkerRoles_Update default %s", o._statusCode, payload)
}

func (o *WorkerRolesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/WorkerRoles/{id}][%d] WorkerRoles_Update default %s", o._statusCode, payload)
}

func (o *WorkerRolesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *WorkerRolesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
