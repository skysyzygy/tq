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

// PaymentGatewayNotificationsCreateNotificationEventReader is a Reader for the PaymentGatewayNotificationsCreateNotificationEvent structure.
type PaymentGatewayNotificationsCreateNotificationEventReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentGatewayNotificationsCreateNotificationEventReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentGatewayNotificationsCreateNotificationEventOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /PaymentGateway/Notifications/Events] PaymentGatewayNotifications_CreateNotificationEvent", response, response.Code())
	}
}

// NewPaymentGatewayNotificationsCreateNotificationEventOK creates a PaymentGatewayNotificationsCreateNotificationEventOK with default headers values
func NewPaymentGatewayNotificationsCreateNotificationEventOK() *PaymentGatewayNotificationsCreateNotificationEventOK {
	return &PaymentGatewayNotificationsCreateNotificationEventOK{}
}

/*
PaymentGatewayNotificationsCreateNotificationEventOK describes a response with status code 200, with default header values.

OK
*/
type PaymentGatewayNotificationsCreateNotificationEventOK struct {
	Payload *models.NotificationEvent
}

// IsSuccess returns true when this payment gateway notifications create notification event o k response has a 2xx status code
func (o *PaymentGatewayNotificationsCreateNotificationEventOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment gateway notifications create notification event o k response has a 3xx status code
func (o *PaymentGatewayNotificationsCreateNotificationEventOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment gateway notifications create notification event o k response has a 4xx status code
func (o *PaymentGatewayNotificationsCreateNotificationEventOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment gateway notifications create notification event o k response has a 5xx status code
func (o *PaymentGatewayNotificationsCreateNotificationEventOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment gateway notifications create notification event o k response a status code equal to that given
func (o *PaymentGatewayNotificationsCreateNotificationEventOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment gateway notifications create notification event o k response
func (o *PaymentGatewayNotificationsCreateNotificationEventOK) Code() int {
	return 200
}

func (o *PaymentGatewayNotificationsCreateNotificationEventOK) Error() string {
	return fmt.Sprintf("[POST /PaymentGateway/Notifications/Events][%d] paymentGatewayNotificationsCreateNotificationEventOK  %+v", 200, o.Payload)
}

func (o *PaymentGatewayNotificationsCreateNotificationEventOK) String() string {
	return fmt.Sprintf("[POST /PaymentGateway/Notifications/Events][%d] paymentGatewayNotificationsCreateNotificationEventOK  %+v", 200, o.Payload)
}

func (o *PaymentGatewayNotificationsCreateNotificationEventOK) GetPayload() *models.NotificationEvent {
	return o.Payload
}

func (o *PaymentGatewayNotificationsCreateNotificationEventOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NotificationEvent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
