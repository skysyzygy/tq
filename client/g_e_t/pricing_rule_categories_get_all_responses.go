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

// PricingRuleCategoriesGetAllReader is a Reader for the PricingRuleCategoriesGetAll structure.
type PricingRuleCategoriesGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingRuleCategoriesGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingRuleCategoriesGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/PricingRuleCategories] PricingRuleCategories_GetAll", response, response.Code())
	}
}

// NewPricingRuleCategoriesGetAllOK creates a PricingRuleCategoriesGetAllOK with default headers values
func NewPricingRuleCategoriesGetAllOK() *PricingRuleCategoriesGetAllOK {
	return &PricingRuleCategoriesGetAllOK{}
}

/*
PricingRuleCategoriesGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PricingRuleCategoriesGetAllOK struct {
	Payload []*models.PricingRuleCategory
}

// IsSuccess returns true when this pricing rule categories get all o k response has a 2xx status code
func (o *PricingRuleCategoriesGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing rule categories get all o k response has a 3xx status code
func (o *PricingRuleCategoriesGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing rule categories get all o k response has a 4xx status code
func (o *PricingRuleCategoriesGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing rule categories get all o k response has a 5xx status code
func (o *PricingRuleCategoriesGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing rule categories get all o k response a status code equal to that given
func (o *PricingRuleCategoriesGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pricing rule categories get all o k response
func (o *PricingRuleCategoriesGetAllOK) Code() int {
	return 200
}

func (o *PricingRuleCategoriesGetAllOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/PricingRuleCategories][%d] pricingRuleCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *PricingRuleCategoriesGetAllOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/PricingRuleCategories][%d] pricingRuleCategoriesGetAllOK  %+v", 200, o.Payload)
}

func (o *PricingRuleCategoriesGetAllOK) GetPayload() []*models.PricingRuleCategory {
	return o.Payload
}

func (o *PricingRuleCategoriesGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
