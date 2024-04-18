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

// DeliveryMethodsGetAllReader is a Reader for the DeliveryMethodsGetAll structure.
type DeliveryMethodsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeliveryMethodsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeliveryMethodsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/DeliveryMethods] DeliveryMethods_GetAll", response, response.Code())
	}
}

// NewDeliveryMethodsGetAllOK creates a DeliveryMethodsGetAllOK with default headers values
func NewDeliveryMethodsGetAllOK() *DeliveryMethodsGetAllOK {
	return &DeliveryMethodsGetAllOK{}
}

/*
DeliveryMethodsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type DeliveryMethodsGetAllOK struct {
	Payload []*models.DeliveryMethod
}

// IsSuccess returns true when this delivery methods get all o k response has a 2xx status code
func (o *DeliveryMethodsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delivery methods get all o k response has a 3xx status code
func (o *DeliveryMethodsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delivery methods get all o k response has a 4xx status code
func (o *DeliveryMethodsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this delivery methods get all o k response has a 5xx status code
func (o *DeliveryMethodsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this delivery methods get all o k response a status code equal to that given
func (o *DeliveryMethodsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the delivery methods get all o k response
func (o *DeliveryMethodsGetAllOK) Code() int {
	return 200
}

func (o *DeliveryMethodsGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/DeliveryMethods][%d] deliveryMethodsGetAllOK  %+v", 200, o.Payload)
}

func (o *DeliveryMethodsGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/DeliveryMethods][%d] deliveryMethodsGetAllOK  %+v", 200, o.Payload)
}

func (o *DeliveryMethodsGetAllOK) GetPayload() []*models.DeliveryMethod {
	return o.Payload
}

func (o *DeliveryMethodsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}