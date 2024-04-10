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

// WebContentTypesGetReader is a Reader for the WebContentTypesGet structure.
type WebContentTypesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebContentTypesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWebContentTypesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/WebContentTypes/{id}] WebContentTypes_Get", response, response.Code())
	}
}

// NewWebContentTypesGetOK creates a WebContentTypesGetOK with default headers values
func NewWebContentTypesGetOK() *WebContentTypesGetOK {
	return &WebContentTypesGetOK{}
}

/*
WebContentTypesGetOK describes a response with status code 200, with default header values.

OK
*/
type WebContentTypesGetOK struct {
	Payload *models.WebContentType
}

// IsSuccess returns true when this web content types get o k response has a 2xx status code
func (o *WebContentTypesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this web content types get o k response has a 3xx status code
func (o *WebContentTypesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this web content types get o k response has a 4xx status code
func (o *WebContentTypesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this web content types get o k response has a 5xx status code
func (o *WebContentTypesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this web content types get o k response a status code equal to that given
func (o *WebContentTypesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the web content types get o k response
func (o *WebContentTypesGetOK) Code() int {
	return 200
}

func (o *WebContentTypesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/WebContentTypes/{id}][%d] webContentTypesGetOK  %+v", 200, o.Payload)
}

func (o *WebContentTypesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/WebContentTypes/{id}][%d] webContentTypesGetOK  %+v", 200, o.Payload)
}

func (o *WebContentTypesGetOK) GetPayload() *models.WebContentType {
	return o.Payload
}

func (o *WebContentTypesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebContentType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
