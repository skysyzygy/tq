// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// PerformancesRescheduleReader is a Reader for the PerformancesReschedule structure.
type PerformancesRescheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancesRescheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPerformancesRescheduleNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancesRescheduleDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancesRescheduleNoContent creates a PerformancesRescheduleNoContent with default headers values
func NewPerformancesRescheduleNoContent() *PerformancesRescheduleNoContent {
	return &PerformancesRescheduleNoContent{}
}

/*
PerformancesRescheduleNoContent describes a response with status code 204, with default header values.

No Content
*/
type PerformancesRescheduleNoContent struct {
}

// IsSuccess returns true when this performances reschedule no content response has a 2xx status code
func (o *PerformancesRescheduleNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performances reschedule no content response has a 3xx status code
func (o *PerformancesRescheduleNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performances reschedule no content response has a 4xx status code
func (o *PerformancesRescheduleNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this performances reschedule no content response has a 5xx status code
func (o *PerformancesRescheduleNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this performances reschedule no content response a status code equal to that given
func (o *PerformancesRescheduleNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the performances reschedule no content response
func (o *PerformancesRescheduleNoContent) Code() int {
	return 204
}

func (o *PerformancesRescheduleNoContent) Error() string {
	return fmt.Sprintf("[POST /TXN/Performances/Reschedule][%d] performancesRescheduleNoContent", 204)
}

func (o *PerformancesRescheduleNoContent) String() string {
	return fmt.Sprintf("[POST /TXN/Performances/Reschedule][%d] performancesRescheduleNoContent", 204)
}

func (o *PerformancesRescheduleNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPerformancesRescheduleDefault creates a PerformancesRescheduleDefault with default headers values
func NewPerformancesRescheduleDefault(code int) *PerformancesRescheduleDefault {
	return &PerformancesRescheduleDefault{
		_statusCode: code,
	}
}

/*
PerformancesRescheduleDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancesRescheduleDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performances reschedule default response has a 2xx status code
func (o *PerformancesRescheduleDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performances reschedule default response has a 3xx status code
func (o *PerformancesRescheduleDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performances reschedule default response has a 4xx status code
func (o *PerformancesRescheduleDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performances reschedule default response has a 5xx status code
func (o *PerformancesRescheduleDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performances reschedule default response a status code equal to that given
func (o *PerformancesRescheduleDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performances reschedule default response
func (o *PerformancesRescheduleDefault) Code() int {
	return o._statusCode
}

func (o *PerformancesRescheduleDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/Performances/Reschedule][%d] Performances_Reschedule default %s", o._statusCode, payload)
}

func (o *PerformancesRescheduleDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/Performances/Reschedule][%d] Performances_Reschedule default %s", o._statusCode, payload)
}

func (o *PerformancesRescheduleDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancesRescheduleDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
