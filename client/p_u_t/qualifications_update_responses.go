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

// QualificationsUpdateReader is a Reader for the QualificationsUpdate structure.
type QualificationsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QualificationsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQualificationsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQualificationsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQualificationsUpdateOK creates a QualificationsUpdateOK with default headers values
func NewQualificationsUpdateOK() *QualificationsUpdateOK {
	return &QualificationsUpdateOK{}
}

/*
QualificationsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type QualificationsUpdateOK struct {
	Payload *models.Qualification
}

// IsSuccess returns true when this qualifications update o k response has a 2xx status code
func (o *QualificationsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qualifications update o k response has a 3xx status code
func (o *QualificationsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qualifications update o k response has a 4xx status code
func (o *QualificationsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this qualifications update o k response has a 5xx status code
func (o *QualificationsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this qualifications update o k response a status code equal to that given
func (o *QualificationsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the qualifications update o k response
func (o *QualificationsUpdateOK) Code() int {
	return 200
}

func (o *QualificationsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Qualifications/{id}][%d] qualificationsUpdateOK %s", 200, payload)
}

func (o *QualificationsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Qualifications/{id}][%d] qualificationsUpdateOK %s", 200, payload)
}

func (o *QualificationsUpdateOK) GetPayload() *models.Qualification {
	return o.Payload
}

func (o *QualificationsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Qualification)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQualificationsUpdateDefault creates a QualificationsUpdateDefault with default headers values
func NewQualificationsUpdateDefault(code int) *QualificationsUpdateDefault {
	return &QualificationsUpdateDefault{
		_statusCode: code,
	}
}

/*
QualificationsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type QualificationsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this qualifications update default response has a 2xx status code
func (o *QualificationsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qualifications update default response has a 3xx status code
func (o *QualificationsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qualifications update default response has a 4xx status code
func (o *QualificationsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qualifications update default response has a 5xx status code
func (o *QualificationsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qualifications update default response a status code equal to that given
func (o *QualificationsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the qualifications update default response
func (o *QualificationsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *QualificationsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Qualifications/{id}][%d] Qualifications_Update default %s", o._statusCode, payload)
}

func (o *QualificationsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/Qualifications/{id}][%d] Qualifications_Update default %s", o._statusCode, payload)
}

func (o *QualificationsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *QualificationsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
