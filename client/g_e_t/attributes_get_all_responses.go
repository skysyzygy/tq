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

// AttributesGetAllReader is a Reader for the AttributesGetAll structure.
type AttributesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AttributesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAttributesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Attributes] Attributes_GetAll", response, response.Code())
	}
}

// NewAttributesGetAllOK creates a AttributesGetAllOK with default headers values
func NewAttributesGetAllOK() *AttributesGetAllOK {
	return &AttributesGetAllOK{}
}

/*
AttributesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type AttributesGetAllOK struct {
	Payload []*models.Attribute
}

// IsSuccess returns true when this attributes get all o k response has a 2xx status code
func (o *AttributesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this attributes get all o k response has a 3xx status code
func (o *AttributesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this attributes get all o k response has a 4xx status code
func (o *AttributesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this attributes get all o k response has a 5xx status code
func (o *AttributesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this attributes get all o k response a status code equal to that given
func (o *AttributesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the attributes get all o k response
func (o *AttributesGetAllOK) Code() int {
	return 200
}

func (o *AttributesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Attributes][%d] attributesGetAllOK  %+v", 200, o.Payload)
}

func (o *AttributesGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/Attributes][%d] attributesGetAllOK  %+v", 200, o.Payload)
}

func (o *AttributesGetAllOK) GetPayload() []*models.Attribute {
	return o.Payload
}

func (o *AttributesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}