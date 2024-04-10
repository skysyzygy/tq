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

// TemplateTypesGetReader is a Reader for the TemplateTypesGet structure.
type TemplateTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplateTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplateTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/TemplateTypes/{id}] TemplateTypes_Get", response, response.Code())
	}
}

// NewTemplateTypesGetOK creates a TemplateTypesGetOK with default headers values
func NewTemplateTypesGetOK() *TemplateTypesGetOK {
	return &TemplateTypesGetOK{}
}

/*
TemplateTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type TemplateTypesGetOK struct {
	Payload *models.TemplateType
}

// IsSuccess returns true when this template types get o k response has a 2xx status code
func (o *TemplateTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template types get o k response has a 3xx status code
func (o *TemplateTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template types get o k response has a 4xx status code
func (o *TemplateTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template types get o k response has a 5xx status code
func (o *TemplateTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template types get o k response a status code equal to that given
func (o *TemplateTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template types get o k response
func (o *TemplateTypesGetOK) Code() int {
	return 200
}

func (o *TemplateTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/TemplateTypes/{id}][%d] templateTypesGetOK  %+v", 200, o.Payload)
}

func (o *TemplateTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/TemplateTypes/{id}][%d] templateTypesGetOK  %+v", 200, o.Payload)
}

func (o *TemplateTypesGetOK) GetPayload() *models.TemplateType {
	return o.Payload
}

func (o *TemplateTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
