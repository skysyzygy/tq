// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CartRemovePaymentReader is a Reader for the CartRemovePayment structure.
type CartRemovePaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartRemovePaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCartRemovePaymentNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /Web/Cart/{sessionKey}/Payments/{paymentId}] Cart_RemovePayment", response, response.Code())
	}
}

// NewCartRemovePaymentNoContent creates a CartRemovePaymentNoContent with default headers values
func NewCartRemovePaymentNoContent() *CartRemovePaymentNoContent {
	return &CartRemovePaymentNoContent{}
}

/*
CartRemovePaymentNoContent describes a response with status code 204, with default header values.

No Content
*/
type CartRemovePaymentNoContent struct {
}

// IsSuccess returns true when this cart remove payment no content response has a 2xx status code
func (o *CartRemovePaymentNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart remove payment no content response has a 3xx status code
func (o *CartRemovePaymentNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart remove payment no content response has a 4xx status code
func (o *CartRemovePaymentNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart remove payment no content response has a 5xx status code
func (o *CartRemovePaymentNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this cart remove payment no content response a status code equal to that given
func (o *CartRemovePaymentNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the cart remove payment no content response
func (o *CartRemovePaymentNoContent) Code() int {
	return 204
}

func (o *CartRemovePaymentNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Web/Cart/{sessionKey}/Payments/{paymentId}][%d] cartRemovePaymentNoContent ", 204)
}

func (o *CartRemovePaymentNoContent) String() string {
	return fmt.Sprintf("[DELETE /Web/Cart/{sessionKey}/Payments/{paymentId}][%d] cartRemovePaymentNoContent ", 204)
}

func (o *CartRemovePaymentNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
