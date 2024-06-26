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

// SeasonsGetSummariesReader is a Reader for the SeasonsGetSummaries structure.
type SeasonsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SeasonsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSeasonsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSeasonsGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSeasonsGetSummariesOK creates a SeasonsGetSummariesOK with default headers values
func NewSeasonsGetSummariesOK() *SeasonsGetSummariesOK {
	return &SeasonsGetSummariesOK{}
}

/*
SeasonsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type SeasonsGetSummariesOK struct {
	Payload []*models.SeasonSummary
}

// IsSuccess returns true when this seasons get summaries o k response has a 2xx status code
func (o *SeasonsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this seasons get summaries o k response has a 3xx status code
func (o *SeasonsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this seasons get summaries o k response has a 4xx status code
func (o *SeasonsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this seasons get summaries o k response has a 5xx status code
func (o *SeasonsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this seasons get summaries o k response a status code equal to that given
func (o *SeasonsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the seasons get summaries o k response
func (o *SeasonsGetSummariesOK) Code() int {
	return 200
}

func (o *SeasonsGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Seasons/Summary][%d] seasonsGetSummariesOK %s", 200, payload)
}

func (o *SeasonsGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Seasons/Summary][%d] seasonsGetSummariesOK %s", 200, payload)
}

func (o *SeasonsGetSummariesOK) GetPayload() []*models.SeasonSummary {
	return o.Payload
}

func (o *SeasonsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSeasonsGetSummariesDefault creates a SeasonsGetSummariesDefault with default headers values
func NewSeasonsGetSummariesDefault(code int) *SeasonsGetSummariesDefault {
	return &SeasonsGetSummariesDefault{
		_statusCode: code,
	}
}

/*
SeasonsGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type SeasonsGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this seasons get summaries default response has a 2xx status code
func (o *SeasonsGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this seasons get summaries default response has a 3xx status code
func (o *SeasonsGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this seasons get summaries default response has a 4xx status code
func (o *SeasonsGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this seasons get summaries default response has a 5xx status code
func (o *SeasonsGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this seasons get summaries default response a status code equal to that given
func (o *SeasonsGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the seasons get summaries default response
func (o *SeasonsGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *SeasonsGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Seasons/Summary][%d] Seasons_GetSummaries default %s", o._statusCode, payload)
}

func (o *SeasonsGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/Seasons/Summary][%d] Seasons_GetSummaries default %s", o._statusCode, payload)
}

func (o *SeasonsGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SeasonsGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
