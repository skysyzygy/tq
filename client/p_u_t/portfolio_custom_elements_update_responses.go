// Code generated by go-swagger; DO NOT EDIT.

package p_u_t

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

// PortfolioCustomElementsUpdateReader is a Reader for the PortfolioCustomElementsUpdate structure.
type PortfolioCustomElementsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortfolioCustomElementsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortfolioCustomElementsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortfolioCustomElementsUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortfolioCustomElementsUpdateOK creates a PortfolioCustomElementsUpdateOK with default headers values
func NewPortfolioCustomElementsUpdateOK() *PortfolioCustomElementsUpdateOK {
	return &PortfolioCustomElementsUpdateOK{}
}

/*
PortfolioCustomElementsUpdateOK describes a response with status code 200, with default header values.

OK
*/
type PortfolioCustomElementsUpdateOK struct {
	Payload *models.PortfolioCustomElement
}

// IsSuccess returns true when this portfolio custom elements update o k response has a 2xx status code
func (o *PortfolioCustomElementsUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portfolio custom elements update o k response has a 3xx status code
func (o *PortfolioCustomElementsUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portfolio custom elements update o k response has a 4xx status code
func (o *PortfolioCustomElementsUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this portfolio custom elements update o k response has a 5xx status code
func (o *PortfolioCustomElementsUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this portfolio custom elements update o k response a status code equal to that given
func (o *PortfolioCustomElementsUpdateOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the portfolio custom elements update o k response
func (o *PortfolioCustomElementsUpdateOK) Code() int {
	return 200
}

func (o *PortfolioCustomElementsUpdateOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PortfolioCustomElements/{id}][%d] portfolioCustomElementsUpdateOK %s", 200, payload)
}

func (o *PortfolioCustomElementsUpdateOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PortfolioCustomElements/{id}][%d] portfolioCustomElementsUpdateOK %s", 200, payload)
}

func (o *PortfolioCustomElementsUpdateOK) GetPayload() *models.PortfolioCustomElement {
	return o.Payload
}

func (o *PortfolioCustomElementsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortfolioCustomElement)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortfolioCustomElementsUpdateDefault creates a PortfolioCustomElementsUpdateDefault with default headers values
func NewPortfolioCustomElementsUpdateDefault(code int) *PortfolioCustomElementsUpdateDefault {
	return &PortfolioCustomElementsUpdateDefault{
		_statusCode: code,
	}
}

/*
PortfolioCustomElementsUpdateDefault describes a response with status code -1, with default header values.

Error
*/
type PortfolioCustomElementsUpdateDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this portfolio custom elements update default response has a 2xx status code
func (o *PortfolioCustomElementsUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this portfolio custom elements update default response has a 3xx status code
func (o *PortfolioCustomElementsUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this portfolio custom elements update default response has a 4xx status code
func (o *PortfolioCustomElementsUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this portfolio custom elements update default response has a 5xx status code
func (o *PortfolioCustomElementsUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this portfolio custom elements update default response a status code equal to that given
func (o *PortfolioCustomElementsUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the portfolio custom elements update default response
func (o *PortfolioCustomElementsUpdateDefault) Code() int {
	return o._statusCode
}

func (o *PortfolioCustomElementsUpdateDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PortfolioCustomElements/{id}][%d] PortfolioCustomElements_Update default %s", o._statusCode, payload)
}

func (o *PortfolioCustomElementsUpdateDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[PUT /ReferenceData/PortfolioCustomElements/{id}][%d] PortfolioCustomElements_Update default %s", o._statusCode, payload)
}

func (o *PortfolioCustomElementsUpdateDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PortfolioCustomElementsUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
