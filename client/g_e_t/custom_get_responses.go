// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CustomGetReader is a Reader for the CustomGet structure.
type CustomGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Custom/{resourceName}/{id}] Custom_Get", response, response.Code())
	}
}

// NewCustomGetOK creates a CustomGetOK with default headers values
func NewCustomGetOK() *CustomGetOK {
	return &CustomGetOK{}
}

/*
CustomGetOK describes a response with status code 200, with default header values.

OK
*/
type CustomGetOK struct {
	Payload interface{}
}

// IsSuccess returns true when this custom get o k response has a 2xx status code
func (o *CustomGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom get o k response has a 3xx status code
func (o *CustomGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom get o k response has a 4xx status code
func (o *CustomGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom get o k response has a 5xx status code
func (o *CustomGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom get o k response a status code equal to that given
func (o *CustomGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom get o k response
func (o *CustomGetOK) Code() int {
	return 200
}

func (o *CustomGetOK) Error() string {
	return fmt.Sprintf("[GET /Custom/{resourceName}/{id}][%d] customGetOK  %+v", 200, o.Payload)
}

func (o *CustomGetOK) String() string {
	return fmt.Sprintf("[GET /Custom/{resourceName}/{id}][%d] customGetOK  %+v", 200, o.Payload)
}

func (o *CustomGetOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CustomGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
