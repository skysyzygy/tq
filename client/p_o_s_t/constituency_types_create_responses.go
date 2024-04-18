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

// ConstituencyTypesCreateReader is a Reader for the ConstituencyTypesCreate structure.
type ConstituencyTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituencyTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituencyTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/ConstituencyTypes] ConstituencyTypes_Create", response, response.Code())
	}
}

// NewConstituencyTypesCreateOK creates a ConstituencyTypesCreateOK with default headers values
func NewConstituencyTypesCreateOK() *ConstituencyTypesCreateOK {
	return &ConstituencyTypesCreateOK{}
}

/*
ConstituencyTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type ConstituencyTypesCreateOK struct {
	Payload *models.ConstituencyType
}

// IsSuccess returns true when this constituency types create o k response has a 2xx status code
func (o *ConstituencyTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituency types create o k response has a 3xx status code
func (o *ConstituencyTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituency types create o k response has a 4xx status code
func (o *ConstituencyTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituency types create o k response has a 5xx status code
func (o *ConstituencyTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituency types create o k response a status code equal to that given
func (o *ConstituencyTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituency types create o k response
func (o *ConstituencyTypesCreateOK) Code() int {
	return 200
}

func (o *ConstituencyTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/ConstituencyTypes][%d] constituencyTypesCreateOK  %+v", 200, o.Payload)
}

func (o *ConstituencyTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/ConstituencyTypes][%d] constituencyTypesCreateOK  %+v", 200, o.Payload)
}

func (o *ConstituencyTypesCreateOK) GetPayload() *models.ConstituencyType {
	return o.Payload
}

func (o *ConstituencyTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituencyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}