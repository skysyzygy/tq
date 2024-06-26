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

// AuthorizationExpirePayByLinkReader is a Reader for the AuthorizationExpirePayByLink structure.
type AuthorizationExpirePayByLinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationExpirePayByLinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationExpirePayByLinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAuthorizationExpirePayByLinkDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAuthorizationExpirePayByLinkOK creates a AuthorizationExpirePayByLinkOK with default headers values
func NewAuthorizationExpirePayByLinkOK() *AuthorizationExpirePayByLinkOK {
	return &AuthorizationExpirePayByLinkOK{}
}

/*
AuthorizationExpirePayByLinkOK describes a response with status code 200, with default header values.

OK
*/
type AuthorizationExpirePayByLinkOK struct {
	Payload *models.ExpirePayByLinkResponse
}

// IsSuccess returns true when this authorization expire pay by link o k response has a 2xx status code
func (o *AuthorizationExpirePayByLinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization expire pay by link o k response has a 3xx status code
func (o *AuthorizationExpirePayByLinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization expire pay by link o k response has a 4xx status code
func (o *AuthorizationExpirePayByLinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization expire pay by link o k response has a 5xx status code
func (o *AuthorizationExpirePayByLinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization expire pay by link o k response a status code equal to that given
func (o *AuthorizationExpirePayByLinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization expire pay by link o k response
func (o *AuthorizationExpirePayByLinkOK) Code() int {
	return 200
}

func (o *AuthorizationExpirePayByLinkOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /PaymentGateway/Authorization/Link/{paymentId}][%d] authorizationExpirePayByLinkOK %s", 200, payload)
}

func (o *AuthorizationExpirePayByLinkOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /PaymentGateway/Authorization/Link/{paymentId}][%d] authorizationExpirePayByLinkOK %s", 200, payload)
}

func (o *AuthorizationExpirePayByLinkOK) GetPayload() *models.ExpirePayByLinkResponse {
	return o.Payload
}

func (o *AuthorizationExpirePayByLinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExpirePayByLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAuthorizationExpirePayByLinkDefault creates a AuthorizationExpirePayByLinkDefault with default headers values
func NewAuthorizationExpirePayByLinkDefault(code int) *AuthorizationExpirePayByLinkDefault {
	return &AuthorizationExpirePayByLinkDefault{
		_statusCode: code,
	}
}

/*
AuthorizationExpirePayByLinkDefault describes a response with status code -1, with default header values.

Error
*/
type AuthorizationExpirePayByLinkDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this authorization expire pay by link default response has a 2xx status code
func (o *AuthorizationExpirePayByLinkDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this authorization expire pay by link default response has a 3xx status code
func (o *AuthorizationExpirePayByLinkDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this authorization expire pay by link default response has a 4xx status code
func (o *AuthorizationExpirePayByLinkDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this authorization expire pay by link default response has a 5xx status code
func (o *AuthorizationExpirePayByLinkDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this authorization expire pay by link default response a status code equal to that given
func (o *AuthorizationExpirePayByLinkDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the authorization expire pay by link default response
func (o *AuthorizationExpirePayByLinkDefault) Code() int {
	return o._statusCode
}

func (o *AuthorizationExpirePayByLinkDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /PaymentGateway/Authorization/Link/{paymentId}][%d] Authorization_ExpirePayByLink default %s", o._statusCode, payload)
}

func (o *AuthorizationExpirePayByLinkDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /PaymentGateway/Authorization/Link/{paymentId}][%d] Authorization_ExpirePayByLink default %s", o._statusCode, payload)
}

func (o *AuthorizationExpirePayByLinkDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *AuthorizationExpirePayByLinkDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
