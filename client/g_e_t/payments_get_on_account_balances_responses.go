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

// PaymentsGetOnAccountBalancesReader is a Reader for the PaymentsGetOnAccountBalances structure.
type PaymentsGetOnAccountBalancesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentsGetOnAccountBalancesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentsGetOnAccountBalancesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/Payments/OnAccount] Payments_GetOnAccountBalances", response, response.Code())
	}
}

// NewPaymentsGetOnAccountBalancesOK creates a PaymentsGetOnAccountBalancesOK with default headers values
func NewPaymentsGetOnAccountBalancesOK() *PaymentsGetOnAccountBalancesOK {
	return &PaymentsGetOnAccountBalancesOK{}
}

/*
PaymentsGetOnAccountBalancesOK describes a response with status code 200, with default header values.

OK
*/
type PaymentsGetOnAccountBalancesOK struct {
	Payload []*models.ConstituentOnAccountBalance
}

// IsSuccess returns true when this payments get on account balances o k response has a 2xx status code
func (o *PaymentsGetOnAccountBalancesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payments get on account balances o k response has a 3xx status code
func (o *PaymentsGetOnAccountBalancesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payments get on account balances o k response has a 4xx status code
func (o *PaymentsGetOnAccountBalancesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payments get on account balances o k response has a 5xx status code
func (o *PaymentsGetOnAccountBalancesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payments get on account balances o k response a status code equal to that given
func (o *PaymentsGetOnAccountBalancesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payments get on account balances o k response
func (o *PaymentsGetOnAccountBalancesOK) Code() int {
	return 200
}

func (o *PaymentsGetOnAccountBalancesOK) Error() string {
	return fmt.Sprintf("[GET /TXN/Payments/OnAccount][%d] paymentsGetOnAccountBalancesOK  %+v", 200, o.Payload)
}

func (o *PaymentsGetOnAccountBalancesOK) String() string {
	return fmt.Sprintf("[GET /TXN/Payments/OnAccount][%d] paymentsGetOnAccountBalancesOK  %+v", 200, o.Payload)
}

func (o *PaymentsGetOnAccountBalancesOK) GetPayload() []*models.ConstituentOnAccountBalance {
	return o.Payload
}

func (o *PaymentsGetOnAccountBalancesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
