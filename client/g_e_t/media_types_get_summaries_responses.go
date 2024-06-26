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

// MediaTypesGetSummariesReader is a Reader for the MediaTypesGetSummaries structure.
type MediaTypesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MediaTypesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMediaTypesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMediaTypesGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMediaTypesGetSummariesOK creates a MediaTypesGetSummariesOK with default headers values
func NewMediaTypesGetSummariesOK() *MediaTypesGetSummariesOK {
	return &MediaTypesGetSummariesOK{}
}

/*
MediaTypesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type MediaTypesGetSummariesOK struct {
	Payload []*models.MediaTypeSummary
}

// IsSuccess returns true when this media types get summaries o k response has a 2xx status code
func (o *MediaTypesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this media types get summaries o k response has a 3xx status code
func (o *MediaTypesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this media types get summaries o k response has a 4xx status code
func (o *MediaTypesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this media types get summaries o k response has a 5xx status code
func (o *MediaTypesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this media types get summaries o k response a status code equal to that given
func (o *MediaTypesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the media types get summaries o k response
func (o *MediaTypesGetSummariesOK) Code() int {
	return 200
}

func (o *MediaTypesGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/MediaTypes/Summary][%d] mediaTypesGetSummariesOK %s", 200, payload)
}

func (o *MediaTypesGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/MediaTypes/Summary][%d] mediaTypesGetSummariesOK %s", 200, payload)
}

func (o *MediaTypesGetSummariesOK) GetPayload() []*models.MediaTypeSummary {
	return o.Payload
}

func (o *MediaTypesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMediaTypesGetSummariesDefault creates a MediaTypesGetSummariesDefault with default headers values
func NewMediaTypesGetSummariesDefault(code int) *MediaTypesGetSummariesDefault {
	return &MediaTypesGetSummariesDefault{
		_statusCode: code,
	}
}

/*
MediaTypesGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type MediaTypesGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this media types get summaries default response has a 2xx status code
func (o *MediaTypesGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this media types get summaries default response has a 3xx status code
func (o *MediaTypesGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this media types get summaries default response has a 4xx status code
func (o *MediaTypesGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this media types get summaries default response has a 5xx status code
func (o *MediaTypesGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this media types get summaries default response a status code equal to that given
func (o *MediaTypesGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the media types get summaries default response
func (o *MediaTypesGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *MediaTypesGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/MediaTypes/Summary][%d] MediaTypes_GetSummaries default %s", o._statusCode, payload)
}

func (o *MediaTypesGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/MediaTypes/Summary][%d] MediaTypes_GetSummaries default %s", o._statusCode, payload)
}

func (o *MediaTypesGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *MediaTypesGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
