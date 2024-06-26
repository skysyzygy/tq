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

// SpecialActivityStatusesCreateReader is a Reader for the SpecialActivityStatusesCreate structure.
type SpecialActivityStatusesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecialActivityStatusesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpecialActivityStatusesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSpecialActivityStatusesCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSpecialActivityStatusesCreateOK creates a SpecialActivityStatusesCreateOK with default headers values
func NewSpecialActivityStatusesCreateOK() *SpecialActivityStatusesCreateOK {
	return &SpecialActivityStatusesCreateOK{}
}

/*
SpecialActivityStatusesCreateOK describes a response with status code 200, with default header values.

OK
*/
type SpecialActivityStatusesCreateOK struct {
	Payload *models.SpecialActivityStatus
}

// IsSuccess returns true when this special activity statuses create o k response has a 2xx status code
func (o *SpecialActivityStatusesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this special activity statuses create o k response has a 3xx status code
func (o *SpecialActivityStatusesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this special activity statuses create o k response has a 4xx status code
func (o *SpecialActivityStatusesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this special activity statuses create o k response has a 5xx status code
func (o *SpecialActivityStatusesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this special activity statuses create o k response a status code equal to that given
func (o *SpecialActivityStatusesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the special activity statuses create o k response
func (o *SpecialActivityStatusesCreateOK) Code() int {
	return 200
}

func (o *SpecialActivityStatusesCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/SpecialActivityStatuses][%d] specialActivityStatusesCreateOK %s", 200, payload)
}

func (o *SpecialActivityStatusesCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/SpecialActivityStatuses][%d] specialActivityStatusesCreateOK %s", 200, payload)
}

func (o *SpecialActivityStatusesCreateOK) GetPayload() *models.SpecialActivityStatus {
	return o.Payload
}

func (o *SpecialActivityStatusesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpecialActivityStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecialActivityStatusesCreateDefault creates a SpecialActivityStatusesCreateDefault with default headers values
func NewSpecialActivityStatusesCreateDefault(code int) *SpecialActivityStatusesCreateDefault {
	return &SpecialActivityStatusesCreateDefault{
		_statusCode: code,
	}
}

/*
SpecialActivityStatusesCreateDefault describes a response with status code -1, with default header values.

Error
*/
type SpecialActivityStatusesCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this special activity statuses create default response has a 2xx status code
func (o *SpecialActivityStatusesCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this special activity statuses create default response has a 3xx status code
func (o *SpecialActivityStatusesCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this special activity statuses create default response has a 4xx status code
func (o *SpecialActivityStatusesCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this special activity statuses create default response has a 5xx status code
func (o *SpecialActivityStatusesCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this special activity statuses create default response a status code equal to that given
func (o *SpecialActivityStatusesCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the special activity statuses create default response
func (o *SpecialActivityStatusesCreateDefault) Code() int {
	return o._statusCode
}

func (o *SpecialActivityStatusesCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/SpecialActivityStatuses][%d] SpecialActivityStatuses_Create default %s", o._statusCode, payload)
}

func (o *SpecialActivityStatusesCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /ReferenceData/SpecialActivityStatuses][%d] SpecialActivityStatuses_Create default %s", o._statusCode, payload)
}

func (o *SpecialActivityStatusesCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SpecialActivityStatusesCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
