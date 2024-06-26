// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// PaymentMethodGroupsUpdateReader is a Reader for the PaymentMethodGroupsUpdate structure.
type PaymentMethodGroupsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentMethodGroupsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentMethodGroupsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentMethodGroupsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentMethodGroupsUpdateOK creates a PaymentMethodGroupsUpdateOK with default headers values
func NewPaymentMethodGroupsUpdateOK() *PaymentMethodGroupsUpdateOK {
	return &PaymentMethodGroupsUpdateOK{}
}

/*
PaymentMethodGroupsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PaymentMethodGroupsUpdateOK struct {
	Payload *models.PaymentMethodGroup
}

// IsSuccess returns true when this payment method groups update o k response has a 2xx status code
func (o *PaymentMethodGroupsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment method groups update o k response has a 3xx status code
func (o *PaymentMethodGroupsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment method groups update o k response has a 4xx status code
func (o *PaymentMethodGroupsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment method groups update o k response has a 5xx status code
func (o *PaymentMethodGroupsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment method groups update o k response a status code equal to that given
func (o *PaymentMethodGroupsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment method groups update o k response
func (o *PaymentMethodGroupsUpdateOK) Code() int {
	return 200
}

func (o *PaymentMethodGroupsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PaymentMethodGroups/{id}][%d] paymentMethodGroupsUpdateOK %s", 200, payload)
}

func (o *PaymentMethodGroupsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PaymentMethodGroups/{id}][%d] paymentMethodGroupsUpdateOK %s", 200, payload)
}

func (o *PaymentMethodGroupsUpdateOK) GetPayload() *models.PaymentMethodGroup {
	return o.Payload
}

func (o *PaymentMethodGroupsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethodGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentMethodGroupsUpdateDefault creates a PaymentMethodGroupsUpdateDefault with default headers values
func NewPaymentMethodGroupsUpdateDefault(code int) *PaymentMethodGroupsUpdateDefault {
	return &PaymentMethodGroupsUpdateDefault{
		_statusCode: code,
	}
}

/*
PaymentMethodGroupsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PaymentMethodGroupsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this payment method groups update default response has a 2xx status code
func (o *PaymentMethodGroupsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment method groups update default response has a 3xx status code
func (o *PaymentMethodGroupsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment method groups update default response has a 4xx status code
func (o *PaymentMethodGroupsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment method groups update default response has a 5xx status code
func (o *PaymentMethodGroupsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment method groups update default response a status code equal to that given
func (o *PaymentMethodGroupsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the payment method groups update default response
func (o *PaymentMethodGroupsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PaymentMethodGroupsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PaymentMethodGroups/{id}][%d] PaymentMethodGroups_Update default %s", o._statusCode, payload)
}

func (o *PaymentMethodGroupsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PaymentMethodGroups/{id}][%d] PaymentMethodGroups_Update default %s", o._statusCode, payload)
}

func (o *PaymentMethodGroupsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PaymentMethodGroupsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
