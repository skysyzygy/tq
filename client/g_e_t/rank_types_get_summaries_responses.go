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

// RankTypesGetSummariesReader is a Reader for the RankTypesGetSummaries structure.
type RankTypesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RankTypesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRankTypesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/RankTypes/Summary] RankTypes_GetSummaries", response, response.Code())
	}
}

// NewRankTypesGetSummariesOK creates a RankTypesGetSummariesOK with default headers values
func NewRankTypesGetSummariesOK() *RankTypesGetSummariesOK {
	return &RankTypesGetSummariesOK{}
}

/*
RankTypesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type RankTypesGetSummariesOK struct {
	Payload []*models.RankTypeSummary
}

// IsSuccess returns true when this rank types get summaries o k response has a 2xx status code
func (o *RankTypesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rank types get summaries o k response has a 3xx status code
func (o *RankTypesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rank types get summaries o k response has a 4xx status code
func (o *RankTypesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rank types get summaries o k response has a 5xx status code
func (o *RankTypesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rank types get summaries o k response a status code equal to that given
func (o *RankTypesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rank types get summaries o k response
func (o *RankTypesGetSummariesOK) Code() int {
	return 200
}

func (o *RankTypesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/RankTypes/Summary][%d] rankTypesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *RankTypesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/RankTypes/Summary][%d] rankTypesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *RankTypesGetSummariesOK) GetPayload() []*models.RankTypeSummary {
	return o.Payload
}

func (o *RankTypesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}