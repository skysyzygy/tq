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

// PerformancePriceLayersGetReader is a Reader for the PerformancePriceLayersGet structure.
type PerformancePriceLayersGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancePriceLayersGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancePriceLayersGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancePriceLayersGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancePriceLayersGetOK creates a PerformancePriceLayersGetOK with default headers values
func NewPerformancePriceLayersGetOK() *PerformancePriceLayersGetOK {
	return &PerformancePriceLayersGetOK{}
}

/*
PerformancePriceLayersGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformancePriceLayersGetOK struct {
	Payload *models.PerformancePriceLayer
}

// IsSuccess returns true when this performance price layers get o k response has a 2xx status code
func (o *PerformancePriceLayersGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance price layers get o k response has a 3xx status code
func (o *PerformancePriceLayersGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance price layers get o k response has a 4xx status code
func (o *PerformancePriceLayersGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance price layers get o k response has a 5xx status code
func (o *PerformancePriceLayersGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance price layers get o k response a status code equal to that given
func (o *PerformancePriceLayersGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance price layers get o k response
func (o *PerformancePriceLayersGetOK) Code() int {
	return 200
}

func (o *PerformancePriceLayersGetOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/{performancePriceLayerId}][%d] performancePriceLayersGetOK %s", 200, payload)
}

func (o *PerformancePriceLayersGetOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/{performancePriceLayerId}][%d] performancePriceLayersGetOK %s", 200, payload)
}

func (o *PerformancePriceLayersGetOK) GetPayload() *models.PerformancePriceLayer {
	return o.Payload
}

func (o *PerformancePriceLayersGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformancePriceLayer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformancePriceLayersGetDefault creates a PerformancePriceLayersGetDefault with default headers values
func NewPerformancePriceLayersGetDefault(code int) *PerformancePriceLayersGetDefault {
	return &PerformancePriceLayersGetDefault{
		_statusCode: code,
	}
}

/*
PerformancePriceLayersGetDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancePriceLayersGetDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performance price layers get default response has a 2xx status code
func (o *PerformancePriceLayersGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance price layers get default response has a 3xx status code
func (o *PerformancePriceLayersGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance price layers get default response has a 4xx status code
func (o *PerformancePriceLayersGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance price layers get default response has a 5xx status code
func (o *PerformancePriceLayersGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance price layers get default response a status code equal to that given
func (o *PerformancePriceLayersGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance price layers get default response
func (o *PerformancePriceLayersGetDefault) Code() int {
	return o._statusCode
}

func (o *PerformancePriceLayersGetDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/{performancePriceLayerId}][%d] PerformancePriceLayers_Get default %s", o._statusCode, payload)
}

func (o *PerformancePriceLayersGetDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/{performancePriceLayerId}][%d] PerformancePriceLayers_Get default %s", o._statusCode, payload)
}

func (o *PerformancePriceLayersGetDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancePriceLayersGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
