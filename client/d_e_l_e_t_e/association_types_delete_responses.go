// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// AssociationTypesDeleteReader is a Reader for the AssociationTypesDelete structure.
type AssociationTypesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AssociationTypesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewAssociationTypesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/AssociationTypes/{id}] AssociationTypes_Delete", response, response.Code())
	}
}

// NewAssociationTypesDeleteNoContent creates a AssociationTypesDeleteNoContent with default headers values
func NewAssociationTypesDeleteNoContent() *AssociationTypesDeleteNoContent {
	return &AssociationTypesDeleteNoContent{}
}

/*
AssociationTypesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type AssociationTypesDeleteNoContent struct {
}

// IsSuccess returns true when this association types delete no content response has a 2xx status code
func (o *AssociationTypesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this association types delete no content response has a 3xx status code
func (o *AssociationTypesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this association types delete no content response has a 4xx status code
func (o *AssociationTypesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this association types delete no content response has a 5xx status code
func (o *AssociationTypesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this association types delete no content response a status code equal to that given
func (o *AssociationTypesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the association types delete no content response
func (o *AssociationTypesDeleteNoContent) Code() int {
	return 204
}

func (o *AssociationTypesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AssociationTypes/{id}][%d] associationTypesDeleteNoContent ", 204)
}

func (o *AssociationTypesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/AssociationTypes/{id}][%d] associationTypesDeleteNoContent ", 204)
}

func (o *AssociationTypesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
