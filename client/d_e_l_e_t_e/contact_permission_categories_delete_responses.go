// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ContactPermissionCategoriesDeleteReader is a Reader for the ContactPermissionCategoriesDelete structure.
type ContactPermissionCategoriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPermissionCategoriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewContactPermissionCategoriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/ContactPermissionCategories/{id}] ContactPermissionCategories_Delete", response, response.Code())
	}
}

// NewContactPermissionCategoriesDeleteNoContent creates a ContactPermissionCategoriesDeleteNoContent with default headers values
func NewContactPermissionCategoriesDeleteNoContent() *ContactPermissionCategoriesDeleteNoContent {
	return &ContactPermissionCategoriesDeleteNoContent{}
}

/*
ContactPermissionCategoriesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ContactPermissionCategoriesDeleteNoContent struct {
}

// IsSuccess returns true when this contact permission categories delete no content response has a 2xx status code
func (o *ContactPermissionCategoriesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact permission categories delete no content response has a 3xx status code
func (o *ContactPermissionCategoriesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact permission categories delete no content response has a 4xx status code
func (o *ContactPermissionCategoriesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact permission categories delete no content response has a 5xx status code
func (o *ContactPermissionCategoriesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this contact permission categories delete no content response a status code equal to that given
func (o *ContactPermissionCategoriesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the contact permission categories delete no content response
func (o *ContactPermissionCategoriesDeleteNoContent) Code() int {
	return 204
}

func (o *ContactPermissionCategoriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ContactPermissionCategories/{id}][%d] contactPermissionCategoriesDeleteNoContent ", 204)
}

func (o *ContactPermissionCategoriesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ContactPermissionCategories/{id}][%d] contactPermissionCategoriesDeleteNoContent ", 204)
}

func (o *ContactPermissionCategoriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
