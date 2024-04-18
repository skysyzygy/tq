// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// SessionGetExpirationReader is a Reader for the SessionGetExpiration structure.
type SessionGetExpirationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionGetExpirationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSessionGetExpirationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Web/Session/{sessionKey}/Expiration] Session_GetExpiration", response, response.Code())
	}
}

// NewSessionGetExpirationOK creates a SessionGetExpirationOK with default headers values
func NewSessionGetExpirationOK() *SessionGetExpirationOK {
	return &SessionGetExpirationOK{}
}

/*
SessionGetExpirationOK describes a response with status code 200, with default header values.

OK
*/
type SessionGetExpirationOK struct {
	Payload *models.SessionExpirationResponse
}

// IsSuccess returns true when this session get expiration o k response has a 2xx status code
func (o *SessionGetExpirationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session get expiration o k response has a 3xx status code
func (o *SessionGetExpirationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session get expiration o k response has a 4xx status code
func (o *SessionGetExpirationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this session get expiration o k response has a 5xx status code
func (o *SessionGetExpirationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this session get expiration o k response a status code equal to that given
func (o *SessionGetExpirationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the session get expiration o k response
func (o *SessionGetExpirationOK) Code() int {
	return 200
}

func (o *SessionGetExpirationOK) Error() string {
	return fmt.Sprintf("[GET /Web/Session/{sessionKey}/Expiration][%d] sessionGetExpirationOK  %+v", 200, o.Payload)
}

func (o *SessionGetExpirationOK) String() string {
	return fmt.Sprintf("[GET /Web/Session/{sessionKey}/Expiration][%d] sessionGetExpirationOK  %+v", 200, o.Payload)
}

func (o *SessionGetExpirationOK) GetPayload() *models.SessionExpirationResponse {
	return o.Payload
}

func (o *SessionGetExpirationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SessionExpirationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}