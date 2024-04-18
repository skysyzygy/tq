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

// TemplatesGetConstituentInfoReader is a Reader for the TemplatesGetConstituentInfo structure.
type TemplatesGetConstituentInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatesGetConstituentInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplatesGetConstituentInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Templates/{templateId}/Constituent/{constituentId}/Info] Templates_GetConstituentInfo", response, response.Code())
	}
}

// NewTemplatesGetConstituentInfoOK creates a TemplatesGetConstituentInfoOK with default headers values
func NewTemplatesGetConstituentInfoOK() *TemplatesGetConstituentInfoOK {
	return &TemplatesGetConstituentInfoOK{}
}

/*
TemplatesGetConstituentInfoOK describes a response with status code 200, with default header values.

OK
*/
type TemplatesGetConstituentInfoOK struct {
	Payload *models.TemplateBody
}

// IsSuccess returns true when this templates get constituent info o k response has a 2xx status code
func (o *TemplatesGetConstituentInfoOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this templates get constituent info o k response has a 3xx status code
func (o *TemplatesGetConstituentInfoOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this templates get constituent info o k response has a 4xx status code
func (o *TemplatesGetConstituentInfoOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this templates get constituent info o k response has a 5xx status code
func (o *TemplatesGetConstituentInfoOK) IsServerError() bool {
	return false
}

// IsCode returns true when this templates get constituent info o k response a status code equal to that given
func (o *TemplatesGetConstituentInfoOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the templates get constituent info o k response
func (o *TemplatesGetConstituentInfoOK) Code() int {
	return 200
}

func (o *TemplatesGetConstituentInfoOK) Error() string {
	return fmt.Sprintf("[POST /Templates/{templateId}/Constituent/{constituentId}/Info][%d] templatesGetConstituentInfoOK  %+v", 200, o.Payload)
}

func (o *TemplatesGetConstituentInfoOK) String() string {
	return fmt.Sprintf("[POST /Templates/{templateId}/Constituent/{constituentId}/Info][%d] templatesGetConstituentInfoOK  %+v", 200, o.Payload)
}

func (o *TemplatesGetConstituentInfoOK) GetPayload() *models.TemplateBody {
	return o.Payload
}

func (o *TemplatesGetConstituentInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}