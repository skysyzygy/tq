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

// SalutationsCreateReader is a Reader for the SalutationsCreate structure.
type SalutationsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SalutationsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSalutationsCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /CRM/Salutations] Salutations_Create", response, response.Code())
	}
}

// NewSalutationsCreateOK creates a SalutationsCreateOK with default headers values
func NewSalutationsCreateOK() *SalutationsCreateOK {
	return &SalutationsCreateOK{}
}

/*
SalutationsCreateOK describes a response with status code 200, with default header values.

OK
*/
type SalutationsCreateOK struct {
	Payload *models.Salutation
}

// IsSuccess returns true when this salutations create o k response has a 2xx status code
func (o *SalutationsCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this salutations create o k response has a 3xx status code
func (o *SalutationsCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this salutations create o k response has a 4xx status code
func (o *SalutationsCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this salutations create o k response has a 5xx status code
func (o *SalutationsCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this salutations create o k response a status code equal to that given
func (o *SalutationsCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the salutations create o k response
func (o *SalutationsCreateOK) Code() int {
	return 200
}

func (o *SalutationsCreateOK) Error() string {
	return fmt.Sprintf("[POST /CRM/Salutations][%d] salutationsCreateOK  %+v", 200, o.Payload)
}

func (o *SalutationsCreateOK) String() string {
	return fmt.Sprintf("[POST /CRM/Salutations][%d] salutationsCreateOK  %+v", 200, o.Payload)
}

func (o *SalutationsCreateOK) GetPayload() *models.Salutation {
	return o.Payload
}

func (o *SalutationsCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Salutation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
