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

// PortfoliosGetPortfolioEntryReader is a Reader for the PortfoliosGetPortfolioEntry structure.
type PortfoliosGetPortfolioEntryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortfoliosGetPortfolioEntryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPortfoliosGetPortfolioEntryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortfoliosGetPortfolioEntryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortfoliosGetPortfolioEntryOK creates a PortfoliosGetPortfolioEntryOK with default headers values
func NewPortfoliosGetPortfolioEntryOK() *PortfoliosGetPortfolioEntryOK {
	return &PortfoliosGetPortfolioEntryOK{}
}

/*
PortfoliosGetPortfolioEntryOK describes a response with status code 200, with default header values.

OK
*/
type PortfoliosGetPortfolioEntryOK struct {
	Payload *models.PortfolioPlan
}

// IsSuccess returns true when this portfolios get portfolio entry o k response has a 2xx status code
func (o *PortfoliosGetPortfolioEntryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portfolios get portfolio entry o k response has a 3xx status code
func (o *PortfoliosGetPortfolioEntryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portfolios get portfolio entry o k response has a 4xx status code
func (o *PortfoliosGetPortfolioEntryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this portfolios get portfolio entry o k response has a 5xx status code
func (o *PortfoliosGetPortfolioEntryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this portfolios get portfolio entry o k response a status code equal to that given
func (o *PortfoliosGetPortfolioEntryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the portfolios get portfolio entry o k response
func (o *PortfoliosGetPortfolioEntryOK) Code() int {
	return 200
}

func (o *PortfoliosGetPortfolioEntryOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Portfolios/Entries][%d] portfoliosGetPortfolioEntryOK %s", 200, payload)
}

func (o *PortfoliosGetPortfolioEntryOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Portfolios/Entries][%d] portfoliosGetPortfolioEntryOK %s", 200, payload)
}

func (o *PortfoliosGetPortfolioEntryOK) GetPayload() *models.PortfolioPlan {
	return o.Payload
}

func (o *PortfoliosGetPortfolioEntryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortfolioPlan)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortfoliosGetPortfolioEntryDefault creates a PortfoliosGetPortfolioEntryDefault with default headers values
func NewPortfoliosGetPortfolioEntryDefault(code int) *PortfoliosGetPortfolioEntryDefault {
	return &PortfoliosGetPortfolioEntryDefault{
		_statusCode: code,
	}
}

/*
PortfoliosGetPortfolioEntryDefault describes a response with status code -1, with default header values.

Error
*/
type PortfoliosGetPortfolioEntryDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this portfolios get portfolio entry default response has a 2xx status code
func (o *PortfoliosGetPortfolioEntryDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this portfolios get portfolio entry default response has a 3xx status code
func (o *PortfoliosGetPortfolioEntryDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this portfolios get portfolio entry default response has a 4xx status code
func (o *PortfoliosGetPortfolioEntryDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this portfolios get portfolio entry default response has a 5xx status code
func (o *PortfoliosGetPortfolioEntryDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this portfolios get portfolio entry default response a status code equal to that given
func (o *PortfoliosGetPortfolioEntryDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the portfolios get portfolio entry default response
func (o *PortfoliosGetPortfolioEntryDefault) Code() int {
	return o._statusCode
}

func (o *PortfoliosGetPortfolioEntryDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Portfolios/Entries][%d] Portfolios_GetPortfolioEntry default %s", o._statusCode, payload)
}

func (o *PortfoliosGetPortfolioEntryDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /Finance/Portfolios/Entries][%d] Portfolios_GetPortfolioEntry default %s", o._statusCode, payload)
}

func (o *PortfoliosGetPortfolioEntryDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *PortfoliosGetPortfolioEntryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
