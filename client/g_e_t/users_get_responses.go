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

// UsersGetReader is a Reader for the UsersGet structure.
type UsersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Security/Users/{userName}] Users_Get", response, response.Code())
	}
}

// NewUsersGetOK creates a UsersGetOK with default headers values
func NewUsersGetOK() *UsersGetOK {
	return &UsersGetOK{}
}

/*
UsersGetOK describes a response with status code 200, with default header values.

OK
*/
type UsersGetOK struct {
	Payload *models.User
}

// IsSuccess returns true when this users get o k response has a 2xx status code
func (o *UsersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this users get o k response has a 3xx status code
func (o *UsersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this users get o k response has a 4xx status code
func (o *UsersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this users get o k response has a 5xx status code
func (o *UsersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this users get o k response a status code equal to that given
func (o *UsersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the users get o k response
func (o *UsersGetOK) Code() int {
	return 200
}

func (o *UsersGetOK) Error() string {
	return fmt.Sprintf("[GET /Security/Users/{userName}][%d] usersGetOK  %+v", 200, o.Payload)
}

func (o *UsersGetOK) String() string {
	return fmt.Sprintf("[GET /Security/Users/{userName}][%d] usersGetOK  %+v", 200, o.Payload)
}

func (o *UsersGetOK) GetPayload() *models.User {
	return o.Payload
}

func (o *UsersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
