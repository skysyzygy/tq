// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PriceCategoriesDeleteReader is a Reader for the PriceCategoriesDelete structure.
type PriceCategoriesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceCategoriesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPriceCategoriesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /ReferenceData/PriceCategories/{id}] PriceCategories_Delete", response, response.Code())
	}
}

// NewPriceCategoriesDeleteNoContent creates a PriceCategoriesDeleteNoContent with default headers values
func NewPriceCategoriesDeleteNoContent() *PriceCategoriesDeleteNoContent {
	return &PriceCategoriesDeleteNoContent{}
}

/*
PriceCategoriesDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type PriceCategoriesDeleteNoContent struct {
}

// IsSuccess returns true when this price categories delete no content response has a 2xx status code
func (o *PriceCategoriesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price categories delete no content response has a 3xx status code
func (o *PriceCategoriesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price categories delete no content response has a 4xx status code
func (o *PriceCategoriesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this price categories delete no content response has a 5xx status code
func (o *PriceCategoriesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this price categories delete no content response a status code equal to that given
func (o *PriceCategoriesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the price categories delete no content response
func (o *PriceCategoriesDeleteNoContent) Code() int {
	return 204
}

func (o *PriceCategoriesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ReferenceData/PriceCategories/{id}][%d] priceCategoriesDeleteNoContent ", 204)
}

func (o *PriceCategoriesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ReferenceData/PriceCategories/{id}][%d] priceCategoriesDeleteNoContent ", 204)
}

func (o *PriceCategoriesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
