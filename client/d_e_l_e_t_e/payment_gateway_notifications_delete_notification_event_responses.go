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

// PaymentGatewayNotificationsDeleteNotificationEventReader is a Reader for the PaymentGatewayNotificationsDeleteNotificationEvent structure.
type PaymentGatewayNotificationsDeleteNotificationEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGatewayNotificationsDeleteNotificationEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPaymentGatewayNotificationsDeleteNotificationEventNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentGatewayNotificationsDeleteNotificationEventDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentGatewayNotificationsDeleteNotificationEventNoContent creates a PaymentGatewayNotificationsDeleteNotificationEventNoContent with default headers values
func NewPaymentGatewayNotificationsDeleteNotificationEventNoContent() *PaymentGatewayNotificationsDeleteNotificationEventNoContent {
	return &PaymentGatewayNotificationsDeleteNotificationEventNoContent{}
}

/*
PaymentGatewayNotificationsDeleteNotificationEventNoContent describes a response with status code 204, with default header values.

No Content
*/
type PaymentGatewayNotificationsDeleteNotificationEventNoContent struct {
}

// IsSuccess returns true when this payment gateway notifications delete notification event no content response has a 2xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment gateway notifications delete notification event no content response has a 3xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment gateway notifications delete notification event no content response has a 4xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment gateway notifications delete notification event no content response has a 5xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this payment gateway notifications delete notification event no content response a status code equal to that given
func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the payment gateway notifications delete notification event no content response
func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) Code() int {
	return 204
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) Error() string {
	return fmt.Sprintf("[DELETE /PaymentGateway/Notifications/Events/{notificationEventId}][%d] paymentGatewayNotificationsDeleteNotificationEventNoContent", 204)
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) String() string {
	return fmt.Sprintf("[DELETE /PaymentGateway/Notifications/Events/{notificationEventId}][%d] paymentGatewayNotificationsDeleteNotificationEventNoContent", 204)
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPaymentGatewayNotificationsDeleteNotificationEventDefault creates a PaymentGatewayNotificationsDeleteNotificationEventDefault with default headers values
func NewPaymentGatewayNotificationsDeleteNotificationEventDefault(code int) *PaymentGatewayNotificationsDeleteNotificationEventDefault {
	return &PaymentGatewayNotificationsDeleteNotificationEventDefault{
		_statusCode: code,
	}
}

/*
PaymentGatewayNotificationsDeleteNotificationEventDefault describes a response with status code -1, with default header values.

Error
*/
type PaymentGatewayNotificationsDeleteNotificationEventDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this payment gateway notifications delete notification event default response has a 2xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment gateway notifications delete notification event default response has a 3xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment gateway notifications delete notification event default response has a 4xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment gateway notifications delete notification event default response has a 5xx status code
func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment gateway notifications delete notification event default response a status code equal to that given
func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the payment gateway notifications delete notification event default response
func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) Code() int {
	return o._statusCode
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /PaymentGateway/Notifications/Events/{notificationEventId}][%d] PaymentGatewayNotifications_DeleteNotificationEvent default %s", o._statusCode, payload)
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /PaymentGateway/Notifications/Events/{notificationEventId}][%d] PaymentGatewayNotifications_DeleteNotificationEvent default %s", o._statusCode, payload)
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PaymentGatewayNotificationsDeleteNotificationEventDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
