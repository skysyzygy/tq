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

// InactiveReasonsGetSummariesReader is a Reader for the InactiveReasonsGetSummaries structure.
type InactiveReasonsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InactiveReasonsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewInactiveReasonsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/InactiveReasons/Summary] InactiveReasons_GetSummaries", response, response.Code())
	}
}

// NewInactiveReasonsGetSummariesOK creates a InactiveReasonsGetSummariesOK with default headers values
func NewInactiveReasonsGetSummariesOK() *InactiveReasonsGetSummariesOK {
	return &InactiveReasonsGetSummariesOK{}
}

/*
InactiveReasonsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type InactiveReasonsGetSummariesOK struct {
	Payload []*models.InactiveReasonSummary
}

// IsSuccess returns true when this inactive reasons get summaries o k response has a 2xx status code
func (o *InactiveReasonsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this inactive reasons get summaries o k response has a 3xx status code
func (o *InactiveReasonsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this inactive reasons get summaries o k response has a 4xx status code
func (o *InactiveReasonsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this inactive reasons get summaries o k response has a 5xx status code
func (o *InactiveReasonsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this inactive reasons get summaries o k response a status code equal to that given
func (o *InactiveReasonsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the inactive reasons get summaries o k response
func (o *InactiveReasonsGetSummariesOK) Code() int {
	return 200
}

func (o *InactiveReasonsGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/InactiveReasons/Summary][%d] inactiveReasonsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *InactiveReasonsGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/InactiveReasons/Summary][%d] inactiveReasonsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *InactiveReasonsGetSummariesOK) GetPayload() []*models.InactiveReasonSummary {
	return o.Payload
}

func (o *InactiveReasonsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
