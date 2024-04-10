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

// PricingRulesGetAllSummaryReader is a Reader for the PricingRulesGetAllSummary structure.
type PricingRulesGetAllSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingRulesGetAllSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingRulesGetAllSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/PricingRules/Summary] PricingRules_GetAllSummary", response, response.Code())
	}
}

// NewPricingRulesGetAllSummaryOK creates a PricingRulesGetAllSummaryOK with default headers values
func NewPricingRulesGetAllSummaryOK() *PricingRulesGetAllSummaryOK {
	return &PricingRulesGetAllSummaryOK{}
}

/*
PricingRulesGetAllSummaryOK describes a response with status code 200, with default header values.

OK
*/
type PricingRulesGetAllSummaryOK struct {
	Payload []*models.PricingRuleSummary
}

// IsSuccess returns true when this pricing rules get all summary o k response has a 2xx status code
func (o *PricingRulesGetAllSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing rules get all summary o k response has a 3xx status code
func (o *PricingRulesGetAllSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing rules get all summary o k response has a 4xx status code
func (o *PricingRulesGetAllSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing rules get all summary o k response has a 5xx status code
func (o *PricingRulesGetAllSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing rules get all summary o k response a status code equal to that given
func (o *PricingRulesGetAllSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pricing rules get all summary o k response
func (o *PricingRulesGetAllSummaryOK) Code() int {
	return 200
}

func (o *PricingRulesGetAllSummaryOK) Error() string {
	return fmt.Sprintf("[GET /TXN/PricingRules/Summary][%d] pricingRulesGetAllSummaryOK  %+v", 200, o.Payload)
}

func (o *PricingRulesGetAllSummaryOK) String() string {
	return fmt.Sprintf("[GET /TXN/PricingRules/Summary][%d] pricingRulesGetAllSummaryOK  %+v", 200, o.Payload)
}

func (o *PricingRulesGetAllSummaryOK) GetPayload() []*models.PricingRuleSummary {
	return o.Payload
}

func (o *PricingRulesGetAllSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
