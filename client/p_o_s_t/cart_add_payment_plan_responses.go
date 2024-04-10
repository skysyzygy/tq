// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// CartAddPaymentPlanReader is a Reader for the CartAddPaymentPlan structure.
type CartAddPaymentPlanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartAddPaymentPlanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartAddPaymentPlanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Web/Cart/{sessionKey}/Payments/Plan/NumberOfPayments] Cart_AddPaymentPlan", response, response.Code())
	}
}

// NewCartAddPaymentPlanOK creates a CartAddPaymentPlanOK with default headers values
func NewCartAddPaymentPlanOK() *CartAddPaymentPlanOK {
	return &CartAddPaymentPlanOK{}
}

/*
CartAddPaymentPlanOK describes a response with status code 200, with default header values.

OK
*/
type CartAddPaymentPlanOK struct {
	Payload []*models.PaymentPlan
}

// IsSuccess returns true when this cart add payment plan o k response has a 2xx status code
func (o *CartAddPaymentPlanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart add payment plan o k response has a 3xx status code
func (o *CartAddPaymentPlanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart add payment plan o k response has a 4xx status code
func (o *CartAddPaymentPlanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart add payment plan o k response has a 5xx status code
func (o *CartAddPaymentPlanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart add payment plan o k response a status code equal to that given
func (o *CartAddPaymentPlanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart add payment plan o k response
func (o *CartAddPaymentPlanOK) Code() int {
	return 200
}

func (o *CartAddPaymentPlanOK) Error() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Payments/Plan/NumberOfPayments][%d] cartAddPaymentPlanOK  %+v", 200, o.Payload)
}

func (o *CartAddPaymentPlanOK) String() string {
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Payments/Plan/NumberOfPayments][%d] cartAddPaymentPlanOK  %+v", 200, o.Payload)
}

func (o *CartAddPaymentPlanOK) GetPayload() []*models.PaymentPlan {
	return o.Payload
}

func (o *CartAddPaymentPlanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
