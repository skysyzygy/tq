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

// PerformancePriceLayersGetCountsReader is a Reader for the PerformancePriceLayersGetCounts structure.
type PerformancePriceLayersGetCountsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancePriceLayersGetCountsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancePriceLayersGetCountsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformancePriceLayersGetCountsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformancePriceLayersGetCountsOK creates a PerformancePriceLayersGetCountsOK with default headers values
func NewPerformancePriceLayersGetCountsOK() *PerformancePriceLayersGetCountsOK {
	return &PerformancePriceLayersGetCountsOK{}
}

/*
PerformancePriceLayersGetCountsOK describes a response with status code 200, with default header values.

OK
*/
type PerformancePriceLayersGetCountsOK struct {
	Payload []*models.PerformancePriceLayerCount
}

// IsSuccess returns true when this performance price layers get counts o k response has a 2xx status code
func (o *PerformancePriceLayersGetCountsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance price layers get counts o k response has a 3xx status code
func (o *PerformancePriceLayersGetCountsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance price layers get counts o k response has a 4xx status code
func (o *PerformancePriceLayersGetCountsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance price layers get counts o k response has a 5xx status code
func (o *PerformancePriceLayersGetCountsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance price layers get counts o k response a status code equal to that given
func (o *PerformancePriceLayersGetCountsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performance price layers get counts o k response
func (o *PerformancePriceLayersGetCountsOK) Code() int {
	return 200
}

func (o *PerformancePriceLayersGetCountsOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/Count][%d] performancePriceLayersGetCountsOK %s", 200, payload)
}

func (o *PerformancePriceLayersGetCountsOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/Count][%d] performancePriceLayersGetCountsOK %s", 200, payload)
}

func (o *PerformancePriceLayersGetCountsOK) GetPayload() []*models.PerformancePriceLayerCount {
	return o.Payload
}

func (o *PerformancePriceLayersGetCountsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformancePriceLayersGetCountsDefault creates a PerformancePriceLayersGetCountsDefault with default headers values
func NewPerformancePriceLayersGetCountsDefault(code int) *PerformancePriceLayersGetCountsDefault {
	return &PerformancePriceLayersGetCountsDefault{
		_statusCode: code,
	}
}

/*
PerformancePriceLayersGetCountsDefault describes a response with status code -1, with default header values.

Error
*/
type PerformancePriceLayersGetCountsDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this performance price layers get counts default response has a 2xx status code
func (o *PerformancePriceLayersGetCountsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance price layers get counts default response has a 3xx status code
func (o *PerformancePriceLayersGetCountsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance price layers get counts default response has a 4xx status code
func (o *PerformancePriceLayersGetCountsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance price layers get counts default response has a 5xx status code
func (o *PerformancePriceLayersGetCountsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance price layers get counts default response a status code equal to that given
func (o *PerformancePriceLayersGetCountsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the performance price layers get counts default response
func (o *PerformancePriceLayersGetCountsDefault) Code() int {
	return o._statusCode
}

func (o *PerformancePriceLayersGetCountsDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/Count][%d] PerformancePriceLayers_GetCounts default %s", o._statusCode, payload)
}

func (o *PerformancePriceLayersGetCountsDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PerformancePriceLayers/Count][%d] PerformancePriceLayers_GetCounts default %s", o._statusCode, payload)
}

func (o *PerformancePriceLayersGetCountsDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PerformancePriceLayersGetCountsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
