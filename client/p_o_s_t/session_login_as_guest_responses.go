// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// SessionLoginAsGuestReader is a Reader for the SessionLoginAsGuest structure.
type SessionLoginAsGuestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionLoginAsGuestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSessionLoginAsGuestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Web/Session/{sessionKey}/LoginAsGuest] Session_LoginAsGuest", response, response.Code())
	}
}

// NewSessionLoginAsGuestOK creates a SessionLoginAsGuestOK with default headers values
func NewSessionLoginAsGuestOK() *SessionLoginAsGuestOK {
	return &SessionLoginAsGuestOK{}
}

/*
SessionLoginAsGuestOK describes a response with status code 200, with default header values.

OK
*/
type SessionLoginAsGuestOK struct {
	Payload *models.Session
}

// IsSuccess returns true when this session login as guest o k response has a 2xx status code
func (o *SessionLoginAsGuestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session login as guest o k response has a 3xx status code
func (o *SessionLoginAsGuestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session login as guest o k response has a 4xx status code
func (o *SessionLoginAsGuestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this session login as guest o k response has a 5xx status code
func (o *SessionLoginAsGuestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this session login as guest o k response a status code equal to that given
func (o *SessionLoginAsGuestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the session login as guest o k response
func (o *SessionLoginAsGuestOK) Code() int {
	return 200
}

func (o *SessionLoginAsGuestOK) Error() string {
	return fmt.Sprintf("[POST /Web/Session/{sessionKey}/LoginAsGuest][%d] sessionLoginAsGuestOK  %+v", 200, o.Payload)
}

func (o *SessionLoginAsGuestOK) String() string {
	return fmt.Sprintf("[POST /Web/Session/{sessionKey}/LoginAsGuest][%d] sessionLoginAsGuestOK  %+v", 200, o.Payload)
}

func (o *SessionLoginAsGuestOK) GetPayload() *models.Session {
	return o.Payload
}

func (o *SessionLoginAsGuestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Session)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
