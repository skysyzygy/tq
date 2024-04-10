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

// UserPreferencesCreateReader is a Reader for the UserPreferencesCreate structure.
type UserPreferencesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserPreferencesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserPreferencesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Security/UserPreferences] UserPreferences_Create", response, response.Code())
	}
}

// NewUserPreferencesCreateOK creates a UserPreferencesCreateOK with default headers values
func NewUserPreferencesCreateOK() *UserPreferencesCreateOK {
	return &UserPreferencesCreateOK{}
}

/*
UserPreferencesCreateOK describes a response with status code 200, with default header values.

OK
*/
type UserPreferencesCreateOK struct {
	Payload *models.UserPreference
}

// IsSuccess returns true when this user preferences create o k response has a 2xx status code
func (o *UserPreferencesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user preferences create o k response has a 3xx status code
func (o *UserPreferencesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user preferences create o k response has a 4xx status code
func (o *UserPreferencesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user preferences create o k response has a 5xx status code
func (o *UserPreferencesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user preferences create o k response a status code equal to that given
func (o *UserPreferencesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user preferences create o k response
func (o *UserPreferencesCreateOK) Code() int {
	return 200
}

func (o *UserPreferencesCreateOK) Error() string {
	return fmt.Sprintf("[POST /Security/UserPreferences][%d] userPreferencesCreateOK  %+v", 200, o.Payload)
}

func (o *UserPreferencesCreateOK) String() string {
	return fmt.Sprintf("[POST /Security/UserPreferences][%d] userPreferencesCreateOK  %+v", 200, o.Payload)
}

func (o *UserPreferencesCreateOK) GetPayload() *models.UserPreference {
	return o.Payload
}

func (o *UserPreferencesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserPreference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
