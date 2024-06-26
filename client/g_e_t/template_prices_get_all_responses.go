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

// TemplatePricesGetAllReader is a Reader for the TemplatePricesGetAll structure.
type TemplatePricesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TemplatePricesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTemplatePricesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTemplatePricesGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTemplatePricesGetAllOK creates a TemplatePricesGetAllOK with default headers values
func NewTemplatePricesGetAllOK() *TemplatePricesGetAllOK {
	return &TemplatePricesGetAllOK{}
}

/*
TemplatePricesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type TemplatePricesGetAllOK struct {
	Payload []*models.TemplatePrice
}

// IsSuccess returns true when this template prices get all o k response has a 2xx status code
func (o *TemplatePricesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this template prices get all o k response has a 3xx status code
func (o *TemplatePricesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this template prices get all o k response has a 4xx status code
func (o *TemplatePricesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this template prices get all o k response has a 5xx status code
func (o *TemplatePricesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this template prices get all o k response a status code equal to that given
func (o *TemplatePricesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the template prices get all o k response
func (o *TemplatePricesGetAllOK) Code() int {
	return 200
}

func (o *TemplatePricesGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/TemplatePrices][%d] templatePricesGetAllOK %s", 200, payload)
}

func (o *TemplatePricesGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/TemplatePrices][%d] templatePricesGetAllOK %s", 200, payload)
}

func (o *TemplatePricesGetAllOK) GetPayload() []*models.TemplatePrice {
	return o.Payload
}

func (o *TemplatePricesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTemplatePricesGetAllDefault creates a TemplatePricesGetAllDefault with default headers values
func NewTemplatePricesGetAllDefault(code int) *TemplatePricesGetAllDefault {
	return &TemplatePricesGetAllDefault{
		_statusCode: code,
	}
}

/*
TemplatePricesGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type TemplatePricesGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this template prices get all default response has a 2xx status code
func (o *TemplatePricesGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this template prices get all default response has a 3xx status code
func (o *TemplatePricesGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this template prices get all default response has a 4xx status code
func (o *TemplatePricesGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this template prices get all default response has a 5xx status code
func (o *TemplatePricesGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this template prices get all default response a status code equal to that given
func (o *TemplatePricesGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the template prices get all default response
func (o *TemplatePricesGetAllDefault) Code() int {
	return o._statusCode
}

func (o *TemplatePricesGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/TemplatePrices][%d] TemplatePrices_GetAll default %s", o._statusCode, payload)
}

func (o *TemplatePricesGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/TemplatePrices][%d] TemplatePrices_GetAll default %s", o._statusCode, payload)
}

func (o *TemplatePricesGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *TemplatePricesGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
