// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// UsersGetUsersForUserGroupReader is a Reader for the UsersGetUsersForUserGroup structure.
type UsersGetUsersForUserGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGetUsersForUserGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGetUsersForUserGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsersGetUsersForUserGroupDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsersGetUsersForUserGroupOK creates a UsersGetUsersForUserGroupOK with default headers values
func NewUsersGetUsersForUserGroupOK() *UsersGetUsersForUserGroupOK {
	return &UsersGetUsersForUserGroupOK{}
}

/*
UsersGetUsersForUserGroupOK describes a response with status code 200, with default header values.

OK
*/
type UsersGetUsersForUserGroupOK struct {
	Payload []*models.User
}

// IsSuccess returns true when this users get users for user group o k response has a 2xx status code
func (o *UsersGetUsersForUserGroupOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users get users for user group o k response has a 3xx status code
func (o *UsersGetUsersForUserGroupOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users get users for user group o k response has a 4xx status code
func (o *UsersGetUsersForUserGroupOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users get users for user group o k response has a 5xx status code
func (o *UsersGetUsersForUserGroupOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users get users for user group o k response a status code equal to that given
func (o *UsersGetUsersForUserGroupOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users get users for user group o k response
func (o *UsersGetUsersForUserGroupOK) Code() int {
	return 200
}

func (o *UsersGetUsersForUserGroupOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/Users/ForGroup][%d] usersGetUsersForUserGroupOK %s", 200, payload)
}

func (o *UsersGetUsersForUserGroupOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/Users/ForGroup][%d] usersGetUsersForUserGroupOK %s", 200, payload)
}

func (o *UsersGetUsersForUserGroupOK) GetPayload() []*models.User {
	return o.Payload
}

func (o *UsersGetUsersForUserGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsersGetUsersForUserGroupDefault creates a UsersGetUsersForUserGroupDefault with default headers values
func NewUsersGetUsersForUserGroupDefault(code int) *UsersGetUsersForUserGroupDefault {
	return &UsersGetUsersForUserGroupDefault{
		_statusCode: code,
	}
}

/*
UsersGetUsersForUserGroupDefault describes a response with status code -1, with default header values.

Error
*/
type UsersGetUsersForUserGroupDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this users get users for user group default response has a 2xx status code
func (o *UsersGetUsersForUserGroupDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this users get users for user group default response has a 3xx status code
func (o *UsersGetUsersForUserGroupDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this users get users for user group default response has a 4xx status code
func (o *UsersGetUsersForUserGroupDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this users get users for user group default response has a 5xx status code
func (o *UsersGetUsersForUserGroupDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this users get users for user group default response a status code equal to that given
func (o *UsersGetUsersForUserGroupDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the users get users for user group default response
func (o *UsersGetUsersForUserGroupDefault) Code() int {
	return o._statusCode
}

func (o *UsersGetUsersForUserGroupDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/Users/ForGroup][%d] Users_GetUsersForUserGroup default %s", o._statusCode, payload)
}

func (o *UsersGetUsersForUserGroupDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/Users/ForGroup][%d] Users_GetUsersForUserGroup default %s", o._statusCode, payload)
}

func (o *UsersGetUsersForUserGroupDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UsersGetUsersForUserGroupDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
