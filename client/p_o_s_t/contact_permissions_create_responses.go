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

// ContactPermissionsCreateReader is a Reader for the ContactPermissionsCreate structure.
type ContactPermissionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPermissionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPermissionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /CRM/ContactPermissions] ContactPermissions_Create", response, response.Code())
	}
}

// NewContactPermissionsCreateOK creates a ContactPermissionsCreateOK with default headers values
func NewContactPermissionsCreateOK() *ContactPermissionsCreateOK {
	return &ContactPermissionsCreateOK{}
}

/*
ContactPermissionsCreateOK describes a response with status code 200, with default header values.

OK
*/
type ContactPermissionsCreateOK struct {
	Payload *models.ContactPermission
}

// IsSuccess returns true when this contact permissions create o k response has a 2xx status code
func (o *ContactPermissionsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact permissions create o k response has a 3xx status code
func (o *ContactPermissionsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact permissions create o k response has a 4xx status code
func (o *ContactPermissionsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact permissions create o k response has a 5xx status code
func (o *ContactPermissionsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact permissions create o k response a status code equal to that given
func (o *ContactPermissionsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact permissions create o k response
func (o *ContactPermissionsCreateOK) Code() int {
	return 200
}

func (o *ContactPermissionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /CRM/ContactPermissions][%d] contactPermissionsCreateOK  %+v", 200, o.Payload)
}

func (o *ContactPermissionsCreateOK) String() string {
	return fmt.Sprintf("[POST /CRM/ContactPermissions][%d] contactPermissionsCreateOK  %+v", 200, o.Payload)
}

func (o *ContactPermissionsCreateOK) GetPayload() *models.ContactPermission {
	return o.Payload
}

func (o *ContactPermissionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContactPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}