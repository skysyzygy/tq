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

// CoreIdentityTokenReader is a Reader for the CoreIdentityToken structure.
type CoreIdentityTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CoreIdentityTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCoreIdentityTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Security/CoreIdentity/Token] CoreIdentity_Token", response, response.Code())
	}
}

// NewCoreIdentityTokenOK creates a CoreIdentityTokenOK with default headers values
func NewCoreIdentityTokenOK() *CoreIdentityTokenOK {
	return &CoreIdentityTokenOK{}
}

/*
CoreIdentityTokenOK describes a response with status code 200, with default header values.

OK
*/
type CoreIdentityTokenOK struct {
	Payload *models.CoreIdentityTokenResponse
}

// IsSuccess returns true when this core identity token o k response has a 2xx status code
func (o *CoreIdentityTokenOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this core identity token o k response has a 3xx status code
func (o *CoreIdentityTokenOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this core identity token o k response has a 4xx status code
func (o *CoreIdentityTokenOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this core identity token o k response has a 5xx status code
func (o *CoreIdentityTokenOK) IsServerError() bool {
	return false
}

// IsCode returns true when this core identity token o k response a status code equal to that given
func (o *CoreIdentityTokenOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the core identity token o k response
func (o *CoreIdentityTokenOK) Code() int {
	return 200
}

func (o *CoreIdentityTokenOK) Error() string {
	return fmt.Sprintf("[GET /Security/CoreIdentity/Token][%d] coreIdentityTokenOK  %+v", 200, o.Payload)
}

func (o *CoreIdentityTokenOK) String() string {
	return fmt.Sprintf("[GET /Security/CoreIdentity/Token][%d] coreIdentityTokenOK  %+v", 200, o.Payload)
}

func (o *CoreIdentityTokenOK) GetPayload() *models.CoreIdentityTokenResponse {
	return o.Payload
}

func (o *CoreIdentityTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreIdentityTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}