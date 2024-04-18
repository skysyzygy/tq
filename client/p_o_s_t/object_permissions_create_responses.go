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

// ObjectPermissionsCreateReader is a Reader for the ObjectPermissionsCreate structure.
type ObjectPermissionsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectPermissionsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectPermissionsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/ObjectPermissions] ObjectPermissions_Create", response, response.Code())
	}
}

// NewObjectPermissionsCreateOK creates a ObjectPermissionsCreateOK with default headers values
func NewObjectPermissionsCreateOK() *ObjectPermissionsCreateOK {
	return &ObjectPermissionsCreateOK{}
}

/*
ObjectPermissionsCreateOK describes a response with status code 200, with default header values.

OK
*/
type ObjectPermissionsCreateOK struct {
	Payload *models.ObjectPermission
}

// IsSuccess returns true when this object permissions create o k response has a 2xx status code
func (o *ObjectPermissionsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this object permissions create o k response has a 3xx status code
func (o *ObjectPermissionsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this object permissions create o k response has a 4xx status code
func (o *ObjectPermissionsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this object permissions create o k response has a 5xx status code
func (o *ObjectPermissionsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this object permissions create o k response a status code equal to that given
func (o *ObjectPermissionsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the object permissions create o k response
func (o *ObjectPermissionsCreateOK) Code() int {
	return 200
}

func (o *ObjectPermissionsCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/ObjectPermissions][%d] objectPermissionsCreateOK  %+v", 200, o.Payload)
}

func (o *ObjectPermissionsCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/ObjectPermissions][%d] objectPermissionsCreateOK  %+v", 200, o.Payload)
}

func (o *ObjectPermissionsCreateOK) GetPayload() *models.ObjectPermission {
	return o.Payload
}

func (o *ObjectPermissionsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ObjectPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}