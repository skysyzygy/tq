// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// CartApplyOtherPaymentReader is a Reader for the CartApplyOtherPayment structure.
type CartApplyOtherPaymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartApplyOtherPaymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartApplyOtherPaymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCartApplyOtherPaymentDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCartApplyOtherPaymentOK creates a CartApplyOtherPaymentOK with default headers values
func NewCartApplyOtherPaymentOK() *CartApplyOtherPaymentOK {
	return &CartApplyOtherPaymentOK{}
}

/*
CartApplyOtherPaymentOK describes a response with status code 200, with default header values.

OK
*/
type CartApplyOtherPaymentOK struct {
	Payload *models.ApplyPaymentResponse
}

// IsSuccess returns true when this cart apply other payment o k response has a 2xx status code
func (o *CartApplyOtherPaymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart apply other payment o k response has a 3xx status code
func (o *CartApplyOtherPaymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart apply other payment o k response has a 4xx status code
func (o *CartApplyOtherPaymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart apply other payment o k response has a 5xx status code
func (o *CartApplyOtherPaymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart apply other payment o k response a status code equal to that given
func (o *CartApplyOtherPaymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart apply other payment o k response
func (o *CartApplyOtherPaymentOK) Code() int {
	return 200
}

func (o *CartApplyOtherPaymentOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Payments/Other][%d] cartApplyOtherPaymentOK %s", 200, payload)
}

func (o *CartApplyOtherPaymentOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Payments/Other][%d] cartApplyOtherPaymentOK %s", 200, payload)
}

func (o *CartApplyOtherPaymentOK) GetPayload() *models.ApplyPaymentResponse {
	return o.Payload
}

func (o *CartApplyOtherPaymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ApplyPaymentResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCartApplyOtherPaymentDefault creates a CartApplyOtherPaymentDefault with default headers values
func NewCartApplyOtherPaymentDefault(code int) *CartApplyOtherPaymentDefault {
	return &CartApplyOtherPaymentDefault{
		_statusCode: code,
	}
}

/*
CartApplyOtherPaymentDefault describes a response with status code -1, with default header values.

Error
*/
type CartApplyOtherPaymentDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this cart apply other payment default response has a 2xx status code
func (o *CartApplyOtherPaymentDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cart apply other payment default response has a 3xx status code
func (o *CartApplyOtherPaymentDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cart apply other payment default response has a 4xx status code
func (o *CartApplyOtherPaymentDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cart apply other payment default response has a 5xx status code
func (o *CartApplyOtherPaymentDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cart apply other payment default response a status code equal to that given
func (o *CartApplyOtherPaymentDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cart apply other payment default response
func (o *CartApplyOtherPaymentDefault) Code() int {
	return o._statusCode
}

func (o *CartApplyOtherPaymentDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Payments/Other][%d] Cart_ApplyOtherPayment default %s", o._statusCode, payload)
}

func (o *CartApplyOtherPaymentDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Payments/Other][%d] Cart_ApplyOtherPayment default %s", o._statusCode, payload)
}

func (o *CartApplyOtherPaymentDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CartApplyOtherPaymentDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
