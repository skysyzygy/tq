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

// ProgramListingsGetReader is a Reader for the ProgramListingsGet structure.
type ProgramListingsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProgramListingsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProgramListingsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProgramListingsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProgramListingsGetOK creates a ProgramListingsGetOK with default headers values
func NewProgramListingsGetOK() *ProgramListingsGetOK {
	return &ProgramListingsGetOK{}
}

/*
ProgramListingsGetOK describes a response with status code 200, with default header values.

OK
*/
type ProgramListingsGetOK struct {
	Payload *models.ProgramListing
}

// IsSuccess returns true when this program listings get o k response has a 2xx status code
func (o *ProgramListingsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this program listings get o k response has a 3xx status code
func (o *ProgramListingsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this program listings get o k response has a 4xx status code
func (o *ProgramListingsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this program listings get o k response has a 5xx status code
func (o *ProgramListingsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this program listings get o k response a status code equal to that given
func (o *ProgramListingsGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the program listings get o k response
func (o *ProgramListingsGetOK) Code() int {
	return 200
}

func (o *ProgramListingsGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/ProgramListings/{programListingId}][%d] programListingsGetOK %s", 200, payload)
}

func (o *ProgramListingsGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/ProgramListings/{programListingId}][%d] programListingsGetOK %s", 200, payload)
}

func (o *ProgramListingsGetOK) GetPayload() *models.ProgramListing {
	return o.Payload
}

func (o *ProgramListingsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProgramListing)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProgramListingsGetDefault creates a ProgramListingsGetDefault with default headers values
func NewProgramListingsGetDefault(code int) *ProgramListingsGetDefault {
	return &ProgramListingsGetDefault{
		_statusCode: code,
	}
}

/*
ProgramListingsGetDefault describes a response with status code -1, with default header values.

Error
*/
type ProgramListingsGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this program listings get default response has a 2xx status code
func (o *ProgramListingsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this program listings get default response has a 3xx status code
func (o *ProgramListingsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this program listings get default response has a 4xx status code
func (o *ProgramListingsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this program listings get default response has a 5xx status code
func (o *ProgramListingsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this program listings get default response a status code equal to that given
func (o *ProgramListingsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the program listings get default response
func (o *ProgramListingsGetDefault) Code() int {
	return o._statusCode
}

func (o *ProgramListingsGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/ProgramListings/{programListingId}][%d] ProgramListings_Get default %s", o._statusCode, payload)
}

func (o *ProgramListingsGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /CRM/ProgramListings/{programListingId}][%d] ProgramListings_Get default %s", o._statusCode, payload)
}

func (o *ProgramListingsGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *ProgramListingsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
