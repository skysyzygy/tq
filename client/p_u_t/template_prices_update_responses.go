// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// TemplatePricesUpdateReader is a Reader for the TemplatePricesUpdate structure.
type TemplatePricesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatePricesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplatePricesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /TXN/TemplatePrices/{templatePriceId}] TemplatePrices_Update", response, response.Code())
	}
}

// NewTemplatePricesUpdateOK creates a TemplatePricesUpdateOK with default headers values
func NewTemplatePricesUpdateOK() *TemplatePricesUpdateOK {
	return &TemplatePricesUpdateOK{}
}

/*
TemplatePricesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type TemplatePricesUpdateOK struct {
	Payload *models.TemplatePrice
}

// IsSuccess returns true when this template prices update o k response has a 2xx status code
func (o *TemplatePricesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template prices update o k response has a 3xx status code
func (o *TemplatePricesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template prices update o k response has a 4xx status code
func (o *TemplatePricesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template prices update o k response has a 5xx status code
func (o *TemplatePricesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template prices update o k response a status code equal to that given
func (o *TemplatePricesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template prices update o k response
func (o *TemplatePricesUpdateOK) Code() int {
	return 200
}

func (o *TemplatePricesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /TXN/TemplatePrices/{templatePriceId}][%d] templatePricesUpdateOK  %+v", 200, o.Payload)
}

func (o *TemplatePricesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /TXN/TemplatePrices/{templatePriceId}][%d] templatePricesUpdateOK  %+v", 200, o.Payload)
}

func (o *TemplatePricesUpdateOK) GetPayload() *models.TemplatePrice {
	return o.Payload
}

func (o *TemplatePricesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplatePrice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}