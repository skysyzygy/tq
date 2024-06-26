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

// CountriesGetSummariesReader is a Reader for the CountriesGetSummaries structure.
type CountriesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CountriesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCountriesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCountriesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCountriesGetSummariesOK creates a CountriesGetSummariesOK with default headers values
func NewCountriesGetSummariesOK() *CountriesGetSummariesOK {
	return &CountriesGetSummariesOK{}
}

/*
CountriesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type CountriesGetSummariesOK struct {
	Payload []*models.CountrySummary
}

// IsSuccess returns true when this countries get summaries o k response has a 2xx status code
func (o *CountriesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this countries get summaries o k response has a 3xx status code
func (o *CountriesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this countries get summaries o k response has a 4xx status code
func (o *CountriesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this countries get summaries o k response has a 5xx status code
func (o *CountriesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this countries get summaries o k response a status code equal to that given
func (o *CountriesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the countries get summaries o k response
func (o *CountriesGetSummariesOK) Code() int {
	return 200
}

func (o *CountriesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Countries/Summary][%d] countriesGetSummariesOK %s", 200, payload)
}

func (o *CountriesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Countries/Summary][%d] countriesGetSummariesOK %s", 200, payload)
}

func (o *CountriesGetSummariesOK) GetPayload() []*models.CountrySummary {
	return o.Payload
}

func (o *CountriesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCountriesGetSummariesDefault creates a CountriesGetSummariesDefault with default headers values
func NewCountriesGetSummariesDefault(code int) *CountriesGetSummariesDefault {
	return &CountriesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
CountriesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type CountriesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this countries get summaries default response has a 2xx status code
func (o *CountriesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this countries get summaries default response has a 3xx status code
func (o *CountriesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this countries get summaries default response has a 4xx status code
func (o *CountriesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this countries get summaries default response has a 5xx status code
func (o *CountriesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this countries get summaries default response a status code equal to that given
func (o *CountriesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the countries get summaries default response
func (o *CountriesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *CountriesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Countries/Summary][%d] Countries_GetSummaries default %s", o._statusCode, payload)
}

func (o *CountriesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Countries/Summary][%d] Countries_GetSummaries default %s", o._statusCode, payload)
}

func (o *CountriesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *CountriesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
