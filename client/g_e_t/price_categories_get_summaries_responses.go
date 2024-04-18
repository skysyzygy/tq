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

// PriceCategoriesGetSummariesReader is a Reader for the PriceCategoriesGetSummaries structure.
type PriceCategoriesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PriceCategoriesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPriceCategoriesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PriceCategories/Summary] PriceCategories_GetSummaries", response, response.Code())
	}
}

// NewPriceCategoriesGetSummariesOK creates a PriceCategoriesGetSummariesOK with default headers values
func NewPriceCategoriesGetSummariesOK() *PriceCategoriesGetSummariesOK {
	return &PriceCategoriesGetSummariesOK{}
}

/*
PriceCategoriesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type PriceCategoriesGetSummariesOK struct {
	Payload []*models.PriceCategorySummary
}

// IsSuccess returns true when this price categories get summaries o k response has a 2xx status code
func (o *PriceCategoriesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this price categories get summaries o k response has a 3xx status code
func (o *PriceCategoriesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this price categories get summaries o k response has a 4xx status code
func (o *PriceCategoriesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this price categories get summaries o k response has a 5xx status code
func (o *PriceCategoriesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this price categories get summaries o k response a status code equal to that given
func (o *PriceCategoriesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the price categories get summaries o k response
func (o *PriceCategoriesGetSummariesOK) Code() int {
	return 200
}

func (o *PriceCategoriesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories/Summary][%d] priceCategoriesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *PriceCategoriesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PriceCategories/Summary][%d] priceCategoriesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *PriceCategoriesGetSummariesOK) GetPayload() []*models.PriceCategorySummary {
	return o.Payload
}

func (o *PriceCategoriesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}