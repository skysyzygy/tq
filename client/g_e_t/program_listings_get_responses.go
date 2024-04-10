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
		return nil, runtime.NewAPIError("[GET /CRM/ProgramListings/{programListingId}] ProgramListings_Get", response, response.Code())
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
	return fmt.Sprintf("[GET /CRM/ProgramListings/{programListingId}][%d] programListingsGetOK  %+v", 200, o.Payload)
}

func (o *ProgramListingsGetOK) String() string {
	return fmt.Sprintf("[GET /CRM/ProgramListings/{programListingId}][%d] programListingsGetOK  %+v", 200, o.Payload)
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
