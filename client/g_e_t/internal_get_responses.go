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

// InternalGetReader is a Reader for the InternalGet structure.
type InternalGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInternalGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Internal/AddressDetails/{addressId}] Internal_Get", response, response.Code())
	}
}

// NewInternalGetOK creates a InternalGetOK with default headers values
func NewInternalGetOK() *InternalGetOK {
	return &InternalGetOK{}
}

/*
InternalGetOK describes a response with status code 200, with default header values.

OK
*/
type InternalGetOK struct {
	Payload *models.AddressDetail
}

// IsSuccess returns true when this internal get o k response has a 2xx status code
func (o *InternalGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal get o k response has a 3xx status code
func (o *InternalGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal get o k response has a 4xx status code
func (o *InternalGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal get o k response has a 5xx status code
func (o *InternalGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this internal get o k response a status code equal to that given
func (o *InternalGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the internal get o k response
func (o *InternalGetOK) Code() int {
	return 200
}

func (o *InternalGetOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Internal/AddressDetails/{addressId}][%d] internalGetOK  %+v", 200, o.Payload)
}

func (o *InternalGetOK) String() string {
	return fmt.Sprintf("[GET /CRM/Internal/AddressDetails/{addressId}][%d] internalGetOK  %+v", 200, o.Payload)
}

func (o *InternalGetOK) GetPayload() *models.AddressDetail {
	return o.Payload
}

func (o *InternalGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddressDetail)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
