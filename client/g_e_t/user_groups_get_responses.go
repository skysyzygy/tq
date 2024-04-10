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

// UserGroupsGetReader is a Reader for the UserGroupsGet structure.
type UserGroupsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserGroupsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserGroupsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/UserGroups/{id}] UserGroups_Get", response, response.Code())
	}
}

// NewUserGroupsGetOK creates a UserGroupsGetOK with default headers values
func NewUserGroupsGetOK() *UserGroupsGetOK {
	return &UserGroupsGetOK{}
}

/*
UserGroupsGetOK describes a response with status code 200, with default header values.

OK
*/
type UserGroupsGetOK struct {
	Payload *models.UserGroup
}

// IsSuccess returns true when this user groups get o k response has a 2xx status code
func (o *UserGroupsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user groups get o k response has a 3xx status code
func (o *UserGroupsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user groups get o k response has a 4xx status code
func (o *UserGroupsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user groups get o k response has a 5xx status code
func (o *UserGroupsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user groups get o k response a status code equal to that given
func (o *UserGroupsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user groups get o k response
func (o *UserGroupsGetOK) Code() int {
	return 200
}

func (o *UserGroupsGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/UserGroups/{id}][%d] userGroupsGetOK  %+v", 200, o.Payload)
}

func (o *UserGroupsGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/UserGroups/{id}][%d] userGroupsGetOK  %+v", 200, o.Payload)
}

func (o *UserGroupsGetOK) GetPayload() *models.UserGroup {
	return o.Payload
}

func (o *UserGroupsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
