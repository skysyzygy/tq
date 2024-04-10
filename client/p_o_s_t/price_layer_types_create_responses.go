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

// PriceLayerTypesCreateReader is a Reader for the PriceLayerTypesCreate structure.
type PriceLayerTypesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceLayerTypesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceLayerTypesCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /ReferenceData/PriceLayerTypes] PriceLayerTypes_Create", response, response.Code())
	}
}

// NewPriceLayerTypesCreateOK creates a PriceLayerTypesCreateOK with default headers values
func NewPriceLayerTypesCreateOK() *PriceLayerTypesCreateOK {
	return &PriceLayerTypesCreateOK{}
}

/*
PriceLayerTypesCreateOK describes a response with status code 200, with default header values.

OK
*/
type PriceLayerTypesCreateOK struct {
	Payload *models.PriceLayerType
}

// IsSuccess returns true when this price layer types create o k response has a 2xx status code
func (o *PriceLayerTypesCreateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price layer types create o k response has a 3xx status code
func (o *PriceLayerTypesCreateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price layer types create o k response has a 4xx status code
func (o *PriceLayerTypesCreateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price layer types create o k response has a 5xx status code
func (o *PriceLayerTypesCreateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price layer types create o k response a status code equal to that given
func (o *PriceLayerTypesCreateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price layer types create o k response
func (o *PriceLayerTypesCreateOK) Code() int {
	return 200
}

func (o *PriceLayerTypesCreateOK) Error() string {
	return fmt.Sprintf("[POST /ReferenceData/PriceLayerTypes][%d] priceLayerTypesCreateOK  %+v", 200, o.Payload)
}

func (o *PriceLayerTypesCreateOK) String() string {
	return fmt.Sprintf("[POST /ReferenceData/PriceLayerTypes][%d] priceLayerTypesCreateOK  %+v", 200, o.Payload)
}

func (o *PriceLayerTypesCreateOK) GetPayload() *models.PriceLayerType {
	return o.Payload
}

func (o *PriceLayerTypesCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriceLayerType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
