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

// PaymentMethodsTranslateMnemonicReader is a Reader for the PaymentMethodsTranslateMnemonic structure.
type PaymentMethodsTranslateMnemonicReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentMethodsTranslateMnemonicReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentMethodsTranslateMnemonicOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/PaymentMethods/TranslateMnemonic] PaymentMethods_TranslateMnemonic", response, response.Code())
	}
}

// NewPaymentMethodsTranslateMnemonicOK creates a PaymentMethodsTranslateMnemonicOK with default headers values
func NewPaymentMethodsTranslateMnemonicOK() *PaymentMethodsTranslateMnemonicOK {
	return &PaymentMethodsTranslateMnemonicOK{}
}

/*
PaymentMethodsTranslateMnemonicOK describes a response with status code 200, with default header values.

OK
*/
type PaymentMethodsTranslateMnemonicOK struct {
	Payload *models.PaymentMethod
}

// IsSuccess returns true when this payment methods translate mnemonic o k response has a 2xx status code
func (o *PaymentMethodsTranslateMnemonicOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment methods translate mnemonic o k response has a 3xx status code
func (o *PaymentMethodsTranslateMnemonicOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment methods translate mnemonic o k response has a 4xx status code
func (o *PaymentMethodsTranslateMnemonicOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment methods translate mnemonic o k response has a 5xx status code
func (o *PaymentMethodsTranslateMnemonicOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment methods translate mnemonic o k response a status code equal to that given
func (o *PaymentMethodsTranslateMnemonicOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment methods translate mnemonic o k response
func (o *PaymentMethodsTranslateMnemonicOK) Code() int {
	return 200
}

func (o *PaymentMethodsTranslateMnemonicOK) Error() string {
	return fmt.Sprintf("[GET /TXN/PaymentMethods/TranslateMnemonic][%d] paymentMethodsTranslateMnemonicOK  %+v", 200, o.Payload)
}

func (o *PaymentMethodsTranslateMnemonicOK) String() string {
	return fmt.Sprintf("[GET /TXN/PaymentMethods/TranslateMnemonic][%d] paymentMethodsTranslateMnemonicOK  %+v", 200, o.Payload)
}

func (o *PaymentMethodsTranslateMnemonicOK) GetPayload() *models.PaymentMethod {
	return o.Payload
}

func (o *PaymentMethodsTranslateMnemonicOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaymentMethod)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
