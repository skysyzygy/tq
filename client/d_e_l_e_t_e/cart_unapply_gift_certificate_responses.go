// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// CartUnapplyGiftCertificateReader is a Reader for the CartUnapplyGiftCertificate structure.
type CartUnapplyGiftCertificateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartUnapplyGiftCertificateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewCartUnapplyGiftCertificateNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /Web/Cart/{sessionKey}/Payments/GiftCertificate/{giftCertificateNumber}] Cart_UnapplyGiftCertificate", response, response.Code())
	}
}

// NewCartUnapplyGiftCertificateNoContent creates a CartUnapplyGiftCertificateNoContent with default headers values
func NewCartUnapplyGiftCertificateNoContent() *CartUnapplyGiftCertificateNoContent {
	return &CartUnapplyGiftCertificateNoContent{}
}

/*
CartUnapplyGiftCertificateNoContent describes a response with status code 204, with default header values.

No Content
*/
type CartUnapplyGiftCertificateNoContent struct {
}

// IsSuccess returns true when this cart unapply gift certificate no content response has a 2xx status code
func (o *CartUnapplyGiftCertificateNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart unapply gift certificate no content response has a 3xx status code
func (o *CartUnapplyGiftCertificateNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart unapply gift certificate no content response has a 4xx status code
func (o *CartUnapplyGiftCertificateNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart unapply gift certificate no content response has a 5xx status code
func (o *CartUnapplyGiftCertificateNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this cart unapply gift certificate no content response a status code equal to that given
func (o *CartUnapplyGiftCertificateNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the cart unapply gift certificate no content response
func (o *CartUnapplyGiftCertificateNoContent) Code() int {
	return 204
}

func (o *CartUnapplyGiftCertificateNoContent) Error() string {
	return fmt.Sprintf("[DELETE /Web/Cart/{sessionKey}/Payments/GiftCertificate/{giftCertificateNumber}][%d] cartUnapplyGiftCertificateNoContent ", 204)
}

func (o *CartUnapplyGiftCertificateNoContent) String() string {
	return fmt.Sprintf("[DELETE /Web/Cart/{sessionKey}/Payments/GiftCertificate/{giftCertificateNumber}][%d] cartUnapplyGiftCertificateNoContent ", 204)
}

func (o *CartUnapplyGiftCertificateNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}