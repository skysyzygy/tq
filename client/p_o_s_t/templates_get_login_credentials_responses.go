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

// TemplatesGetLoginCredentialsReader is a Reader for the TemplatesGetLoginCredentials structure.
type TemplatesGetLoginCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatesGetLoginCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplatesGetLoginCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Templates/{templateId}/Login/{loginId}/Credentials] Templates_GetLoginCredentials", response, response.Code())
	}
}

// NewTemplatesGetLoginCredentialsOK creates a TemplatesGetLoginCredentialsOK with default headers values
func NewTemplatesGetLoginCredentialsOK() *TemplatesGetLoginCredentialsOK {
	return &TemplatesGetLoginCredentialsOK{}
}

/*
TemplatesGetLoginCredentialsOK describes a response with status code 200, with default header values.

OK
*/
type TemplatesGetLoginCredentialsOK struct {
	Payload *models.TemplateBody
}

// IsSuccess returns true when this templates get login credentials o k response has a 2xx status code
func (o *TemplatesGetLoginCredentialsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this templates get login credentials o k response has a 3xx status code
func (o *TemplatesGetLoginCredentialsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this templates get login credentials o k response has a 4xx status code
func (o *TemplatesGetLoginCredentialsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this templates get login credentials o k response has a 5xx status code
func (o *TemplatesGetLoginCredentialsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this templates get login credentials o k response a status code equal to that given
func (o *TemplatesGetLoginCredentialsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the templates get login credentials o k response
func (o *TemplatesGetLoginCredentialsOK) Code() int {
	return 200
}

func (o *TemplatesGetLoginCredentialsOK) Error() string {
	return fmt.Sprintf("[POST /Templates/{templateId}/Login/{loginId}/Credentials][%d] templatesGetLoginCredentialsOK  %+v", 200, o.Payload)
}

func (o *TemplatesGetLoginCredentialsOK) String() string {
	return fmt.Sprintf("[POST /Templates/{templateId}/Login/{loginId}/Credentials][%d] templatesGetLoginCredentialsOK  %+v", 200, o.Payload)
}

func (o *TemplatesGetLoginCredentialsOK) GetPayload() *models.TemplateBody {
	return o.Payload
}

func (o *TemplatesGetLoginCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplateBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}