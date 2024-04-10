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

// PhoneTypesGetAllReader is a Reader for the PhoneTypesGetAll structure.
type PhoneTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PhoneTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPhoneTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PhoneTypes] PhoneTypes_GetAll", response, response.Code())
	}
}

// NewPhoneTypesGetAllOK creates a PhoneTypesGetAllOK with default headers values
func NewPhoneTypesGetAllOK() *PhoneTypesGetAllOK {
	return &PhoneTypesGetAllOK{}
}

/*
PhoneTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PhoneTypesGetAllOK struct {
	Payload []*models.PhoneType
}

// IsSuccess returns true when this phone types get all o k response has a 2xx status code
func (o *PhoneTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this phone types get all o k response has a 3xx status code
func (o *PhoneTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this phone types get all o k response has a 4xx status code
func (o *PhoneTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this phone types get all o k response has a 5xx status code
func (o *PhoneTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this phone types get all o k response a status code equal to that given
func (o *PhoneTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the phone types get all o k response
func (o *PhoneTypesGetAllOK) Code() int {
	return 200
}

func (o *PhoneTypesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PhoneTypes][%d] phoneTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *PhoneTypesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PhoneTypes][%d] phoneTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *PhoneTypesGetAllOK) GetPayload() []*models.PhoneType {
	return o.Payload
}

func (o *PhoneTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
