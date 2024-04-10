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

// DiscountTypesGetSummariesReader is a Reader for the DiscountTypesGetSummaries structure.
type DiscountTypesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiscountTypesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiscountTypesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/DiscountTypes/Summary] DiscountTypes_GetSummaries", response, response.Code())
	}
}

// NewDiscountTypesGetSummariesOK creates a DiscountTypesGetSummariesOK with default headers values
func NewDiscountTypesGetSummariesOK() *DiscountTypesGetSummariesOK {
	return &DiscountTypesGetSummariesOK{}
}

/*
DiscountTypesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type DiscountTypesGetSummariesOK struct {
	Payload []*models.DiscountTypeSummary
}

// IsSuccess returns true when this discount types get summaries o k response has a 2xx status code
func (o *DiscountTypesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this discount types get summaries o k response has a 3xx status code
func (o *DiscountTypesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this discount types get summaries o k response has a 4xx status code
func (o *DiscountTypesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this discount types get summaries o k response has a 5xx status code
func (o *DiscountTypesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this discount types get summaries o k response a status code equal to that given
func (o *DiscountTypesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the discount types get summaries o k response
func (o *DiscountTypesGetSummariesOK) Code() int {
	return 200
}

func (o *DiscountTypesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/DiscountTypes/Summary][%d] discountTypesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *DiscountTypesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/DiscountTypes/Summary][%d] discountTypesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *DiscountTypesGetSummariesOK) GetPayload() []*models.DiscountTypeSummary {
	return o.Payload
}

func (o *DiscountTypesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
