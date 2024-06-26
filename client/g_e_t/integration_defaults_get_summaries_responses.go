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

// IntegrationDefaultsGetSummariesReader is a Reader for the IntegrationDefaultsGetSummaries structure.
type IntegrationDefaultsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IntegrationDefaultsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIntegrationDefaultsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIntegrationDefaultsGetSummariesDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIntegrationDefaultsGetSummariesOK creates a IntegrationDefaultsGetSummariesOK with default headers values
func NewIntegrationDefaultsGetSummariesOK() *IntegrationDefaultsGetSummariesOK {
	return &IntegrationDefaultsGetSummariesOK{}
}

/*
IntegrationDefaultsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type IntegrationDefaultsGetSummariesOK struct {
	Payload []*models.IntegrationDefaultSummary
}

// IsSuccess returns true when this integration defaults get summaries o k response has a 2xx status code
func (o *IntegrationDefaultsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this integration defaults get summaries o k response has a 3xx status code
func (o *IntegrationDefaultsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this integration defaults get summaries o k response has a 4xx status code
func (o *IntegrationDefaultsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this integration defaults get summaries o k response has a 5xx status code
func (o *IntegrationDefaultsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this integration defaults get summaries o k response a status code equal to that given
func (o *IntegrationDefaultsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the integration defaults get summaries o k response
func (o *IntegrationDefaultsGetSummariesOK) Code() int {
	return 200
}

func (o *IntegrationDefaultsGetSummariesOK) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/IntegrationDefaults/Summary][%d] integrationDefaultsGetSummariesOK %s", 200, payload)
}

func (o *IntegrationDefaultsGetSummariesOK) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/IntegrationDefaults/Summary][%d] integrationDefaultsGetSummariesOK %s", 200, payload)
}

func (o *IntegrationDefaultsGetSummariesOK) GetPayload() []*models.IntegrationDefaultSummary {
	return o.Payload
}

func (o *IntegrationDefaultsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIntegrationDefaultsGetSummariesDefault creates a IntegrationDefaultsGetSummariesDefault with default headers values
func NewIntegrationDefaultsGetSummariesDefault(code int) *IntegrationDefaultsGetSummariesDefault {
	return &IntegrationDefaultsGetSummariesDefault{
		_statusCode: code,
	}
}

/*
IntegrationDefaultsGetSummariesDefault describes a response with status code -1, with default header values.

Error
*/
type IntegrationDefaultsGetSummariesDefault struct {
	_statusCode int

	Payload *models.ErrorMessage
}

// IsSuccess returns true when this integration defaults get summaries default response has a 2xx status code
func (o *IntegrationDefaultsGetSummariesDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this integration defaults get summaries default response has a 3xx status code
func (o *IntegrationDefaultsGetSummariesDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this integration defaults get summaries default response has a 4xx status code
func (o *IntegrationDefaultsGetSummariesDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this integration defaults get summaries default response has a 5xx status code
func (o *IntegrationDefaultsGetSummariesDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this integration defaults get summaries default response a status code equal to that given
func (o *IntegrationDefaultsGetSummariesDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the integration defaults get summaries default response
func (o *IntegrationDefaultsGetSummariesDefault) Code() int {
	return o._statusCode
}

func (o *IntegrationDefaultsGetSummariesDefault) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/IntegrationDefaults/Summary][%d] IntegrationDefaults_GetSummaries default %s", o._statusCode, payload)
}

func (o *IntegrationDefaultsGetSummariesDefault) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[GET /ReferenceData/IntegrationDefaults/Summary][%d] IntegrationDefaults_GetSummaries default %s", o._statusCode, payload)
}

func (o *IntegrationDefaultsGetSummariesDefault) GetPayload() *models.ErrorMessage {
	return o.Payload
}

func (o *IntegrationDefaultsGetSummariesDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorMessage)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
