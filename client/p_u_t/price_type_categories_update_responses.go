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

// PriceTypeCategoriesUpdateReader is a Reader for the PriceTypeCategoriesUpdate structure.
type PriceTypeCategoriesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceTypeCategoriesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceTypeCategoriesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/PriceTypeCategories/{id}] PriceTypeCategories_Update", response, response.Code())
	}
}

// NewPriceTypeCategoriesUpdateOK creates a PriceTypeCategoriesUpdateOK with default headers values
func NewPriceTypeCategoriesUpdateOK() *PriceTypeCategoriesUpdateOK {
	return &PriceTypeCategoriesUpdateOK{}
}

/*
PriceTypeCategoriesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PriceTypeCategoriesUpdateOK struct {
	Payload *models.PriceTypeCategory
}

// IsSuccess returns true when this price type categories update o k response has a 2xx status code
func (o *PriceTypeCategoriesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price type categories update o k response has a 3xx status code
func (o *PriceTypeCategoriesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price type categories update o k response has a 4xx status code
func (o *PriceTypeCategoriesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price type categories update o k response has a 5xx status code
func (o *PriceTypeCategoriesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price type categories update o k response a status code equal to that given
func (o *PriceTypeCategoriesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price type categories update o k response
func (o *PriceTypeCategoriesUpdateOK) Code() int {
	return 200
}

func (o *PriceTypeCategoriesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/PriceTypeCategories/{id}][%d] priceTypeCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *PriceTypeCategoriesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/PriceTypeCategories/{id}][%d] priceTypeCategoriesUpdateOK  %+v", 200, o.Payload)
}

func (o *PriceTypeCategoriesUpdateOK) GetPayload() *models.PriceTypeCategory {
	return o.Payload
}

func (o *PriceTypeCategoriesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PriceTypeCategory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
