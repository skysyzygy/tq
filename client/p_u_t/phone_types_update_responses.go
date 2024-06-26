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

// PhoneTypesUpdateReader is a Reader for the PhoneTypesUpdate structure.
type PhoneTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhoneTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhoneTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPhoneTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPhoneTypesUpdateOK creates a PhoneTypesUpdateOK with default headers values
func NewPhoneTypesUpdateOK() *PhoneTypesUpdateOK {
	return &PhoneTypesUpdateOK{}
}

/*
PhoneTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PhoneTypesUpdateOK struct {
	Payload *models.PhoneType
}

// IsSuccess returns true when this phone types update o k response has a 2xx status code
func (o *PhoneTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phone types update o k response has a 3xx status code
func (o *PhoneTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phone types update o k response has a 4xx status code
func (o *PhoneTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phone types update o k response has a 5xx status code
func (o *PhoneTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phone types update o k response a status code equal to that given
func (o *PhoneTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phone types update o k response
func (o *PhoneTypesUpdateOK) Code() int {
	return 200
}

func (o *PhoneTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PhoneTypes/{id}][%d] phoneTypesUpdateOK %s", 200, payload)
}

func (o *PhoneTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PhoneTypes/{id}][%d] phoneTypesUpdateOK %s", 200, payload)
}

func (o *PhoneTypesUpdateOK) GetPayload() *models.PhoneType {
	return o.Payload
}

func (o *PhoneTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PhoneType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPhoneTypesUpdateDefault creates a PhoneTypesUpdateDefault with default headers values
func NewPhoneTypesUpdateDefault(code int) *PhoneTypesUpdateDefault {
	return &PhoneTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
PhoneTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PhoneTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this phone types update default response has a 2xx status code
func (o *PhoneTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this phone types update default response has a 3xx status code
func (o *PhoneTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this phone types update default response has a 4xx status code
func (o *PhoneTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this phone types update default response has a 5xx status code
func (o *PhoneTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this phone types update default response a status code equal to that given
func (o *PhoneTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the phone types update default response
func (o *PhoneTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PhoneTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PhoneTypes/{id}][%d] PhoneTypes_Update default %s", o._statusCode, payload)
}

func (o *PhoneTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PhoneTypes/{id}][%d] PhoneTypes_Update default %s", o._statusCode, payload)
}

func (o *PhoneTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PhoneTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
