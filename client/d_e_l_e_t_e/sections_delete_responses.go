// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// SectionsDeleteReader is a Reader for the SectionsDelete structure.
type SectionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SectionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewSectionsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/Sections/{id}] Sections_Delete", response, response.Code())
	}
}

// NewSectionsDeleteNoContent creates a SectionsDeleteNoContent with default headers values
func NewSectionsDeleteNoContent() *SectionsDeleteNoContent {
	return &SectionsDeleteNoContent{}
}

/*
SectionsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type SectionsDeleteNoContent struct {
}

// IsSuccess returns true when this sections delete no content response has a 2xx status code
func (o *SectionsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sections delete no content response has a 3xx status code
func (o *SectionsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sections delete no content response has a 4xx status code
func (o *SectionsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this sections delete no content response has a 5xx status code
func (o *SectionsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this sections delete no content response a status code equal to that given
func (o *SectionsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the sections delete no content response
func (o *SectionsDeleteNoContent) Code() int {
	return 204
}

func (o *SectionsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/Sections/{id}][%d] sectionsDeleteNoContent ", 204)
}

func (o *SectionsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/Sections/{id}][%d] sectionsDeleteNoContent ", 204)
}

func (o *SectionsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
