// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CustomCreateReader is a Reader for the CustomCreate structure.
type CustomCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Custom/{resourceName}] Custom_Create", response, response.Code())
	}
}

// NewCustomCreateOK creates a CustomCreateOK with default headers values
func NewCustomCreateOK() *CustomCreateOK {
	return &CustomCreateOK{}
}

/*
CustomCreateOK describes a response with status code 200, with default header values.

OK
*/
type CustomCreateOK struct {
	Payload interface{}
}

// IsSuccess returns true when this custom create o k response has a 2xx status code
func (o *CustomCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom create o k response has a 3xx status code
func (o *CustomCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom create o k response has a 4xx status code
func (o *CustomCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom create o k response has a 5xx status code
func (o *CustomCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom create o k response a status code equal to that given
func (o *CustomCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom create o k response
func (o *CustomCreateOK) Code() int {
	return 200
}

func (o *CustomCreateOK) Error() string {
	return fmt.Sprintf("[POST /Custom/{resourceName}][%d] customCreateOK  %+v", 200, o.Payload)
}

func (o *CustomCreateOK) String() string {
	return fmt.Sprintf("[POST /Custom/{resourceName}][%d] customCreateOK  %+v", 200, o.Payload)
}

func (o *CustomCreateOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
