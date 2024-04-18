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

// ControlGroupUserGroupsGetReader is a Reader for the ControlGroupUserGroupsGet structure.
type ControlGroupUserGroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControlGroupUserGroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewControlGroupUserGroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ControlGroupUserGroups/{id}] ControlGroupUserGroups_Get", response, response.Code())
	}
}

// NewControlGroupUserGroupsGetOK creates a ControlGroupUserGroupsGetOK with default headers values
func NewControlGroupUserGroupsGetOK() *ControlGroupUserGroupsGetOK {
	return &ControlGroupUserGroupsGetOK{}
}

/*
ControlGroupUserGroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type ControlGroupUserGroupsGetOK struct {
	Payload *models.ControlGroupUserGroup
}

// IsSuccess returns true when this control group user groups get o k response has a 2xx status code
func (o *ControlGroupUserGroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this control group user groups get o k response has a 3xx status code
func (o *ControlGroupUserGroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control group user groups get o k response has a 4xx status code
func (o *ControlGroupUserGroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this control group user groups get o k response has a 5xx status code
func (o *ControlGroupUserGroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this control group user groups get o k response a status code equal to that given
func (o *ControlGroupUserGroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the control group user groups get o k response
func (o *ControlGroupUserGroupsGetOK) Code() int {
	return 200
}

func (o *ControlGroupUserGroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ControlGroupUserGroups/{id}][%d] controlGroupUserGroupsGetOK  %+v", 200, o.Payload)
}

func (o *ControlGroupUserGroupsGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ControlGroupUserGroups/{id}][%d] controlGroupUserGroupsGetOK  %+v", 200, o.Payload)
}

func (o *ControlGroupUserGroupsGetOK) GetPayload() *models.ControlGroupUserGroup {
	return o.Payload
}

func (o *ControlGroupUserGroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ControlGroupUserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}