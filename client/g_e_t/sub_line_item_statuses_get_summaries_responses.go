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

// SubLineItemStatusesGetSummariesReader is a Reader for the SubLineItemStatusesGetSummaries structure.
type SubLineItemStatusesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SubLineItemStatusesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSubLineItemStatusesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/SubLineItemStatuses/Summary] SubLineItemStatuses_GetSummaries", response, response.Code())
	}
}

// NewSubLineItemStatusesGetSummariesOK creates a SubLineItemStatusesGetSummariesOK with default headers values
func NewSubLineItemStatusesGetSummariesOK() *SubLineItemStatusesGetSummariesOK {
	return &SubLineItemStatusesGetSummariesOK{}
}

/*
SubLineItemStatusesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type SubLineItemStatusesGetSummariesOK struct {
	Payload []*models.SubLineItemStatusSummary
}

// IsSuccess returns true when this sub line item statuses get summaries o k response has a 2xx status code
func (o *SubLineItemStatusesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this sub line item statuses get summaries o k response has a 3xx status code
func (o *SubLineItemStatusesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this sub line item statuses get summaries o k response has a 4xx status code
func (o *SubLineItemStatusesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this sub line item statuses get summaries o k response has a 5xx status code
func (o *SubLineItemStatusesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this sub line item statuses get summaries o k response a status code equal to that given
func (o *SubLineItemStatusesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the sub line item statuses get summaries o k response
func (o *SubLineItemStatusesGetSummariesOK) Code() int {
	return 200
}

func (o *SubLineItemStatusesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/SubLineItemStatuses/Summary][%d] subLineItemStatusesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *SubLineItemStatusesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/SubLineItemStatuses/Summary][%d] subLineItemStatusesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *SubLineItemStatusesGetSummariesOK) GetPayload() []*models.SubLineItemStatusSummary {
	return o.Payload
}

func (o *SubLineItemStatusesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
