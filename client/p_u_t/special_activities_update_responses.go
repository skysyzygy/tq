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

// SpecialActivitiesUpdateReader is a Reader for the SpecialActivitiesUpdate structure.
type SpecialActivitiesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SpecialActivitiesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSpecialActivitiesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSpecialActivitiesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSpecialActivitiesUpdateOK creates a SpecialActivitiesUpdateOK with default headers values
func NewSpecialActivitiesUpdateOK() *SpecialActivitiesUpdateOK {
	return &SpecialActivitiesUpdateOK{}
}

/*
SpecialActivitiesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type SpecialActivitiesUpdateOK struct {
	Payload *models.SpecialActivity
}

// IsSuccess returns true when this special activities update o k response has a 2xx status code
func (o *SpecialActivitiesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this special activities update o k response has a 3xx status code
func (o *SpecialActivitiesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this special activities update o k response has a 4xx status code
func (o *SpecialActivitiesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this special activities update o k response has a 5xx status code
func (o *SpecialActivitiesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this special activities update o k response a status code equal to that given
func (o *SpecialActivitiesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the special activities update o k response
func (o *SpecialActivitiesUpdateOK) Code() int {
	return 200
}

func (o *SpecialActivitiesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/SpecialActivities/{specialActivityId}][%d] specialActivitiesUpdateOK %s", 200, payload)
}

func (o *SpecialActivitiesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/SpecialActivities/{specialActivityId}][%d] specialActivitiesUpdateOK %s", 200, payload)
}

func (o *SpecialActivitiesUpdateOK) GetPayload() *models.SpecialActivity {
	return o.Payload
}

func (o *SpecialActivitiesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SpecialActivity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSpecialActivitiesUpdateDefault creates a SpecialActivitiesUpdateDefault with default headers values
func NewSpecialActivitiesUpdateDefault(code int) *SpecialActivitiesUpdateDefault {
	return &SpecialActivitiesUpdateDefault{
		_statusCode: code,
	}
}

/*
SpecialActivitiesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type SpecialActivitiesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this special activities update default response has a 2xx status code
func (o *SpecialActivitiesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this special activities update default response has a 3xx status code
func (o *SpecialActivitiesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this special activities update default response has a 4xx status code
func (o *SpecialActivitiesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this special activities update default response has a 5xx status code
func (o *SpecialActivitiesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this special activities update default response a status code equal to that given
func (o *SpecialActivitiesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the special activities update default response
func (o *SpecialActivitiesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *SpecialActivitiesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/SpecialActivities/{specialActivityId}][%d] SpecialActivities_Update default %s", o._statusCode, payload)
}

func (o *SpecialActivitiesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /CRM/SpecialActivities/{specialActivityId}][%d] SpecialActivities_Update default %s", o._statusCode, payload)
}

func (o *SpecialActivitiesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SpecialActivitiesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
