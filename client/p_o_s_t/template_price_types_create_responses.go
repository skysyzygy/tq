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

// TemplatePriceTypesCreateReader is a Reader for the TemplatePriceTypesCreate structure.
type TemplatePriceTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatePriceTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplatePriceTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /TXN/TemplatePriceTypes] TemplatePriceTypes_Create", response, response.Code())
	}
}

// NewTemplatePriceTypesCreateOK creates a TemplatePriceTypesCreateOK with default headers values
func NewTemplatePriceTypesCreateOK() *TemplatePriceTypesCreateOK {
	return &TemplatePriceTypesCreateOK{}
}

/*
TemplatePriceTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type TemplatePriceTypesCreateOK struct {
	Payload *models.TemplatePriceType
}

// IsSuccess returns true when this template price types create o k response has a 2xx status code
func (o *TemplatePriceTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template price types create o k response has a 3xx status code
func (o *TemplatePriceTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template price types create o k response has a 4xx status code
func (o *TemplatePriceTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template price types create o k response has a 5xx status code
func (o *TemplatePriceTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template price types create o k response a status code equal to that given
func (o *TemplatePriceTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template price types create o k response
func (o *TemplatePriceTypesCreateOK) Code() int {
	return 200
}

func (o *TemplatePriceTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /TXN/TemplatePriceTypes][%d] templatePriceTypesCreateOK  %+v", 200, o.Payload)
}

func (o *TemplatePriceTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /TXN/TemplatePriceTypes][%d] templatePriceTypesCreateOK  %+v", 200, o.Payload)
}

func (o *TemplatePriceTypesCreateOK) GetPayload() *models.TemplatePriceType {
	return o.Payload
}

func (o *TemplatePriceTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TemplatePriceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
