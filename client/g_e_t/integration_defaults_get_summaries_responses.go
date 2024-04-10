// Code generated by go-swagger; DO NOT EDIT.

package g_e_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
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
		return nil, runtime.NewAPIError("[GET /ReferenceData/IntegrationDefaults/Summary] IntegrationDefaults_GetSummaries", response, response.Code())
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
	return fmt.Sprintf("[GET /ReferenceData/IntegrationDefaults/Summary][%d] integrationDefaultsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *IntegrationDefaultsGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/IntegrationDefaults/Summary][%d] integrationDefaultsGetSummariesOK  %+v", 200, o.Payload)
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
