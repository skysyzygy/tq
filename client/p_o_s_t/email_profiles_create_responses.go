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

// EmailProfilesCreateReader is a Reader for the EmailProfilesCreate structure.
type EmailProfilesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailProfilesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEmailProfilesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/EmailProfiles] EmailProfiles_Create", response, response.Code())
	}
}

// NewEmailProfilesCreateOK creates a EmailProfilesCreateOK with default headers values
func NewEmailProfilesCreateOK() *EmailProfilesCreateOK {
	return &EmailProfilesCreateOK{}
}

/*
EmailProfilesCreateOK describes a response with status code 200, with default header values.

OK
*/
type EmailProfilesCreateOK struct {
	Payload *models.EmailProfile
}

// IsSuccess returns true when this email profiles create o k response has a 2xx status code
func (o *EmailProfilesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this email profiles create o k response has a 3xx status code
func (o *EmailProfilesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this email profiles create o k response has a 4xx status code
func (o *EmailProfilesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this email profiles create o k response has a 5xx status code
func (o *EmailProfilesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this email profiles create o k response a status code equal to that given
func (o *EmailProfilesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the email profiles create o k response
func (o *EmailProfilesCreateOK) Code() int {
	return 200
}

func (o *EmailProfilesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/EmailProfiles][%d] emailProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *EmailProfilesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/EmailProfiles][%d] emailProfilesCreateOK  %+v", 200, o.Payload)
}

func (o *EmailProfilesCreateOK) GetPayload() *models.EmailProfile {
	return o.Payload
}

func (o *EmailProfilesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EmailProfile)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
