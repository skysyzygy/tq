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

// PaymentTypesGetAllReader is a Reader for the PaymentTypesGetAll structure.
type PaymentTypesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PaymentTypesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPaymentTypesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PaymentTypes] PaymentTypes_GetAll", response, response.Code())
	}
}

// NewPaymentTypesGetAllOK creates a PaymentTypesGetAllOK with default headers values
func NewPaymentTypesGetAllOK() *PaymentTypesGetAllOK {
	return &PaymentTypesGetAllOK{}
}

/*
PaymentTypesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PaymentTypesGetAllOK struct {
	Payload []*models.PaymentType
}

// IsSuccess returns true when this payment types get all o k response has a 2xx status code
func (o *PaymentTypesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this payment types get all o k response has a 3xx status code
func (o *PaymentTypesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this payment types get all o k response has a 4xx status code
func (o *PaymentTypesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this payment types get all o k response has a 5xx status code
func (o *PaymentTypesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this payment types get all o k response a status code equal to that given
func (o *PaymentTypesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the payment types get all o k response
func (o *PaymentTypesGetAllOK) Code() int {
	return 200
}

func (o *PaymentTypesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PaymentTypes][%d] paymentTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *PaymentTypesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PaymentTypes][%d] paymentTypesGetAllOK  %+v", 200, o.Payload)
}

func (o *PaymentTypesGetAllOK) GetPayload() []*models.PaymentType {
	return o.Payload
}

func (o *PaymentTypesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
