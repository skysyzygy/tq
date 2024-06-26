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

// InterestCategoriesGetSummariesReader is a Reader for the InterestCategoriesGetSummaries structure.
type InterestCategoriesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InterestCategoriesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInterestCategoriesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewInterestCategoriesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewInterestCategoriesGetSummariesOK creates a InterestCategoriesGetSummariesOK with default headers values
func NewInterestCategoriesGetSummariesOK() *InterestCategoriesGetSummariesOK {
	return &InterestCategoriesGetSummariesOK{}
}

/*
InterestCategoriesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type InterestCategoriesGetSummariesOK struct {
	Payload []*models.InterestCategorySummary
}

// IsSuccess returns true when this interest categories get summaries o k response has a 2xx status code
func (o *InterestCategoriesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this interest categories get summaries o k response has a 3xx status code
func (o *InterestCategoriesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this interest categories get summaries o k response has a 4xx status code
func (o *InterestCategoriesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this interest categories get summaries o k response has a 5xx status code
func (o *InterestCategoriesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this interest categories get summaries o k response a status code equal to that given
func (o *InterestCategoriesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the interest categories get summaries o k response
func (o *InterestCategoriesGetSummariesOK) Code() int {
	return 200
}

func (o *InterestCategoriesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestCategories/Summary][%d] interestCategoriesGetSummariesOK %s", 200, payload)
}

func (o *InterestCategoriesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestCategories/Summary][%d] interestCategoriesGetSummariesOK %s", 200, payload)
}

func (o *InterestCategoriesGetSummariesOK) GetPayload() []*models.InterestCategorySummary {
	return o.Payload
}

func (o *InterestCategoriesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInterestCategoriesGetSummariesDefault creates a InterestCategoriesGetSummariesDefault with default headers values
func NewInterestCategoriesGetSummariesDefault(code int) *InterestCategoriesGetSummariesDefault {
	return &InterestCategoriesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
InterestCategoriesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type InterestCategoriesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this interest categories get summaries default response has a 2xx status code
func (o *InterestCategoriesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this interest categories get summaries default response has a 3xx status code
func (o *InterestCategoriesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this interest categories get summaries default response has a 4xx status code
func (o *InterestCategoriesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this interest categories get summaries default response has a 5xx status code
func (o *InterestCategoriesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this interest categories get summaries default response a status code equal to that given
func (o *InterestCategoriesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the interest categories get summaries default response
func (o *InterestCategoriesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *InterestCategoriesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestCategories/Summary][%d] InterestCategories_GetSummaries default %s", o._statusCode, payload)
}

func (o *InterestCategoriesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/InterestCategories/Summary][%d] InterestCategories_GetSummaries default %s", o._statusCode, payload)
}

func (o *InterestCategoriesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *InterestCategoriesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
