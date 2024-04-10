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

// ConstituencyTypesGetReader is a Reader for the ConstituencyTypesGet structure.
type ConstituencyTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituencyTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituencyTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ConstituencyTypes/{id}] ConstituencyTypes_Get", response, response.Code())
	}
}

// NewConstituencyTypesGetOK creates a ConstituencyTypesGetOK with default headers values
func NewConstituencyTypesGetOK() *ConstituencyTypesGetOK {
	return &ConstituencyTypesGetOK{}
}

/*
ConstituencyTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type ConstituencyTypesGetOK struct {
	Payload *models.ConstituencyType
}

// IsSuccess returns true when this constituency types get o k response has a 2xx status code
func (o *ConstituencyTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituency types get o k response has a 3xx status code
func (o *ConstituencyTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituency types get o k response has a 4xx status code
func (o *ConstituencyTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituency types get o k response has a 5xx status code
func (o *ConstituencyTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituency types get o k response a status code equal to that given
func (o *ConstituencyTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituency types get o k response
func (o *ConstituencyTypesGetOK) Code() int {
	return 200
}

func (o *ConstituencyTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ConstituencyTypes/{id}][%d] constituencyTypesGetOK  %+v", 200, o.Payload)
}

func (o *ConstituencyTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ConstituencyTypes/{id}][%d] constituencyTypesGetOK  %+v", 200, o.Payload)
}

func (o *ConstituencyTypesGetOK) GetPayload() *models.ConstituencyType {
	return o.Payload
}

func (o *ConstituencyTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConstituencyType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
