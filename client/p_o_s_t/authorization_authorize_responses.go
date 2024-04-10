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

// AuthorizationAuthorizeReader is a Reader for the AuthorizationAuthorize structure.
type AuthorizationAuthorizeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationAuthorizeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationAuthorizeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /PaymentGateway/Authorization/Authorize] Authorization_Authorize", response, response.Code())
	}
}

// NewAuthorizationAuthorizeOK creates a AuthorizationAuthorizeOK with default headers values
func NewAuthorizationAuthorizeOK() *AuthorizationAuthorizeOK {
	return &AuthorizationAuthorizeOK{}
}

/*
AuthorizationAuthorizeOK describes a response with status code 200, with default header values.

OK
*/
type AuthorizationAuthorizeOK struct {
	Payload *models.AuthorizationResponse
}

// IsSuccess returns true when this authorization authorize o k response has a 2xx status code
func (o *AuthorizationAuthorizeOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization authorize o k response has a 3xx status code
func (o *AuthorizationAuthorizeOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization authorize o k response has a 4xx status code
func (o *AuthorizationAuthorizeOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization authorize o k response has a 5xx status code
func (o *AuthorizationAuthorizeOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization authorize o k response a status code equal to that given
func (o *AuthorizationAuthorizeOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization authorize o k response
func (o *AuthorizationAuthorizeOK) Code() int {
	return 200
}

func (o *AuthorizationAuthorizeOK) Error() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/Authorize][%d] authorizationAuthorizeOK  %+v", 200, o.Payload)
}

func (o *AuthorizationAuthorizeOK) String() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/Authorize][%d] authorizationAuthorizeOK  %+v", 200, o.Payload)
}

func (o *AuthorizationAuthorizeOK) GetPayload() *models.AuthorizationResponse {
	return o.Payload
}

func (o *AuthorizationAuthorizeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthorizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
