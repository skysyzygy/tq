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

// SuffixesGetSummariesReader is a Reader for the SuffixesGetSummaries structure.
type SuffixesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SuffixesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSuffixesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSuffixesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSuffixesGetSummariesOK creates a SuffixesGetSummariesOK with default headers values
func NewSuffixesGetSummariesOK() *SuffixesGetSummariesOK {
	return &SuffixesGetSummariesOK{}
}

/*
SuffixesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type SuffixesGetSummariesOK struct {
	Payload []*models.SuffixSummary
}

// IsSuccess returns true when this suffixes get summaries o k response has a 2xx status code
func (o *SuffixesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this suffixes get summaries o k response has a 3xx status code
func (o *SuffixesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this suffixes get summaries o k response has a 4xx status code
func (o *SuffixesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this suffixes get summaries o k response has a 5xx status code
func (o *SuffixesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this suffixes get summaries o k response a status code equal to that given
func (o *SuffixesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the suffixes get summaries o k response
func (o *SuffixesGetSummariesOK) Code() int {
	return 200
}

func (o *SuffixesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Suffixes/Summary][%d] suffixesGetSummariesOK %s", 200, payload)
}

func (o *SuffixesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Suffixes/Summary][%d] suffixesGetSummariesOK %s", 200, payload)
}

func (o *SuffixesGetSummariesOK) GetPayload() []*models.SuffixSummary {
	return o.Payload
}

func (o *SuffixesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSuffixesGetSummariesDefault creates a SuffixesGetSummariesDefault with default headers values
func NewSuffixesGetSummariesDefault(code int) *SuffixesGetSummariesDefault {
	return &SuffixesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
SuffixesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type SuffixesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this suffixes get summaries default response has a 2xx status code
func (o *SuffixesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this suffixes get summaries default response has a 3xx status code
func (o *SuffixesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this suffixes get summaries default response has a 4xx status code
func (o *SuffixesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this suffixes get summaries default response has a 5xx status code
func (o *SuffixesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this suffixes get summaries default response a status code equal to that given
func (o *SuffixesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the suffixes get summaries default response
func (o *SuffixesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *SuffixesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Suffixes/Summary][%d] Suffixes_GetSummaries default %s", o._statusCode, payload)
}

func (o *SuffixesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Suffixes/Summary][%d] Suffixes_GetSummaries default %s", o._statusCode, payload)
}

func (o *SuffixesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SuffixesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
