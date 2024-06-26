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

// SessionUpdateVariableReader is a Reader for the SessionUpdateVariable structure.
type SessionUpdateVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionUpdateVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSessionUpdateVariableOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSessionUpdateVariableDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSessionUpdateVariableOK creates a SessionUpdateVariableOK with default headers values
func NewSessionUpdateVariableOK() *SessionUpdateVariableOK {
	return &SessionUpdateVariableOK{}
}

/*
SessionUpdateVariableOK describes a response with status code 200, with default header values.

OK
*/
type SessionUpdateVariableOK struct {
	Payload *models.SessionVariable
}

// IsSuccess returns true when this session update variable o k response has a 2xx status code
func (o *SessionUpdateVariableOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session update variable o k response has a 3xx status code
func (o *SessionUpdateVariableOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session update variable o k response has a 4xx status code
func (o *SessionUpdateVariableOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this session update variable o k response has a 5xx status code
func (o *SessionUpdateVariableOK) IsServerError() bool {
	return false
}

// IsCode returns true when this session update variable o k response a status code equal to that given
func (o *SessionUpdateVariableOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the session update variable o k response
func (o *SessionUpdateVariableOK) Code() int {
	return 200
}

func (o *SessionUpdateVariableOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Web/Session/{sessionKey}/Variables][%d] sessionUpdateVariableOK %s", 200, payload)
}

func (o *SessionUpdateVariableOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Web/Session/{sessionKey}/Variables][%d] sessionUpdateVariableOK %s", 200, payload)
}

func (o *SessionUpdateVariableOK) GetPayload() *models.SessionVariable {
	return o.Payload
}

func (o *SessionUpdateVariableOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionVariable)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSessionUpdateVariableDefault creates a SessionUpdateVariableDefault with default headers values
func NewSessionUpdateVariableDefault(code int) *SessionUpdateVariableDefault {
	return &SessionUpdateVariableDefault{
		_statusCode: code,
	}
}

/*
SessionUpdateVariableDefault describes a response with status code -1, with default header values.

Error
*/
type SessionUpdateVariableDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this session update variable default response has a 2xx status code
func (o *SessionUpdateVariableDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this session update variable default response has a 3xx status code
func (o *SessionUpdateVariableDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this session update variable default response has a 4xx status code
func (o *SessionUpdateVariableDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this session update variable default response has a 5xx status code
func (o *SessionUpdateVariableDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this session update variable default response a status code equal to that given
func (o *SessionUpdateVariableDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the session update variable default response
func (o *SessionUpdateVariableDefault) Code() int {
	return o._statusCode
}

func (o *SessionUpdateVariableDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Web/Session/{sessionKey}/Variables][%d] Session_UpdateVariable default %s", o._statusCode, payload)
}

func (o *SessionUpdateVariableDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /Web/Session/{sessionKey}/Variables][%d] Session_UpdateVariable default %s", o._statusCode, payload)
}

func (o *SessionUpdateVariableDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SessionUpdateVariableDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
