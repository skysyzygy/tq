// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

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

// PaymentGatewayConfigurationGetConfigurationReader is a Reader for the PaymentGatewayConfigurationGetConfiguration structure.
type PaymentGatewayConfigurationGetConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGatewayConfigurationGetConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGatewayConfigurationGetConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentGatewayConfigurationGetConfigurationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentGatewayConfigurationGetConfigurationOK creates a PaymentGatewayConfigurationGetConfigurationOK with default headers values
func NewPaymentGatewayConfigurationGetConfigurationOK() *PaymentGatewayConfigurationGetConfigurationOK {
	return &PaymentGatewayConfigurationGetConfigurationOK{}
}

/*
PaymentGatewayConfigurationGetConfigurationOK describes a response with status code 200, with default header values.

OK
*/
type PaymentGatewayConfigurationGetConfigurationOK struct {
	Payload *models.PaymentGatewayConfiguration
}

// IsSuccess returns true when this payment gateway configuration get configuration o k response has a 2xx status code
func (o *PaymentGatewayConfigurationGetConfigurationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment gateway configuration get configuration o k response has a 3xx status code
func (o *PaymentGatewayConfigurationGetConfigurationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment gateway configuration get configuration o k response has a 4xx status code
func (o *PaymentGatewayConfigurationGetConfigurationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment gateway configuration get configuration o k response has a 5xx status code
func (o *PaymentGatewayConfigurationGetConfigurationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment gateway configuration get configuration o k response a status code equal to that given
func (o *PaymentGatewayConfigurationGetConfigurationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment gateway configuration get configuration o k response
func (o *PaymentGatewayConfigurationGetConfigurationOK) Code() int {
	return 200
}

func (o *PaymentGatewayConfigurationGetConfigurationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Configuration][%d] paymentGatewayConfigurationGetConfigurationOK %s", 200, payload)
}

func (o *PaymentGatewayConfigurationGetConfigurationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Configuration][%d] paymentGatewayConfigurationGetConfigurationOK %s", 200, payload)
}

func (o *PaymentGatewayConfigurationGetConfigurationOK) GetPayload() *models.PaymentGatewayConfiguration {
	return o.Payload
}

func (o *PaymentGatewayConfigurationGetConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentGatewayConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGatewayConfigurationGetConfigurationDefault creates a PaymentGatewayConfigurationGetConfigurationDefault with default headers values
func NewPaymentGatewayConfigurationGetConfigurationDefault(code int) *PaymentGatewayConfigurationGetConfigurationDefault {
	return &PaymentGatewayConfigurationGetConfigurationDefault{
		_statusCode: code,
	}
}

/*
PaymentGatewayConfigurationGetConfigurationDefault describes a response with status code -1, with default header values.

Error
*/
type PaymentGatewayConfigurationGetConfigurationDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this payment gateway configuration get configuration default response has a 2xx status code
func (o *PaymentGatewayConfigurationGetConfigurationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment gateway configuration get configuration default response has a 3xx status code
func (o *PaymentGatewayConfigurationGetConfigurationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment gateway configuration get configuration default response has a 4xx status code
func (o *PaymentGatewayConfigurationGetConfigurationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment gateway configuration get configuration default response has a 5xx status code
func (o *PaymentGatewayConfigurationGetConfigurationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment gateway configuration get configuration default response a status code equal to that given
func (o *PaymentGatewayConfigurationGetConfigurationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the payment gateway configuration get configuration default response
func (o *PaymentGatewayConfigurationGetConfigurationDefault) Code() int {
	return o._statusCode
}

func (o *PaymentGatewayConfigurationGetConfigurationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Configuration][%d] PaymentGatewayConfiguration_GetConfiguration default %s", o._statusCode, payload)
}

func (o *PaymentGatewayConfigurationGetConfigurationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Configuration][%d] PaymentGatewayConfiguration_GetConfiguration default %s", o._statusCode, payload)
}

func (o *PaymentGatewayConfigurationGetConfigurationDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PaymentGatewayConfigurationGetConfigurationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
