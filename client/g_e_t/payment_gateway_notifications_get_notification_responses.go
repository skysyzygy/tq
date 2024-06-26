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

// PaymentGatewayNotificationsGetNotificationReader is a Reader for the PaymentGatewayNotificationsGetNotification structure.
type PaymentGatewayNotificationsGetNotificationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGatewayNotificationsGetNotificationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGatewayNotificationsGetNotificationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPaymentGatewayNotificationsGetNotificationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPaymentGatewayNotificationsGetNotificationOK creates a PaymentGatewayNotificationsGetNotificationOK with default headers values
func NewPaymentGatewayNotificationsGetNotificationOK() *PaymentGatewayNotificationsGetNotificationOK {
	return &PaymentGatewayNotificationsGetNotificationOK{}
}

/*
PaymentGatewayNotificationsGetNotificationOK describes a response with status code 200, with default header values.

OK
*/
type PaymentGatewayNotificationsGetNotificationOK struct {
	Payload []*models.TessituraPaymentsNotification
}

// IsSuccess returns true when this payment gateway notifications get notification o k response has a 2xx status code
func (o *PaymentGatewayNotificationsGetNotificationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment gateway notifications get notification o k response has a 3xx status code
func (o *PaymentGatewayNotificationsGetNotificationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment gateway notifications get notification o k response has a 4xx status code
func (o *PaymentGatewayNotificationsGetNotificationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment gateway notifications get notification o k response has a 5xx status code
func (o *PaymentGatewayNotificationsGetNotificationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment gateway notifications get notification o k response a status code equal to that given
func (o *PaymentGatewayNotificationsGetNotificationOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment gateway notifications get notification o k response
func (o *PaymentGatewayNotificationsGetNotificationOK) Code() int {
	return 200
}

func (o *PaymentGatewayNotificationsGetNotificationOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Notifications][%d] paymentGatewayNotificationsGetNotificationOK %s", 200, payload)
}

func (o *PaymentGatewayNotificationsGetNotificationOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Notifications][%d] paymentGatewayNotificationsGetNotificationOK %s", 200, payload)
}

func (o *PaymentGatewayNotificationsGetNotificationOK) GetPayload() []*models.TessituraPaymentsNotification {
	return o.Payload
}

func (o *PaymentGatewayNotificationsGetNotificationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPaymentGatewayNotificationsGetNotificationDefault creates a PaymentGatewayNotificationsGetNotificationDefault with default headers values
func NewPaymentGatewayNotificationsGetNotificationDefault(code int) *PaymentGatewayNotificationsGetNotificationDefault {
	return &PaymentGatewayNotificationsGetNotificationDefault{
		_statusCode: code,
	}
}

/*
PaymentGatewayNotificationsGetNotificationDefault describes a response with status code -1, with default header values.

Error
*/
type PaymentGatewayNotificationsGetNotificationDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this payment gateway notifications get notification default response has a 2xx status code
func (o *PaymentGatewayNotificationsGetNotificationDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this payment gateway notifications get notification default response has a 3xx status code
func (o *PaymentGatewayNotificationsGetNotificationDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this payment gateway notifications get notification default response has a 4xx status code
func (o *PaymentGatewayNotificationsGetNotificationDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this payment gateway notifications get notification default response has a 5xx status code
func (o *PaymentGatewayNotificationsGetNotificationDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this payment gateway notifications get notification default response a status code equal to that given
func (o *PaymentGatewayNotificationsGetNotificationDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the payment gateway notifications get notification default response
func (o *PaymentGatewayNotificationsGetNotificationDefault) Code() int {
	return o._statusCode
}

func (o *PaymentGatewayNotificationsGetNotificationDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Notifications][%d] PaymentGatewayNotifications_GetNotification default %s", o._statusCode, payload)
}

func (o *PaymentGatewayNotificationsGetNotificationDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /PaymentGateway/Notifications][%d] PaymentGatewayNotifications_GetNotification default %s", o._statusCode, payload)
}

func (o *PaymentGatewayNotificationsGetNotificationDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PaymentGatewayNotificationsGetNotificationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
