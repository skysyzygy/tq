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

// PriceCategoriesGetReader is a Reader for the PriceCategoriesGet structure.
type PriceCategoriesGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceCategoriesGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceCategoriesGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PriceCategories/{id}] PriceCategories_Get", response, response.Code())
	}
}

// NewPriceCategoriesGetOK creates a PriceCategoriesGetOK with default headers values
func NewPriceCategoriesGetOK() *PriceCategoriesGetOK {
	return &PriceCategoriesGetOK{}
}

/*
PriceCategoriesGetOK describes a response with status code 200, with default header values.

OK
*/
type PriceCategoriesGetOK struct {
	Payload *models.PriceCategory
}

// IsSuccess returns true when this price categories get o k response has a 2xx status code
func (o *PriceCategoriesGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price categories get o k response has a 3xx status code
func (o *PriceCategoriesGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price categories get o k response has a 4xx status code
func (o *PriceCategoriesGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price categories get o k response has a 5xx status code
func (o *PriceCategoriesGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price categories get o k response a status code equal to that given
func (o *PriceCategoriesGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price categories get o k response
func (o *PriceCategoriesGetOK) Code() int {
	return 200
}

func (o *PriceCategoriesGetOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories/{id}][%d] priceCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *PriceCategoriesGetOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories/{id}][%d] priceCategoriesGetOK  %+v", 200, o.Payload)
}

func (o *PriceCategoriesGetOK) GetPayload() *models.PriceCategory {
	return o.Payload
}

func (o *PriceCategoriesGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriceCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
