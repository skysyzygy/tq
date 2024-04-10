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

// AuthorizationLinkReader is a Reader for the AuthorizationLink structure.
type AuthorizationLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /PaymentGateway/Authorization/Link] Authorization_Link", response, response.Code())
	}
}

// NewAuthorizationLinkOK creates a AuthorizationLinkOK with default headers values
func NewAuthorizationLinkOK() *AuthorizationLinkOK {
	return &AuthorizationLinkOK{}
}

/*
AuthorizationLinkOK describes a response with status code 200, with default header values.

OK
*/
type AuthorizationLinkOK struct {
	Payload *models.PayByLinkResponse
}

// IsSuccess returns true when this authorization link o k response has a 2xx status code
func (o *AuthorizationLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization link o k response has a 3xx status code
func (o *AuthorizationLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization link o k response has a 4xx status code
func (o *AuthorizationLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization link o k response has a 5xx status code
func (o *AuthorizationLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization link o k response a status code equal to that given
func (o *AuthorizationLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization link o k response
func (o *AuthorizationLinkOK) Code() int {
	return 200
}

func (o *AuthorizationLinkOK) Error() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/Link][%d] authorizationLinkOK  %+v", 200, o.Payload)
}

func (o *AuthorizationLinkOK) String() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/Link][%d] authorizationLinkOK  %+v", 200, o.Payload)
}

func (o *AuthorizationLinkOK) GetPayload() *models.PayByLinkResponse {
	return o.Payload
}

func (o *AuthorizationLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayByLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
