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
		result := NewPricingRulesGetAllSummaryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
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
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRules/Summary][%d] pricingRulesGetAllSummaryOK %s", 200, payload)
}

func (o *PricingRulesGetAllSummaryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRules/Summary][%d] pricingRulesGetAllSummaryOK %s", 200, payload)
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

// NewPricingRulesGetAllSummaryDefault creates a PricingRulesGetAllSummaryDefault with default headers values
func NewPricingRulesGetAllSummaryDefault(code int) *PricingRulesGetAllSummaryDefault {
	return &PricingRulesGetAllSummaryDefault{
		_statusCode: code,
	}
}

/*
PricingRulesGetAllSummaryDefault describes a response with status code -1, with default header values.

Error
*/
type PricingRulesGetAllSummaryDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this pricing rules get all summary default response has a 2xx status code
func (o *PricingRulesGetAllSummaryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing rules get all summary default response has a 3xx status code
func (o *PricingRulesGetAllSummaryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing rules get all summary default response has a 4xx status code
func (o *PricingRulesGetAllSummaryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing rules get all summary default response has a 5xx status code
func (o *PricingRulesGetAllSummaryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing rules get all summary default response a status code equal to that given
func (o *PricingRulesGetAllSummaryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pricing rules get all summary default response
func (o *PricingRulesGetAllSummaryDefault) Code() int {
	return o._statusCode
}

func (o *PricingRulesGetAllSummaryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRules/Summary][%d] PricingRules_GetAllSummary default %s", o._statusCode, payload)
}

func (o *PricingRulesGetAllSummaryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRules/Summary][%d] PricingRules_GetAllSummary default %s", o._statusCode, payload)
}

func (o *PricingRulesGetAllSummaryDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PricingRulesGetAllSummaryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
