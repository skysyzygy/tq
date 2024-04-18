// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ConstituentTypesDeleteReader is a Reader for the ConstituentTypesDelete structure.
type ConstituentTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewConstituentTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/ConstituentTypes/{id}] ConstituentTypes_Delete", response, response.Code())
	}
}

// NewConstituentTypesDeleteNoContent creates a ConstituentTypesDeleteNoContent with default headers values
func NewConstituentTypesDeleteNoContent() *ConstituentTypesDeleteNoContent {
	return &ConstituentTypesDeleteNoContent{}
}

/*
ConstituentTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ConstituentTypesDeleteNoContent struct {
}

// IsSuccess returns true when this constituent types delete no content response has a 2xx status code
func (o *ConstituentTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent types delete no content response has a 3xx status code
func (o *ConstituentTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent types delete no content response has a 4xx status code
func (o *ConstituentTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent types delete no content response has a 5xx status code
func (o *ConstituentTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent types delete no content response a status code equal to that given
func (o *ConstituentTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the constituent types delete no content response
func (o *ConstituentTypesDeleteNoContent) Code() int {
	return 204
}

func (o *ConstituentTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentTypes/{id}][%d] constituentTypesDeleteNoContent ", 204)
}

func (o *ConstituentTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ConstituentTypes/{id}][%d] constituentTypesDeleteNoContent ", 204)
}

func (o *ConstituentTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}