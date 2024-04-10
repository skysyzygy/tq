// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// PaymentHistoryGetAllReader is a Reader for the PaymentHistoryGetAll structure.
type PaymentHistoryGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentHistoryGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentHistoryGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/PaymentHistory] PaymentHistory_GetAll", response, response.Code())
	}
}

// NewPaymentHistoryGetAllOK creates a PaymentHistoryGetAllOK with default headers values
func NewPaymentHistoryGetAllOK() *PaymentHistoryGetAllOK {
	return &PaymentHistoryGetAllOK{}
}

/*
PaymentHistoryGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PaymentHistoryGetAllOK struct {
	Payload *models.PaymentHistoryResponse
}

// IsSuccess returns true when this payment history get all o k response has a 2xx status code
func (o *PaymentHistoryGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment history get all o k response has a 3xx status code
func (o *PaymentHistoryGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment history get all o k response has a 4xx status code
func (o *PaymentHistoryGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment history get all o k response has a 5xx status code
func (o *PaymentHistoryGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment history get all o k response a status code equal to that given
func (o *PaymentHistoryGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment history get all o k response
func (o *PaymentHistoryGetAllOK) Code() int {
	return 200
}

func (o *PaymentHistoryGetAllOK) Error() string {
	return fmt.Sprintf("[GET /CRM/PaymentHistory][%d] paymentHistoryGetAllOK  %+v", 200, o.Payload)
}

func (o *PaymentHistoryGetAllOK) String() string {
	return fmt.Sprintf("[GET /CRM/PaymentHistory][%d] paymentHistoryGetAllOK  %+v", 200, o.Payload)
}

func (o *PaymentHistoryGetAllOK) GetPayload() *models.PaymentHistoryResponse {
	return o.Payload
}

func (o *PaymentHistoryGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentHistoryResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
