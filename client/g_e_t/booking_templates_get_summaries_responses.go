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

// BookingTemplatesGetSummariesReader is a Reader for the BookingTemplatesGetSummaries structure.
type BookingTemplatesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BookingTemplatesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBookingTemplatesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBookingTemplatesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBookingTemplatesGetSummariesOK creates a BookingTemplatesGetSummariesOK with default headers values
func NewBookingTemplatesGetSummariesOK() *BookingTemplatesGetSummariesOK {
	return &BookingTemplatesGetSummariesOK{}
}

/*
BookingTemplatesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type BookingTemplatesGetSummariesOK struct {
	Payload []*models.BookingTemplateSummary
}

// IsSuccess returns true when this booking templates get summaries o k response has a 2xx status code
func (o *BookingTemplatesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this booking templates get summaries o k response has a 3xx status code
func (o *BookingTemplatesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this booking templates get summaries o k response has a 4xx status code
func (o *BookingTemplatesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this booking templates get summaries o k response has a 5xx status code
func (o *BookingTemplatesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this booking templates get summaries o k response a status code equal to that given
func (o *BookingTemplatesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the booking templates get summaries o k response
func (o *BookingTemplatesGetSummariesOK) Code() int {
	return 200
}

func (o *BookingTemplatesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/BookingTemplates/Summary][%d] bookingTemplatesGetSummariesOK %s", 200, payload)
}

func (o *BookingTemplatesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/BookingTemplates/Summary][%d] bookingTemplatesGetSummariesOK %s", 200, payload)
}

func (o *BookingTemplatesGetSummariesOK) GetPayload() []*models.BookingTemplateSummary {
	return o.Payload
}

func (o *BookingTemplatesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBookingTemplatesGetSummariesDefault creates a BookingTemplatesGetSummariesDefault with default headers values
func NewBookingTemplatesGetSummariesDefault(code int) *BookingTemplatesGetSummariesDefault {
	return &BookingTemplatesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
BookingTemplatesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type BookingTemplatesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this booking templates get summaries default response has a 2xx status code
func (o *BookingTemplatesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this booking templates get summaries default response has a 3xx status code
func (o *BookingTemplatesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this booking templates get summaries default response has a 4xx status code
func (o *BookingTemplatesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this booking templates get summaries default response has a 5xx status code
func (o *BookingTemplatesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this booking templates get summaries default response a status code equal to that given
func (o *BookingTemplatesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the booking templates get summaries default response
func (o *BookingTemplatesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *BookingTemplatesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/BookingTemplates/Summary][%d] BookingTemplates_GetSummaries default %s", o._statusCode, payload)
}

func (o *BookingTemplatesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /EventsManagement/BookingTemplates/Summary][%d] BookingTemplates_GetSummaries default %s", o._statusCode, payload)
}

func (o *BookingTemplatesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *BookingTemplatesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
