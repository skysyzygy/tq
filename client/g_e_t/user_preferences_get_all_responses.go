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

// UserPreferencesGetAllReader is a Reader for the UserPreferencesGetAll structure.
type UserPreferencesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UserPreferencesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUserPreferencesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUserPreferencesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUserPreferencesGetAllOK creates a UserPreferencesGetAllOK with default headers values
func NewUserPreferencesGetAllOK() *UserPreferencesGetAllOK {
	return &UserPreferencesGetAllOK{}
}

/*
UserPreferencesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type UserPreferencesGetAllOK struct {
	Payload []*models.UserPreference
}

// IsSuccess returns true when this user preferences get all o k response has a 2xx status code
func (o *UserPreferencesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this user preferences get all o k response has a 3xx status code
func (o *UserPreferencesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this user preferences get all o k response has a 4xx status code
func (o *UserPreferencesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this user preferences get all o k response has a 5xx status code
func (o *UserPreferencesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this user preferences get all o k response a status code equal to that given
func (o *UserPreferencesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the user preferences get all o k response
func (o *UserPreferencesGetAllOK) Code() int {
	return 200
}

func (o *UserPreferencesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/UserPreferences][%d] userPreferencesGetAllOK %s", 200, payload)
}

func (o *UserPreferencesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/UserPreferences][%d] userPreferencesGetAllOK %s", 200, payload)
}

func (o *UserPreferencesGetAllOK) GetPayload() []*models.UserPreference {
	return o.Payload
}

func (o *UserPreferencesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUserPreferencesGetAllDefault creates a UserPreferencesGetAllDefault with default headers values
func NewUserPreferencesGetAllDefault(code int) *UserPreferencesGetAllDefault {
	return &UserPreferencesGetAllDefault{
		_statusCode: code,
	}
}

/*
UserPreferencesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type UserPreferencesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this user preferences get all default response has a 2xx status code
func (o *UserPreferencesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this user preferences get all default response has a 3xx status code
func (o *UserPreferencesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this user preferences get all default response has a 4xx status code
func (o *UserPreferencesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this user preferences get all default response has a 5xx status code
func (o *UserPreferencesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this user preferences get all default response a status code equal to that given
func (o *UserPreferencesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the user preferences get all default response
func (o *UserPreferencesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *UserPreferencesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/UserPreferences][%d] UserPreferences_GetAll default %s", o._statusCode, payload)
}

func (o *UserPreferencesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Security/UserPreferences][%d] UserPreferences_GetAll default %s", o._statusCode, payload)
}

func (o *UserPreferencesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *UserPreferencesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
