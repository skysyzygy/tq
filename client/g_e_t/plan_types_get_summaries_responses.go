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

// PlanTypesGetSummariesReader is a Reader for the PlanTypesGetSummaries structure.
type PlanTypesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PlanTypesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPlanTypesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPlanTypesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPlanTypesGetSummariesOK creates a PlanTypesGetSummariesOK with default headers values
func NewPlanTypesGetSummariesOK() *PlanTypesGetSummariesOK {
	return &PlanTypesGetSummariesOK{}
}

/*
PlanTypesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type PlanTypesGetSummariesOK struct {
	Payload []*models.PlanTypeSummary
}

// IsSuccess returns true when this plan types get summaries o k response has a 2xx status code
func (o *PlanTypesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this plan types get summaries o k response has a 3xx status code
func (o *PlanTypesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this plan types get summaries o k response has a 4xx status code
func (o *PlanTypesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this plan types get summaries o k response has a 5xx status code
func (o *PlanTypesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this plan types get summaries o k response a status code equal to that given
func (o *PlanTypesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the plan types get summaries o k response
func (o *PlanTypesGetSummariesOK) Code() int {
	return 200
}

func (o *PlanTypesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanTypes/Summary][%d] planTypesGetSummariesOK %s", 200, payload)
}

func (o *PlanTypesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanTypes/Summary][%d] planTypesGetSummariesOK %s", 200, payload)
}

func (o *PlanTypesGetSummariesOK) GetPayload() []*models.PlanTypeSummary {
	return o.Payload
}

func (o *PlanTypesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPlanTypesGetSummariesDefault creates a PlanTypesGetSummariesDefault with default headers values
func NewPlanTypesGetSummariesDefault(code int) *PlanTypesGetSummariesDefault {
	return &PlanTypesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
PlanTypesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type PlanTypesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this plan types get summaries default response has a 2xx status code
func (o *PlanTypesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this plan types get summaries default response has a 3xx status code
func (o *PlanTypesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this plan types get summaries default response has a 4xx status code
func (o *PlanTypesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this plan types get summaries default response has a 5xx status code
func (o *PlanTypesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this plan types get summaries default response a status code equal to that given
func (o *PlanTypesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the plan types get summaries default response
func (o *PlanTypesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *PlanTypesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanTypes/Summary][%d] PlanTypes_GetSummaries default %s", o._statusCode, payload)
}

func (o *PlanTypesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PlanTypes/Summary][%d] PlanTypes_GetSummaries default %s", o._statusCode, payload)
}

func (o *PlanTypesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PlanTypesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
