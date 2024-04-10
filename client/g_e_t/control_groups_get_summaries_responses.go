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

// ControlGroupsGetSummariesReader is a Reader for the ControlGroupsGetSummaries structure.
type ControlGroupsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ControlGroupsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewControlGroupsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ControlGroups/Summary] ControlGroups_GetSummaries", response, response.Code())
	}
}

// NewControlGroupsGetSummariesOK creates a ControlGroupsGetSummariesOK with default headers values
func NewControlGroupsGetSummariesOK() *ControlGroupsGetSummariesOK {
	return &ControlGroupsGetSummariesOK{}
}

/*
ControlGroupsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type ControlGroupsGetSummariesOK struct {
	Payload []*models.ControlGroupSummary
}

// IsSuccess returns true when this control groups get summaries o k response has a 2xx status code
func (o *ControlGroupsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this control groups get summaries o k response has a 3xx status code
func (o *ControlGroupsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this control groups get summaries o k response has a 4xx status code
func (o *ControlGroupsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this control groups get summaries o k response has a 5xx status code
func (o *ControlGroupsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this control groups get summaries o k response a status code equal to that given
func (o *ControlGroupsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the control groups get summaries o k response
func (o *ControlGroupsGetSummariesOK) Code() int {
	return 200
}

func (o *ControlGroupsGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ControlGroups/Summary][%d] controlGroupsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ControlGroupsGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ControlGroups/Summary][%d] controlGroupsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ControlGroupsGetSummariesOK) GetPayload() []*models.ControlGroupSummary {
	return o.Payload
}

func (o *ControlGroupsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
