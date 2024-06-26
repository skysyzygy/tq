// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// PerformancePriceLayersCreateReader is a Reader for the PerformancePriceLayersCreate structure.
type PerformancePriceLayersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancePriceLayersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancePriceLayersCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancePriceLayersCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancePriceLayersCreateOK creates a PerformancePriceLayersCreateOK with default headers values
func NewPerformancePriceLayersCreateOK() *PerformancePriceLayersCreateOK {
	return &PerformancePriceLayersCreateOK{}
}

/*
PerformancePriceLayersCreateOK describes a response with status code 200, with default header values.

OK
*/
type PerformancePriceLayersCreateOK struct {
	Payload *models.PerformancePriceLayer
}

// IsSuccess returns true when this performance price layers create o k response has a 2xx status code
func (o *PerformancePriceLayersCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance price layers create o k response has a 3xx status code
func (o *PerformancePriceLayersCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance price layers create o k response has a 4xx status code
func (o *PerformancePriceLayersCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance price layers create o k response has a 5xx status code
func (o *PerformancePriceLayersCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance price layers create o k response a status code equal to that given
func (o *PerformancePriceLayersCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance price layers create o k response
func (o *PerformancePriceLayersCreateOK) Code() int {
	return 200
}

func (o *PerformancePriceLayersCreateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/PerformancePriceLayers][%d] performancePriceLayersCreateOK %s", 200, payload)
}

func (o *PerformancePriceLayersCreateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/PerformancePriceLayers][%d] performancePriceLayersCreateOK %s", 200, payload)
}

func (o *PerformancePriceLayersCreateOK) GetPayload() *models.PerformancePriceLayer {
	return o.Payload
}

func (o *PerformancePriceLayersCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformancePriceLayer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformancePriceLayersCreateDefault creates a PerformancePriceLayersCreateDefault with default headers values
func NewPerformancePriceLayersCreateDefault(code int) *PerformancePriceLayersCreateDefault {
	return &PerformancePriceLayersCreateDefault{
		_statusCode: code,
	}
}

/*
PerformancePriceLayersCreateDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancePriceLayersCreateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performance price layers create default response has a 2xx status code
func (o *PerformancePriceLayersCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance price layers create default response has a 3xx status code
func (o *PerformancePriceLayersCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance price layers create default response has a 4xx status code
func (o *PerformancePriceLayersCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance price layers create default response has a 5xx status code
func (o *PerformancePriceLayersCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance price layers create default response a status code equal to that given
func (o *PerformancePriceLayersCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance price layers create default response
func (o *PerformancePriceLayersCreateDefault) Code() int {
	return o._statusCode
}

func (o *PerformancePriceLayersCreateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/PerformancePriceLayers][%d] PerformancePriceLayers_Create default %s", o._statusCode, payload)
}

func (o *PerformancePriceLayersCreateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /TXN/PerformancePriceLayers][%d] PerformancePriceLayers_Create default %s", o._statusCode, payload)
}

func (o *PerformancePriceLayersCreateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancePriceLayersCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
