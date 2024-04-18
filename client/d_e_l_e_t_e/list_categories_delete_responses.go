// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// ListCategoriesDeleteReader is a Reader for the ListCategoriesDelete structure.
type ListCategoriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCategoriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewListCategoriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/ListCategories/{id}] ListCategories_Delete", response, response.Code())
	}
}

// NewListCategoriesDeleteNoContent creates a ListCategoriesDeleteNoContent with default headers values
func NewListCategoriesDeleteNoContent() *ListCategoriesDeleteNoContent {
	return &ListCategoriesDeleteNoContent{}
}

/*
ListCategoriesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type ListCategoriesDeleteNoContent struct {
}

// IsSuccess returns true when this list categories delete no content response has a 2xx status code
func (o *ListCategoriesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this list categories delete no content response has a 3xx status code
func (o *ListCategoriesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this list categories delete no content response has a 4xx status code
func (o *ListCategoriesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this list categories delete no content response has a 5xx status code
func (o *ListCategoriesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this list categories delete no content response a status code equal to that given
func (o *ListCategoriesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the list categories delete no content response
func (o *ListCategoriesDeleteNoContent) Code() int {
	return 204
}

func (o *ListCategoriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ListCategories/{id}][%d] listCategoriesDeleteNoContent ", 204)
}

func (o *ListCategoriesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/ListCategories/{id}][%d] listCategoriesDeleteNoContent ", 204)
}

func (o *ListCategoriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}