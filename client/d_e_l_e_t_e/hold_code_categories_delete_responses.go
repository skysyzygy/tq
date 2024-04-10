// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// HoldCodeCategoriesDeleteReader is a Reader for the HoldCodeCategoriesDelete structure.
type HoldCodeCategoriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *HoldCodeCategoriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewHoldCodeCategoriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/HoldCodeCategories/{id}] HoldCodeCategories_Delete", response, response.Code())
	}
}

// NewHoldCodeCategoriesDeleteNoContent creates a HoldCodeCategoriesDeleteNoContent with default headers values
func NewHoldCodeCategoriesDeleteNoContent() *HoldCodeCategoriesDeleteNoContent {
	return &HoldCodeCategoriesDeleteNoContent{}
}

/*
HoldCodeCategoriesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type HoldCodeCategoriesDeleteNoContent struct {
}

// IsSuccess returns true when this hold code categories delete no content response has a 2xx status code
func (o *HoldCodeCategoriesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this hold code categories delete no content response has a 3xx status code
func (o *HoldCodeCategoriesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this hold code categories delete no content response has a 4xx status code
func (o *HoldCodeCategoriesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this hold code categories delete no content response has a 5xx status code
func (o *HoldCodeCategoriesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this hold code categories delete no content response a status code equal to that given
func (o *HoldCodeCategoriesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the hold code categories delete no content response
func (o *HoldCodeCategoriesDeleteNoContent) Code() int {
	return 204
}

func (o *HoldCodeCategoriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/HoldCodeCategories/{id}][%d] holdCodeCategoriesDeleteNoContent ", 204)
}

func (o *HoldCodeCategoriesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/HoldCodeCategories/{id}][%d] holdCodeCategoriesDeleteNoContent ", 204)
}

func (o *HoldCodeCategoriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
