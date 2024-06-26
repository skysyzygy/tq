// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

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

// CartAddContributionReader is a Reader for the CartAddContribution structure.
type CartAddContributionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CartAddContributionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCartAddContributionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCartAddContributionDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCartAddContributionOK creates a CartAddContributionOK with default headers values
func NewCartAddContributionOK() *CartAddContributionOK {
	return &CartAddContributionOK{}
}

/*
CartAddContributionOK describes a response with status code 200, with default header values.

OK
*/
type CartAddContributionOK struct {
	Payload *models.AddContributionResponse
}

// IsSuccess returns true when this cart add contribution o k response has a 2xx status code
func (o *CartAddContributionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cart add contribution o k response has a 3xx status code
func (o *CartAddContributionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cart add contribution o k response has a 4xx status code
func (o *CartAddContributionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cart add contribution o k response has a 5xx status code
func (o *CartAddContributionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cart add contribution o k response a status code equal to that given
func (o *CartAddContributionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the cart add contribution o k response
func (o *CartAddContributionOK) Code() int {
	return 200
}

func (o *CartAddContributionOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Contributions][%d] cartAddContributionOK %s", 200, payload)
}

func (o *CartAddContributionOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Contributions][%d] cartAddContributionOK %s", 200, payload)
}

func (o *CartAddContributionOK) GetPayload() *models.AddContributionResponse {
	return o.Payload
}

func (o *CartAddContributionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AddContributionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCartAddContributionDefault creates a CartAddContributionDefault with default headers values
func NewCartAddContributionDefault(code int) *CartAddContributionDefault {
	return &CartAddContributionDefault{
		_statusCode: code,
	}
}

/*
CartAddContributionDefault describes a response with status code -1, with default header values.

Error
*/
type CartAddContributionDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this cart add contribution default response has a 2xx status code
func (o *CartAddContributionDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cart add contribution default response has a 3xx status code
func (o *CartAddContributionDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cart add contribution default response has a 4xx status code
func (o *CartAddContributionDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cart add contribution default response has a 5xx status code
func (o *CartAddContributionDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cart add contribution default response a status code equal to that given
func (o *CartAddContributionDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the cart add contribution default response
func (o *CartAddContributionDefault) Code() int {
	return o._statusCode
}

func (o *CartAddContributionDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Contributions][%d] Cart_AddContribution default %s", o._statusCode, payload)
}

func (o *CartAddContributionDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /Web/Cart/{sessionKey}/Contributions][%d] Cart_AddContribution default %s", o._statusCode, payload)
}

func (o *CartAddContributionDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CartAddContributionDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
