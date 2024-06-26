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

// SessionLoginUsingTokenReader is a Reader for the SessionLoginUsingToken structure.
type SessionLoginUsingTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionLoginUsingTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSessionLoginUsingTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSessionLoginUsingTokenDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSessionLoginUsingTokenOK creates a SessionLoginUsingTokenOK with default headers values
func NewSessionLoginUsingTokenOK() *SessionLoginUsingTokenOK {
	return &SessionLoginUsingTokenOK{}
}

/*
SessionLoginUsingTokenOK describes a response with status code 200, with default header values.

OK
*/
type SessionLoginUsingTokenOK struct {
	Payload *models.Session
}

// IsSuccess returns true when this session login using token o k response has a 2xx status code
func (o *SessionLoginUsingTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session login using token o k response has a 3xx status code
func (o *SessionLoginUsingTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session login using token o k response has a 4xx status code
func (o *SessionLoginUsingTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this session login using token o k response has a 5xx status code
func (o *SessionLoginUsingTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this session login using token o k response a status code equal to that given
func (o *SessionLoginUsingTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the session login using token o k response
func (o *SessionLoginUsingTokenOK) Code() int {
	return 200
}

func (o *SessionLoginUsingTokenOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Session/{sessionKey}/Login/Token][%d] sessionLoginUsingTokenOK %s", 200, payload)
}

func (o *SessionLoginUsingTokenOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Session/{sessionKey}/Login/Token][%d] sessionLoginUsingTokenOK %s", 200, payload)
}

func (o *SessionLoginUsingTokenOK) GetPayload() *models.Session {
	return o.Payload
}

func (o *SessionLoginUsingTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSessionLoginUsingTokenDefault creates a SessionLoginUsingTokenDefault with default headers values
func NewSessionLoginUsingTokenDefault(code int) *SessionLoginUsingTokenDefault {
	return &SessionLoginUsingTokenDefault{
		_statusCode: code,
	}
}

/*
SessionLoginUsingTokenDefault describes a response with status code -1, with default header values.

Error
*/
type SessionLoginUsingTokenDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this session login using token default response has a 2xx status code
func (o *SessionLoginUsingTokenDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this session login using token default response has a 3xx status code
func (o *SessionLoginUsingTokenDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this session login using token default response has a 4xx status code
func (o *SessionLoginUsingTokenDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this session login using token default response has a 5xx status code
func (o *SessionLoginUsingTokenDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this session login using token default response a status code equal to that given
func (o *SessionLoginUsingTokenDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the session login using token default response
func (o *SessionLoginUsingTokenDefault) Code() int {
	return o._statusCode
}

func (o *SessionLoginUsingTokenDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Session/{sessionKey}/Login/Token][%d] Session_LoginUsingToken default %s", o._statusCode, payload)
}

func (o *SessionLoginUsingTokenDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Session/{sessionKey}/Login/Token][%d] Session_LoginUsingToken default %s", o._statusCode, payload)
}

func (o *SessionLoginUsingTokenDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SessionLoginUsingTokenDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
