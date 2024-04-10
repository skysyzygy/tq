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

// TheatersCreateReader is a Reader for the TheatersCreate structure.
type TheatersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TheatersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTheatersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/Theaters] Theaters_Create", response, response.Code())
	}
}

// NewTheatersCreateOK creates a TheatersCreateOK with default headers values
func NewTheatersCreateOK() *TheatersCreateOK {
	return &TheatersCreateOK{}
}

/*
TheatersCreateOK describes a response with status code 200, with default header values.

OK
*/
type TheatersCreateOK struct {
	Payload *models.Theater
}

// IsSuccess returns true when this theaters create o k response has a 2xx status code
func (o *TheatersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this theaters create o k response has a 3xx status code
func (o *TheatersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this theaters create o k response has a 4xx status code
func (o *TheatersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this theaters create o k response has a 5xx status code
func (o *TheatersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this theaters create o k response a status code equal to that given
func (o *TheatersCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the theaters create o k response
func (o *TheatersCreateOK) Code() int {
	return 200
}

func (o *TheatersCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/Theaters][%d] theatersCreateOK  %+v", 200, o.Payload)
}

func (o *TheatersCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/Theaters][%d] theatersCreateOK  %+v", 200, o.Payload)
}

func (o *TheatersCreateOK) GetPayload() *models.Theater {
	return o.Payload
}

func (o *TheatersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Theater)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
