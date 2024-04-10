// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// EmailsSendConstituentInfoReader is a Reader for the EmailsSendConstituentInfo structure.
type EmailsSendConstituentInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailsSendConstituentInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewEmailsSendConstituentInfoNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Emails/ConstituentInfo/{constituentId}/Send] Emails_SendConstituentInfo", response, response.Code())
	}
}

// NewEmailsSendConstituentInfoNoContent creates a EmailsSendConstituentInfoNoContent with default headers values
func NewEmailsSendConstituentInfoNoContent() *EmailsSendConstituentInfoNoContent {
	return &EmailsSendConstituentInfoNoContent{}
}

/*
EmailsSendConstituentInfoNoContent describes a response with status code 204, with default header values.

No Content
*/
type EmailsSendConstituentInfoNoContent struct {
}

// IsSuccess returns true when this emails send constituent info no content response has a 2xx status code
func (o *EmailsSendConstituentInfoNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this emails send constituent info no content response has a 3xx status code
func (o *EmailsSendConstituentInfoNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this emails send constituent info no content response has a 4xx status code
func (o *EmailsSendConstituentInfoNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this emails send constituent info no content response has a 5xx status code
func (o *EmailsSendConstituentInfoNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this emails send constituent info no content response a status code equal to that given
func (o *EmailsSendConstituentInfoNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the emails send constituent info no content response
func (o *EmailsSendConstituentInfoNoContent) Code() int {
	return 204
}

func (o *EmailsSendConstituentInfoNoContent) Error() string {
	return fmt.Sprintf("[POST /Emails/ConstituentInfo/{constituentId}/Send][%d] emailsSendConstituentInfoNoContent ", 204)
}

func (o *EmailsSendConstituentInfoNoContent) String() string {
	return fmt.Sprintf("[POST /Emails/ConstituentInfo/{constituentId}/Send][%d] emailsSendConstituentInfoNoContent ", 204)
}

func (o *EmailsSendConstituentInfoNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
