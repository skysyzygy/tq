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

// AuthorizationConfirmPayByLinkReader is a Reader for the AuthorizationConfirmPayByLink structure.
type AuthorizationConfirmPayByLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationConfirmPayByLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationConfirmPayByLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /PaymentGateway/Authorization/Link/{paymentId}/Confirm] Authorization_ConfirmPayByLink", response, response.Code())
	}
}

// NewAuthorizationConfirmPayByLinkOK creates a AuthorizationConfirmPayByLinkOK with default headers values
func NewAuthorizationConfirmPayByLinkOK() *AuthorizationConfirmPayByLinkOK {
	return &AuthorizationConfirmPayByLinkOK{}
}

/*
AuthorizationConfirmPayByLinkOK describes a response with status code 200, with default header values.

OK
*/
type AuthorizationConfirmPayByLinkOK struct {
	Payload *models.PayByLinkAuthorizationResponse
}

// IsSuccess returns true when this authorization confirm pay by link o k response has a 2xx status code
func (o *AuthorizationConfirmPayByLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization confirm pay by link o k response has a 3xx status code
func (o *AuthorizationConfirmPayByLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization confirm pay by link o k response has a 4xx status code
func (o *AuthorizationConfirmPayByLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization confirm pay by link o k response has a 5xx status code
func (o *AuthorizationConfirmPayByLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization confirm pay by link o k response a status code equal to that given
func (o *AuthorizationConfirmPayByLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization confirm pay by link o k response
func (o *AuthorizationConfirmPayByLinkOK) Code() int {
	return 200
}

func (o *AuthorizationConfirmPayByLinkOK) Error() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/Link/{paymentId}/Confirm][%d] authorizationConfirmPayByLinkOK  %+v", 200, o.Payload)
}

func (o *AuthorizationConfirmPayByLinkOK) String() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/Link/{paymentId}/Confirm][%d] authorizationConfirmPayByLinkOK  %+v", 200, o.Payload)
}

func (o *AuthorizationConfirmPayByLinkOK) GetPayload() *models.PayByLinkAuthorizationResponse {
	return o.Payload
}

func (o *AuthorizationConfirmPayByLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PayByLinkAuthorizationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}