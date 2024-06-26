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

// PlanPrioritiesGetAllReader is a Reader for the PlanPrioritiesGetAll structure.
type PlanPrioritiesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlanPrioritiesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlanPrioritiesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPlanPrioritiesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPlanPrioritiesGetAllOK creates a PlanPrioritiesGetAllOK with default headers values
func NewPlanPrioritiesGetAllOK() *PlanPrioritiesGetAllOK {
	return &PlanPrioritiesGetAllOK{}
}

/*
PlanPrioritiesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PlanPrioritiesGetAllOK struct {
	Payload []*models.PlanPriority
}

// IsSuccess returns true when this plan priorities get all o k response has a 2xx status code
func (o *PlanPrioritiesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plan priorities get all o k response has a 3xx status code
func (o *PlanPrioritiesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plan priorities get all o k response has a 4xx status code
func (o *PlanPrioritiesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plan priorities get all o k response has a 5xx status code
func (o *PlanPrioritiesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plan priorities get all o k response a status code equal to that given
func (o *PlanPrioritiesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plan priorities get all o k response
func (o *PlanPrioritiesGetAllOK) Code() int {
	return 200
}

func (o *PlanPrioritiesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanPriorities][%d] planPrioritiesGetAllOK %s", 200, payload)
}

func (o *PlanPrioritiesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanPriorities][%d] planPrioritiesGetAllOK %s", 200, payload)
}

func (o *PlanPrioritiesGetAllOK) GetPayload() []*models.PlanPriority {
	return o.Payload
}

func (o *PlanPrioritiesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanPrioritiesGetAllDefault creates a PlanPrioritiesGetAllDefault with default headers values
func NewPlanPrioritiesGetAllDefault(code int) *PlanPrioritiesGetAllDefault {
	return &PlanPrioritiesGetAllDefault{
		_statusCode: code,
	}
}

/*
PlanPrioritiesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type PlanPrioritiesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this plan priorities get all default response has a 2xx status code
func (o *PlanPrioritiesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plan priorities get all default response has a 3xx status code
func (o *PlanPrioritiesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plan priorities get all default response has a 4xx status code
func (o *PlanPrioritiesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plan priorities get all default response has a 5xx status code
func (o *PlanPrioritiesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plan priorities get all default response a status code equal to that given
func (o *PlanPrioritiesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plan priorities get all default response
func (o *PlanPrioritiesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *PlanPrioritiesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanPriorities][%d] PlanPriorities_GetAll default %s", o._statusCode, payload)
}

func (o *PlanPrioritiesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanPriorities][%d] PlanPriorities_GetAll default %s", o._statusCode, payload)
}

func (o *PlanPrioritiesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PlanPrioritiesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
