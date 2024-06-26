// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// PriceLayerTypesUpdateReader is a Reader for the PriceLayerTypesUpdate structure.
type PriceLayerTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceLayerTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceLayerTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPriceLayerTypesUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPriceLayerTypesUpdateOK creates a PriceLayerTypesUpdateOK with default headers values
func NewPriceLayerTypesUpdateOK() *PriceLayerTypesUpdateOK {
	return &PriceLayerTypesUpdateOK{}
}

/*
PriceLayerTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PriceLayerTypesUpdateOK struct {
	Payload *models.PriceLayerType
}

// IsSuccess returns true when this price layer types update o k response has a 2xx status code
func (o *PriceLayerTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price layer types update o k response has a 3xx status code
func (o *PriceLayerTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price layer types update o k response has a 4xx status code
func (o *PriceLayerTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price layer types update o k response has a 5xx status code
func (o *PriceLayerTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price layer types update o k response a status code equal to that given
func (o *PriceLayerTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price layer types update o k response
func (o *PriceLayerTypesUpdateOK) Code() int {
	return 200
}

func (o *PriceLayerTypesUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PriceLayerTypes/{id}][%d] priceLayerTypesUpdateOK %s", 200, payload)
}

func (o *PriceLayerTypesUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PriceLayerTypes/{id}][%d] priceLayerTypesUpdateOK %s", 200, payload)
}

func (o *PriceLayerTypesUpdateOK) GetPayload() *models.PriceLayerType {
	return o.Payload
}

func (o *PriceLayerTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriceLayerType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPriceLayerTypesUpdateDefault creates a PriceLayerTypesUpdateDefault with default headers values
func NewPriceLayerTypesUpdateDefault(code int) *PriceLayerTypesUpdateDefault {
	return &PriceLayerTypesUpdateDefault{
		_statusCode: code,
	}
}

/*
PriceLayerTypesUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PriceLayerTypesUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this price layer types update default response has a 2xx status code
func (o *PriceLayerTypesUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this price layer types update default response has a 3xx status code
func (o *PriceLayerTypesUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this price layer types update default response has a 4xx status code
func (o *PriceLayerTypesUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this price layer types update default response has a 5xx status code
func (o *PriceLayerTypesUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this price layer types update default response a status code equal to that given
func (o *PriceLayerTypesUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the price layer types update default response
func (o *PriceLayerTypesUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PriceLayerTypesUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PriceLayerTypes/{id}][%d] PriceLayerTypes_Update default %s", o._statusCode, payload)
}

func (o *PriceLayerTypesUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PriceLayerTypes/{id}][%d] PriceLayerTypes_Update default %s", o._statusCode, payload)
}

func (o *PriceLayerTypesUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PriceLayerTypesUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
