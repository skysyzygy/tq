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

// PricingRuleMessageTypesUpdateReader is a Reader for the PricingRuleMessageTypesUpdate structure.
type PricingRuleMessageTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingRuleMessageTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingRuleMessageTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /ReferenceData/PricingRuleMessageTypes/{id}] PricingRuleMessageTypes_Update", response, response.Code())
	}
}

// NewPricingRuleMessageTypesUpdateOK creates a PricingRuleMessageTypesUpdateOK with default headers values
func NewPricingRuleMessageTypesUpdateOK() *PricingRuleMessageTypesUpdateOK {
	return &PricingRuleMessageTypesUpdateOK{}
}

/*
PricingRuleMessageTypesUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PricingRuleMessageTypesUpdateOK struct {
	Payload *models.PricingRuleMessageType
}

// IsSuccess returns true when this pricing rule message types update o k response has a 2xx status code
func (o *PricingRuleMessageTypesUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing rule message types update o k response has a 3xx status code
func (o *PricingRuleMessageTypesUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing rule message types update o k response has a 4xx status code
func (o *PricingRuleMessageTypesUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing rule message types update o k response has a 5xx status code
func (o *PricingRuleMessageTypesUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing rule message types update o k response a status code equal to that given
func (o *PricingRuleMessageTypesUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pricing rule message types update o k response
func (o *PricingRuleMessageTypesUpdateOK) Code() int {
	return 200
}

func (o *PricingRuleMessageTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ReferenceData/PricingRuleMessageTypes/{id}][%d] pricingRuleMessageTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *PricingRuleMessageTypesUpdateOK) String() string {
	return fmt.Sprintf("[PUT /ReferenceData/PricingRuleMessageTypes/{id}][%d] pricingRuleMessageTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *PricingRuleMessageTypesUpdateOK) GetPayload() *models.PricingRuleMessageType {
	return o.Payload
}

func (o *PricingRuleMessageTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PricingRuleMessageType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
