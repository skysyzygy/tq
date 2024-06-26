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

// SeasonTypesGetSummariesReader is a Reader for the SeasonTypesGetSummaries structure.
type SeasonTypesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SeasonTypesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSeasonTypesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSeasonTypesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSeasonTypesGetSummariesOK creates a SeasonTypesGetSummariesOK with default headers values
func NewSeasonTypesGetSummariesOK() *SeasonTypesGetSummariesOK {
	return &SeasonTypesGetSummariesOK{}
}

/*
SeasonTypesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type SeasonTypesGetSummariesOK struct {
	Payload []*models.SeasonTypeSummary
}

// IsSuccess returns true when this season types get summaries o k response has a 2xx status code
func (o *SeasonTypesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this season types get summaries o k response has a 3xx status code
func (o *SeasonTypesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this season types get summaries o k response has a 4xx status code
func (o *SeasonTypesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this season types get summaries o k response has a 5xx status code
func (o *SeasonTypesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this season types get summaries o k response a status code equal to that given
func (o *SeasonTypesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the season types get summaries o k response
func (o *SeasonTypesGetSummariesOK) Code() int {
	return 200
}

func (o *SeasonTypesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/SeasonTypes/Summary][%d] seasonTypesGetSummariesOK %s", 200, payload)
}

func (o *SeasonTypesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/SeasonTypes/Summary][%d] seasonTypesGetSummariesOK %s", 200, payload)
}

func (o *SeasonTypesGetSummariesOK) GetPayload() []*models.SeasonTypeSummary {
	return o.Payload
}

func (o *SeasonTypesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSeasonTypesGetSummariesDefault creates a SeasonTypesGetSummariesDefault with default headers values
func NewSeasonTypesGetSummariesDefault(code int) *SeasonTypesGetSummariesDefault {
	return &SeasonTypesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
SeasonTypesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type SeasonTypesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this season types get summaries default response has a 2xx status code
func (o *SeasonTypesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this season types get summaries default response has a 3xx status code
func (o *SeasonTypesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this season types get summaries default response has a 4xx status code
func (o *SeasonTypesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this season types get summaries default response has a 5xx status code
func (o *SeasonTypesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this season types get summaries default response a status code equal to that given
func (o *SeasonTypesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the season types get summaries default response
func (o *SeasonTypesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *SeasonTypesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/SeasonTypes/Summary][%d] SeasonTypes_GetSummaries default %s", o._statusCode, payload)
}

func (o *SeasonTypesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/SeasonTypes/Summary][%d] SeasonTypes_GetSummaries default %s", o._statusCode, payload)
}

func (o *SeasonTypesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *SeasonTypesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
