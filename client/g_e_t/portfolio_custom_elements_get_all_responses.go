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

// PortfolioCustomElementsGetAllReader is a Reader for the PortfolioCustomElementsGetAll structure.
type PortfolioCustomElementsGetAllReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortfolioCustomElementsGetAllReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortfolioCustomElementsGetAllOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortfolioCustomElementsGetAllDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortfolioCustomElementsGetAllOK creates a PortfolioCustomElementsGetAllOK with default headers values
func NewPortfolioCustomElementsGetAllOK() *PortfolioCustomElementsGetAllOK {
	return &PortfolioCustomElementsGetAllOK{}
}

/*
PortfolioCustomElementsGetAllOK describes a response with status code 200, with default header values.

OK
*/
type PortfolioCustomElementsGetAllOK struct {
	Payload []*models.PortfolioCustomElement
}

// IsSuccess returns true when this portfolio custom elements get all o k response has a 2xx status code
func (o *PortfolioCustomElementsGetAllOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portfolio custom elements get all o k response has a 3xx status code
func (o *PortfolioCustomElementsGetAllOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portfolio custom elements get all o k response has a 4xx status code
func (o *PortfolioCustomElementsGetAllOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this portfolio custom elements get all o k response has a 5xx status code
func (o *PortfolioCustomElementsGetAllOK) IsServerError() bool {
	return false
}

// IsCode returns true when this portfolio custom elements get all o k response a status code equal to that given
func (o *PortfolioCustomElementsGetAllOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the portfolio custom elements get all o k response
func (o *PortfolioCustomElementsGetAllOK) Code() int {
	return 200
}

func (o *PortfolioCustomElementsGetAllOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PortfolioCustomElements][%d] portfolioCustomElementsGetAllOK %s", 200, payload)
}

func (o *PortfolioCustomElementsGetAllOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PortfolioCustomElements][%d] portfolioCustomElementsGetAllOK %s", 200, payload)
}

func (o *PortfolioCustomElementsGetAllOK) GetPayload() []*models.PortfolioCustomElement {
	return o.Payload
}

func (o *PortfolioCustomElementsGetAllOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortfolioCustomElementsGetAllDefault creates a PortfolioCustomElementsGetAllDefault with default headers values
func NewPortfolioCustomElementsGetAllDefault(code int) *PortfolioCustomElementsGetAllDefault {
	return &PortfolioCustomElementsGetAllDefault{
		_statusCode: code,
	}
}

/*
PortfolioCustomElementsGetAllDefault describes a response with status code -1, with default header values.

Error
*/
type PortfolioCustomElementsGetAllDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this portfolio custom elements get all default response has a 2xx status code
func (o *PortfolioCustomElementsGetAllDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this portfolio custom elements get all default response has a 3xx status code
func (o *PortfolioCustomElementsGetAllDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this portfolio custom elements get all default response has a 4xx status code
func (o *PortfolioCustomElementsGetAllDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this portfolio custom elements get all default response has a 5xx status code
func (o *PortfolioCustomElementsGetAllDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this portfolio custom elements get all default response a status code equal to that given
func (o *PortfolioCustomElementsGetAllDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the portfolio custom elements get all default response
func (o *PortfolioCustomElementsGetAllDefault) Code() int {
	return o._statusCode
}

func (o *PortfolioCustomElementsGetAllDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PortfolioCustomElements][%d] PortfolioCustomElements_GetAll default %s", o._statusCode, payload)
}

func (o *PortfolioCustomElementsGetAllDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/PortfolioCustomElements][%d] PortfolioCustomElements_GetAll default %s", o._statusCode, payload)
}

func (o *PortfolioCustomElementsGetAllDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PortfolioCustomElementsGetAllDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
