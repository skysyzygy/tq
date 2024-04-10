// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CartClearCartReader is a Reader for the CartClearCart structure.
type CartClearCartReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartClearCartReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCartClearCartNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /Web/Cart/{sessionKey}] Cart_ClearCart", response, response.Code())
	}
}

// NewCartClearCartNoContent creates a CartClearCartNoContent with default headers values
func NewCartClearCartNoContent() *CartClearCartNoContent {
	return &CartClearCartNoContent{}
}

/*
CartClearCartNoContent describes a response with status code 204, with default header values.

No Content
*/
type CartClearCartNoContent struct {
}

// IsSuccess returns true when this cart clear cart no content response has a 2xx status code
func (o *CartClearCartNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart clear cart no content response has a 3xx status code
func (o *CartClearCartNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart clear cart no content response has a 4xx status code
func (o *CartClearCartNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart clear cart no content response has a 5xx status code
func (o *CartClearCartNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this cart clear cart no content response a status code equal to that given
func (o *CartClearCartNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the cart clear cart no content response
func (o *CartClearCartNoContent) Code() int {
	return 204
}

func (o *CartClearCartNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Web/Cart/{sessionKey}][%d] cartClearCartNoContent ", 204)
}

func (o *CartClearCartNoContent) String() string {
	return fmt.Sprintf("[DELETE /Web/Cart/{sessionKey}][%d] cartClearCartNoContent ", 204)
}

func (o *CartClearCartNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
