// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SessionDeleteVariableReader is a Reader for the SessionDeleteVariable structure.
type SessionDeleteVariableReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionDeleteVariableReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSessionDeleteVariableNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /Web/Session/{sessionKey}/Variables/{variableName}] Session_DeleteVariable", response, response.Code())
	}
}

// NewSessionDeleteVariableNoContent creates a SessionDeleteVariableNoContent with default headers values
func NewSessionDeleteVariableNoContent() *SessionDeleteVariableNoContent {
	return &SessionDeleteVariableNoContent{}
}

/*
SessionDeleteVariableNoContent describes a response with status code 204, with default header values.

No Content
*/
type SessionDeleteVariableNoContent struct {
}

// IsSuccess returns true when this session delete variable no content response has a 2xx status code
func (o *SessionDeleteVariableNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session delete variable no content response has a 3xx status code
func (o *SessionDeleteVariableNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session delete variable no content response has a 4xx status code
func (o *SessionDeleteVariableNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this session delete variable no content response has a 5xx status code
func (o *SessionDeleteVariableNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this session delete variable no content response a status code equal to that given
func (o *SessionDeleteVariableNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the session delete variable no content response
func (o *SessionDeleteVariableNoContent) Code() int {
	return 204
}

func (o *SessionDeleteVariableNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Web/Session/{sessionKey}/Variables/{variableName}][%d] sessionDeleteVariableNoContent ", 204)
}

func (o *SessionDeleteVariableNoContent) String() string {
	return fmt.Sprintf("[DELETE /Web/Session/{sessionKey}/Variables/{variableName}][%d] sessionDeleteVariableNoContent ", 204)
}

func (o *SessionDeleteVariableNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
