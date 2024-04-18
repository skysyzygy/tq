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

// SessionGetDeliveryMethodsReader is a Reader for the SessionGetDeliveryMethods structure.
type SessionGetDeliveryMethodsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SessionGetDeliveryMethodsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSessionGetDeliveryMethodsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Web/Session/{sessionKey}/DeliveryMethods] Session_GetDeliveryMethods", response, response.Code())
	}
}

// NewSessionGetDeliveryMethodsOK creates a SessionGetDeliveryMethodsOK with default headers values
func NewSessionGetDeliveryMethodsOK() *SessionGetDeliveryMethodsOK {
	return &SessionGetDeliveryMethodsOK{}
}

/*
SessionGetDeliveryMethodsOK describes a response with status code 200, with default header values.

OK
*/
type SessionGetDeliveryMethodsOK struct {
	Payload []*models.WebDeliveryMethod
}

// IsSuccess returns true when this session get delivery methods o k response has a 2xx status code
func (o *SessionGetDeliveryMethodsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this session get delivery methods o k response has a 3xx status code
func (o *SessionGetDeliveryMethodsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this session get delivery methods o k response has a 4xx status code
func (o *SessionGetDeliveryMethodsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this session get delivery methods o k response has a 5xx status code
func (o *SessionGetDeliveryMethodsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this session get delivery methods o k response a status code equal to that given
func (o *SessionGetDeliveryMethodsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the session get delivery methods o k response
func (o *SessionGetDeliveryMethodsOK) Code() int {
	return 200
}

func (o *SessionGetDeliveryMethodsOK) Error() string {
	return fmt.Sprintf("[GET /Web/Session/{sessionKey}/DeliveryMethods][%d] sessionGetDeliveryMethodsOK  %+v", 200, o.Payload)
}

func (o *SessionGetDeliveryMethodsOK) String() string {
	return fmt.Sprintf("[GET /Web/Session/{sessionKey}/DeliveryMethods][%d] sessionGetDeliveryMethodsOK  %+v", 200, o.Payload)
}

func (o *SessionGetDeliveryMethodsOK) GetPayload() []*models.WebDeliveryMethod {
	return o.Payload
}

func (o *SessionGetDeliveryMethodsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}