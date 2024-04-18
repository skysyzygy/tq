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

// PerformancesGetSummariesReader is a Reader for the PerformancesGetSummaries structure.
type PerformancesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformancesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformancesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /TXN/Performances/Summary] Performances_GetSummaries", response, response.Code())
	}
}

// NewPerformancesGetSummariesOK creates a PerformancesGetSummariesOK with default headers values
func NewPerformancesGetSummariesOK() *PerformancesGetSummariesOK {
	return &PerformancesGetSummariesOK{}
}

/*
PerformancesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type PerformancesGetSummariesOK struct {
	Payload []*models.PerformanceSummary
}

// IsSuccess returns true when this performances get summaries o k response has a 2xx status code
func (o *PerformancesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performances get summaries o k response has a 3xx status code
func (o *PerformancesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performances get summaries o k response has a 4xx status code
func (o *PerformancesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performances get summaries o k response has a 5xx status code
func (o *PerformancesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performances get summaries o k response a status code equal to that given
func (o *PerformancesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the performances get summaries o k response
func (o *PerformancesGetSummariesOK) Code() int {
	return 200
}

func (o *PerformancesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /TXN/Performances/Summary][%d] performancesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *PerformancesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /TXN/Performances/Summary][%d] performancesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *PerformancesGetSummariesOK) GetPayload() []*models.PerformanceSummary {
	return o.Payload
}

func (o *PerformancesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}