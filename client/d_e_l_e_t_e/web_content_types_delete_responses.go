// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WebContentTypesDeleteReader is a Reader for the WebContentTypesDelete structure.
type WebContentTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WebContentTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewWebContentTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/WebContentTypes/{id}] WebContentTypes_Delete", response, response.Code())
	}
}

// NewWebContentTypesDeleteNoContent creates a WebContentTypesDeleteNoContent with default headers values
func NewWebContentTypesDeleteNoContent() *WebContentTypesDeleteNoContent {
	return &WebContentTypesDeleteNoContent{}
}

/*
WebContentTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type WebContentTypesDeleteNoContent struct {
}

// IsSuccess returns true when this web content types delete no content response has a 2xx status code
func (o *WebContentTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this web content types delete no content response has a 3xx status code
func (o *WebContentTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this web content types delete no content response has a 4xx status code
func (o *WebContentTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this web content types delete no content response has a 5xx status code
func (o *WebContentTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this web content types delete no content response a status code equal to that given
func (o *WebContentTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the web content types delete no content response
func (o *WebContentTypesDeleteNoContent) Code() int {
	return 204
}

func (o *WebContentTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/WebContentTypes/{id}][%d] webContentTypesDeleteNoContent ", 204)
}

func (o *WebContentTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/WebContentTypes/{id}][%d] webContentTypesDeleteNoContent ", 204)
}

func (o *WebContentTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
