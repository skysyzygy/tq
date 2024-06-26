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

// PricingRuleSetsGetAllReader is a Reader for the PricingRuleSetsGetAll structure.
type PricingRuleSetsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PricingRuleSetsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPricingRuleSetsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPricingRuleSetsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPricingRuleSetsGetAllOK creates a PricingRuleSetsGetAllOK with default headers values
func NewPricingRuleSetsGetAllOK() *PricingRuleSetsGetAllOK {
	return &PricingRuleSetsGetAllOK{}
}

/*
PricingRuleSetsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PricingRuleSetsGetAllOK struct {
	Payload []*models.PricingRuleSet
}

// IsSuccess returns true when this pricing rule sets get all o k response has a 2xx status code
func (o *PricingRuleSetsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pricing rule sets get all o k response has a 3xx status code
func (o *PricingRuleSetsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pricing rule sets get all o k response has a 4xx status code
func (o *PricingRuleSetsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pricing rule sets get all o k response has a 5xx status code
func (o *PricingRuleSetsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pricing rule sets get all o k response a status code equal to that given
func (o *PricingRuleSetsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the pricing rule sets get all o k response
func (o *PricingRuleSetsGetAllOK) Code() int {
	return 200
}

func (o *PricingRuleSetsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRuleSets][%d] pricingRuleSetsGetAllOK %s", 200, payload)
}

func (o *PricingRuleSetsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRuleSets][%d] pricingRuleSetsGetAllOK %s", 200, payload)
}

func (o *PricingRuleSetsGetAllOK) GetPayload() []*models.PricingRuleSet {
	return o.Payload
}

func (o *PricingRuleSetsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPricingRuleSetsGetAllDefault creates a PricingRuleSetsGetAllDefault with default headers values
func NewPricingRuleSetsGetAllDefault(code int) *PricingRuleSetsGetAllDefault {
	return &PricingRuleSetsGetAllDefault{
		_statusCode: code,
	}
}

/*
PricingRuleSetsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type PricingRuleSetsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this pricing rule sets get all default response has a 2xx status code
func (o *PricingRuleSetsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this pricing rule sets get all default response has a 3xx status code
func (o *PricingRuleSetsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this pricing rule sets get all default response has a 4xx status code
func (o *PricingRuleSetsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this pricing rule sets get all default response has a 5xx status code
func (o *PricingRuleSetsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this pricing rule sets get all default response a status code equal to that given
func (o *PricingRuleSetsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the pricing rule sets get all default response
func (o *PricingRuleSetsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *PricingRuleSetsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRuleSets][%d] PricingRuleSets_GetAll default %s", o._statusCode, payload)
}

func (o *PricingRuleSetsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /TXN/PricingRuleSets][%d] PricingRuleSets_GetAll default %s", o._statusCode, payload)
}

func (o *PricingRuleSetsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PricingRuleSetsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
