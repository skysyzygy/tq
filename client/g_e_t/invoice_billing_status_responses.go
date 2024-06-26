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

// InvoiceBillingStatusReader is a Reader for the InvoiceBillingStatus structure.
type InvoiceBillingStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InvoiceBillingStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInvoiceBillingStatusOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInvoiceBillingStatusDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInvoiceBillingStatusOK creates a InvoiceBillingStatusOK with default headers values
func NewInvoiceBillingStatusOK() *InvoiceBillingStatusOK {
	return &InvoiceBillingStatusOK{}
}

/*
InvoiceBillingStatusOK describes a response with status code 200, with default header values.

OK
*/
type InvoiceBillingStatusOK struct {
	Payload *models.AutomatedBillingStatus
}

// IsSuccess returns true when this invoice billing status o k response has a 2xx status code
func (o *InvoiceBillingStatusOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this invoice billing status o k response has a 3xx status code
func (o *InvoiceBillingStatusOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this invoice billing status o k response has a 4xx status code
func (o *InvoiceBillingStatusOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this invoice billing status o k response has a 5xx status code
func (o *InvoiceBillingStatusOK) IsServerError() bool {
	return false
}

// IsCode returns true when this invoice billing status o k response a status code equal to that given
func (o *InvoiceBillingStatusOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the invoice billing status o k response
func (o *InvoiceBillingStatusOK) Code() int {
	return 200
}

func (o *InvoiceBillingStatusOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/InvoiceBilling/{id}/Status][%d] invoiceBillingStatusOK %s", 200, payload)
}

func (o *InvoiceBillingStatusOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/InvoiceBilling/{id}/Status][%d] invoiceBillingStatusOK %s", 200, payload)
}

func (o *InvoiceBillingStatusOK) GetPayload() *models.AutomatedBillingStatus {
	return o.Payload
}

func (o *InvoiceBillingStatusOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AutomatedBillingStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInvoiceBillingStatusDefault creates a InvoiceBillingStatusDefault with default headers values
func NewInvoiceBillingStatusDefault(code int) *InvoiceBillingStatusDefault {
	return &InvoiceBillingStatusDefault{
		_statusCode: code,
	}
}

/*
InvoiceBillingStatusDefault describes a response with status code -1, with default header values.

Error
*/
type InvoiceBillingStatusDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this invoice billing status default response has a 2xx status code
func (o *InvoiceBillingStatusDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this invoice billing status default response has a 3xx status code
func (o *InvoiceBillingStatusDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this invoice billing status default response has a 4xx status code
func (o *InvoiceBillingStatusDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this invoice billing status default response has a 5xx status code
func (o *InvoiceBillingStatusDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this invoice billing status default response a status code equal to that given
func (o *InvoiceBillingStatusDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the invoice billing status default response
func (o *InvoiceBillingStatusDefault) Code() int {
	return o._statusCode
}

func (o *InvoiceBillingStatusDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/InvoiceBilling/{id}/Status][%d] InvoiceBilling_Status default %s", o._statusCode, payload)
}

func (o *InvoiceBillingStatusDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/InvoiceBilling/{id}/Status][%d] InvoiceBilling_Status default %s", o._statusCode, payload)
}

func (o *InvoiceBillingStatusDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *InvoiceBillingStatusDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
