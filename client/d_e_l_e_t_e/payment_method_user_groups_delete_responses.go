// Code generated by go-swagger; DO NOT EDIT.

package d_e_l_e_t_e

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

// PaymentMethodUserGroupsDeleteReader is a Reader for the PaymentMethodUserGroupsDelete structure.
type PaymentMethodUserGroupsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentMethodUserGroupsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPaymentMethodUserGroupsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentMethodUserGroupsDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentMethodUserGroupsDeleteNoContent creates a PaymentMethodUserGroupsDeleteNoContent with default headers values
func NewPaymentMethodUserGroupsDeleteNoContent() *PaymentMethodUserGroupsDeleteNoContent {
	return &PaymentMethodUserGroupsDeleteNoContent{}
}

/*
PaymentMethodUserGroupsDeleteNoContent describes a response with status code 204, with default header values.

No Content
*/
type PaymentMethodUserGroupsDeleteNoContent struct {
}

// IsSuccess returns true when this payment method user groups delete no content response has a 2xx status code
func (o *PaymentMethodUserGroupsDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment method user groups delete no content response has a 3xx status code
func (o *PaymentMethodUserGroupsDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment method user groups delete no content response has a 4xx status code
func (o *PaymentMethodUserGroupsDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment method user groups delete no content response has a 5xx status code
func (o *PaymentMethodUserGroupsDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this payment method user groups delete no content response a status code equal to that given
func (o *PaymentMethodUserGroupsDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the payment method user groups delete no content response
func (o *PaymentMethodUserGroupsDeleteNoContent) Code() int {
	return 204
}

func (o *PaymentMethodUserGroupsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /TXN/PaymentMethodUserGroups/{paymentMethodUserGroupId}][%d] paymentMethodUserGroupsDeleteNoContent", 204)
}

func (o *PaymentMethodUserGroupsDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /TXN/PaymentMethodUserGroups/{paymentMethodUserGroupId}][%d] paymentMethodUserGroupsDeleteNoContent", 204)
}

func (o *PaymentMethodUserGroupsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPaymentMethodUserGroupsDeleteDefault creates a PaymentMethodUserGroupsDeleteDefault with default headers values
func NewPaymentMethodUserGroupsDeleteDefault(code int) *PaymentMethodUserGroupsDeleteDefault {
	return &PaymentMethodUserGroupsDeleteDefault{
		_statusCode: code,
	}
}

/*
PaymentMethodUserGroupsDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type PaymentMethodUserGroupsDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this payment method user groups delete default response has a 2xx status code
func (o *PaymentMethodUserGroupsDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment method user groups delete default response has a 3xx status code
func (o *PaymentMethodUserGroupsDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment method user groups delete default response has a 4xx status code
func (o *PaymentMethodUserGroupsDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment method user groups delete default response has a 5xx status code
func (o *PaymentMethodUserGroupsDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment method user groups delete default response a status code equal to that given
func (o *PaymentMethodUserGroupsDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the payment method user groups delete default response
func (o *PaymentMethodUserGroupsDeleteDefault) Code() int {
	return o._statusCode
}

func (o *PaymentMethodUserGroupsDeleteDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /TXN/PaymentMethodUserGroups/{paymentMethodUserGroupId}][%d] PaymentMethodUserGroups_Delete default %s", o._statusCode, payload)
}

func (o *PaymentMethodUserGroupsDeleteDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /TXN/PaymentMethodUserGroups/{paymentMethodUserGroupId}][%d] PaymentMethodUserGroups_Delete default %s", o._statusCode, payload)
}

func (o *PaymentMethodUserGroupsDeleteDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PaymentMethodUserGroupsDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
