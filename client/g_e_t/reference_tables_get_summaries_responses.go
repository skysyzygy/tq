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

// ReferenceTablesGetSummariesReader is a Reader for the ReferenceTablesGetSummaries structure.
type ReferenceTablesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReferenceTablesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReferenceTablesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ReferenceTables/Summary] ReferenceTables_GetSummaries", response, response.Code())
	}
}

// NewReferenceTablesGetSummariesOK creates a ReferenceTablesGetSummariesOK with default headers values
func NewReferenceTablesGetSummariesOK() *ReferenceTablesGetSummariesOK {
	return &ReferenceTablesGetSummariesOK{}
}

/*
ReferenceTablesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type ReferenceTablesGetSummariesOK struct {
	Payload []*models.ReferenceTableSummary
}

// IsSuccess returns true when this reference tables get summaries o k response has a 2xx status code
func (o *ReferenceTablesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this reference tables get summaries o k response has a 3xx status code
func (o *ReferenceTablesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this reference tables get summaries o k response has a 4xx status code
func (o *ReferenceTablesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this reference tables get summaries o k response has a 5xx status code
func (o *ReferenceTablesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this reference tables get summaries o k response a status code equal to that given
func (o *ReferenceTablesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the reference tables get summaries o k response
func (o *ReferenceTablesGetSummariesOK) Code() int {
	return 200
}

func (o *ReferenceTablesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ReferenceTables/Summary][%d] referenceTablesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ReferenceTablesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ReferenceTables/Summary][%d] referenceTablesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ReferenceTablesGetSummariesOK) GetPayload() []*models.ReferenceTableSummary {
	return o.Payload
}

func (o *ReferenceTablesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}