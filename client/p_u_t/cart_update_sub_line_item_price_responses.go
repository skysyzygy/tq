// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// CartUpdateSubLineItemPriceReader is a Reader for the CartUpdateSubLineItemPrice structure.
type CartUpdateSubLineItemPriceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartUpdateSubLineItemPriceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartUpdateSubLineItemPriceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /Web/Cart/{sessionKey}/SubLineItems/{subLineItemId}/Price] Cart_UpdateSubLineItemPrice", response, response.Code())
	}
}

// NewCartUpdateSubLineItemPriceOK creates a CartUpdateSubLineItemPriceOK with default headers values
func NewCartUpdateSubLineItemPriceOK() *CartUpdateSubLineItemPriceOK {
	return &CartUpdateSubLineItemPriceOK{}
}

/*
CartUpdateSubLineItemPriceOK describes a response with status code 200, with default header values.

OK
*/
type CartUpdateSubLineItemPriceOK struct {
	Payload *models.UpdatePriceResponse
}

// IsSuccess returns true when this cart update sub line item price o k response has a 2xx status code
func (o *CartUpdateSubLineItemPriceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart update sub line item price o k response has a 3xx status code
func (o *CartUpdateSubLineItemPriceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart update sub line item price o k response has a 4xx status code
func (o *CartUpdateSubLineItemPriceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart update sub line item price o k response has a 5xx status code
func (o *CartUpdateSubLineItemPriceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart update sub line item price o k response a status code equal to that given
func (o *CartUpdateSubLineItemPriceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart update sub line item price o k response
func (o *CartUpdateSubLineItemPriceOK) Code() int {
	return 200
}

func (o *CartUpdateSubLineItemPriceOK) Error() string {
	return fmt.Sprintf("[PUT /Web/Cart/{sessionKey}/SubLineItems/{subLineItemId}/Price][%d] cartUpdateSubLineItemPriceOK  %+v", 200, o.Payload)
}

func (o *CartUpdateSubLineItemPriceOK) String() string {
	return fmt.Sprintf("[PUT /Web/Cart/{sessionKey}/SubLineItems/{subLineItemId}/Price][%d] cartUpdateSubLineItemPriceOK  %+v", 200, o.Payload)
}

func (o *CartUpdateSubLineItemPriceOK) GetPayload() *models.UpdatePriceResponse {
	return o.Payload
}

func (o *CartUpdateSubLineItemPriceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UpdatePriceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
