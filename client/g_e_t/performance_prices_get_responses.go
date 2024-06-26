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

// PerformancePricesGetReader is a Reader for the PerformancePricesGet structure.
type PerformancePricesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancePricesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancePricesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancePricesGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancePricesGetOK creates a PerformancePricesGetOK with default headers values
func NewPerformancePricesGetOK() *PerformancePricesGetOK {
	return &PerformancePricesGetOK{}
}

/*
PerformancePricesGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformancePricesGetOK struct {
	Payload *models.PerformancePrice
}

// IsSuccess returns true when this performance prices get o k response has a 2xx status code
func (o *PerformancePricesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance prices get o k response has a 3xx status code
func (o *PerformancePricesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance prices get o k response has a 4xx status code
func (o *PerformancePricesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance prices get o k response has a 5xx status code
func (o *PerformancePricesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance prices get o k response a status code equal to that given
func (o *PerformancePricesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance prices get o k response
func (o *PerformancePricesGetOK) Code() int {
	return 200
}

func (o *PerformancePricesGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePrices/{performancePriceId}][%d] performancePricesGetOK %s", 200, payload)
}

func (o *PerformancePricesGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePrices/{performancePriceId}][%d] performancePricesGetOK %s", 200, payload)
}

func (o *PerformancePricesGetOK) GetPayload() *models.PerformancePrice {
	return o.Payload
}

func (o *PerformancePricesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformancePrice)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformancePricesGetDefault creates a PerformancePricesGetDefault with default headers values
func NewPerformancePricesGetDefault(code int) *PerformancePricesGetDefault {
	return &PerformancePricesGetDefault{
		_statusCode: code,
	}
}

/*
PerformancePricesGetDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancePricesGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performance prices get default response has a 2xx status code
func (o *PerformancePricesGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance prices get default response has a 3xx status code
func (o *PerformancePricesGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance prices get default response has a 4xx status code
func (o *PerformancePricesGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance prices get default response has a 5xx status code
func (o *PerformancePricesGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance prices get default response a status code equal to that given
func (o *PerformancePricesGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance prices get default response
func (o *PerformancePricesGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformancePricesGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePrices/{performancePriceId}][%d] PerformancePrices_Get default %s", o._statusCode, payload)
}

func (o *PerformancePricesGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePrices/{performancePriceId}][%d] PerformancePrices_Get default %s", o._statusCode, payload)
}

func (o *PerformancePricesGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancePricesGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
