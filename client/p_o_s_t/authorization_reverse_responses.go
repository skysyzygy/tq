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

// AuthorizationReverseReader is a Reader for the AuthorizationReverse structure.
type AuthorizationReverseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AuthorizationReverseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAuthorizationReverseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /PaymentGateway/Authorization/{referenceNumber}/Reverse] Authorization_Reverse", response, response.Code())
	}
}

// NewAuthorizationReverseOK creates a AuthorizationReverseOK with default headers values
func NewAuthorizationReverseOK() *AuthorizationReverseOK {
	return &AuthorizationReverseOK{}
}

/*
AuthorizationReverseOK describes a response with status code 200, with default header values.

OK
*/
type AuthorizationReverseOK struct {
	Payload *models.ReversalResponse
}

// IsSuccess returns true when this authorization reverse o k response has a 2xx status code
func (o *AuthorizationReverseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this authorization reverse o k response has a 3xx status code
func (o *AuthorizationReverseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this authorization reverse o k response has a 4xx status code
func (o *AuthorizationReverseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this authorization reverse o k response has a 5xx status code
func (o *AuthorizationReverseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this authorization reverse o k response a status code equal to that given
func (o *AuthorizationReverseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the authorization reverse o k response
func (o *AuthorizationReverseOK) Code() int {
	return 200
}

func (o *AuthorizationReverseOK) Error() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/{referenceNumber}/Reverse][%d] authorizationReverseOK  %+v", 200, o.Payload)
}

func (o *AuthorizationReverseOK) String() string {
	return fmt.Sprintf("[POST /PaymentGateway/Authorization/{referenceNumber}/Reverse][%d] authorizationReverseOK  %+v", 200, o.Payload)
}

func (o *AuthorizationReverseOK) GetPayload() *models.ReversalResponse {
	return o.Payload
}

func (o *AuthorizationReverseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReversalResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
