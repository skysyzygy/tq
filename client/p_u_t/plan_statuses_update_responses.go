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

// PlanStatusesUpdateReader is a Reader for the PlanStatusesUpdate structure.
type PlanStatusesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlanStatusesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlanStatusesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPlanStatusesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPlanStatusesUpdateOK creates a PlanStatusesUpdateOK with default headers values
func NewPlanStatusesUpdateOK() *PlanStatusesUpdateOK {
	return &PlanStatusesUpdateOK{}
}

/*
PlanStatusesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PlanStatusesUpdateOK struct {
	Payload *models.PlanStatus
}

// IsSuccess returns true when this plan statuses update o k response has a 2xx status code
func (o *PlanStatusesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plan statuses update o k response has a 3xx status code
func (o *PlanStatusesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plan statuses update o k response has a 4xx status code
func (o *PlanStatusesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plan statuses update o k response has a 5xx status code
func (o *PlanStatusesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plan statuses update o k response a status code equal to that given
func (o *PlanStatusesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plan statuses update o k response
func (o *PlanStatusesUpdateOK) Code() int {
	return 200
}

func (o *PlanStatusesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PlanStatuses/{id}][%d] planStatusesUpdateOK %s", 200, payload)
}

func (o *PlanStatusesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PlanStatuses/{id}][%d] planStatusesUpdateOK %s", 200, payload)
}

func (o *PlanStatusesUpdateOK) GetPayload() *models.PlanStatus {
	return o.Payload
}

func (o *PlanStatusesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PlanStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanStatusesUpdateDefault creates a PlanStatusesUpdateDefault with default headers values
func NewPlanStatusesUpdateDefault(code int) *PlanStatusesUpdateDefault {
	return &PlanStatusesUpdateDefault{
		_statusCode: code,
	}
}

/*
PlanStatusesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PlanStatusesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this plan statuses update default response has a 2xx status code
func (o *PlanStatusesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plan statuses update default response has a 3xx status code
func (o *PlanStatusesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plan statuses update default response has a 4xx status code
func (o *PlanStatusesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plan statuses update default response has a 5xx status code
func (o *PlanStatusesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plan statuses update default response a status code equal to that given
func (o *PlanStatusesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plan statuses update default response
func (o *PlanStatusesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PlanStatusesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PlanStatuses/{id}][%d] PlanStatuses_Update default %s", o._statusCode, payload)
}

func (o *PlanStatusesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PlanStatuses/{id}][%d] PlanStatuses_Update default %s", o._statusCode, payload)
}

func (o *PlanStatusesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PlanStatusesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
