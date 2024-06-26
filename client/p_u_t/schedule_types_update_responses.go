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

// ScheduleTypesUpdateReader is a Reader for the ScheduleTypesUpdate structure.
type ScheduleTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ScheduleTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewScheduleTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewScheduleTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewScheduleTypesUpdateOK creates a ScheduleTypesUpdateOK with default headers values
func NewScheduleTypesUpdateOK() *ScheduleTypesUpdateOK {
	return &ScheduleTypesUpdateOK{}
}

/*
ScheduleTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type ScheduleTypesUpdateOK struct {
	Payload *models.ScheduleType
}

// IsSuccess returns true when this schedule types update o k response has a 2xx status code
func (o *ScheduleTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this schedule types update o k response has a 3xx status code
func (o *ScheduleTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this schedule types update o k response has a 4xx status code
func (o *ScheduleTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this schedule types update o k response has a 5xx status code
func (o *ScheduleTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this schedule types update o k response a status code equal to that given
func (o *ScheduleTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the schedule types update o k response
func (o *ScheduleTypesUpdateOK) Code() int {
	return 200
}

func (o *ScheduleTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ScheduleTypes/{id}][%d] scheduleTypesUpdateOK %s", 200, payload)
}

func (o *ScheduleTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ScheduleTypes/{id}][%d] scheduleTypesUpdateOK %s", 200, payload)
}

func (o *ScheduleTypesUpdateOK) GetPayload() *models.ScheduleType {
	return o.Payload
}

func (o *ScheduleTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ScheduleType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewScheduleTypesUpdateDefault creates a ScheduleTypesUpdateDefault with default headers values
func NewScheduleTypesUpdateDefault(code int) *ScheduleTypesUpdateDefault {
	return &ScheduleTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
ScheduleTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type ScheduleTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this schedule types update default response has a 2xx status code
func (o *ScheduleTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this schedule types update default response has a 3xx status code
func (o *ScheduleTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this schedule types update default response has a 4xx status code
func (o *ScheduleTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this schedule types update default response has a 5xx status code
func (o *ScheduleTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this schedule types update default response a status code equal to that given
func (o *ScheduleTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the schedule types update default response
func (o *ScheduleTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *ScheduleTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ScheduleTypes/{id}][%d] ScheduleTypes_Update default %s", o._statusCode, payload)
}

func (o *ScheduleTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/ScheduleTypes/{id}][%d] ScheduleTypes_Update default %s", o._statusCode, payload)
}

func (o *ScheduleTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ScheduleTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
