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

// ContactPointCategoryPurposesGetSummariesReader is a Reader for the ContactPointCategoryPurposesGetSummaries structure.
type ContactPointCategoryPurposesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPointCategoryPurposesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPointCategoryPurposesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ContactPointCategoryPurposes/Summary] ContactPointCategoryPurposes_GetSummaries", response, response.Code())
	}
}

// NewContactPointCategoryPurposesGetSummariesOK creates a ContactPointCategoryPurposesGetSummariesOK with default headers values
func NewContactPointCategoryPurposesGetSummariesOK() *ContactPointCategoryPurposesGetSummariesOK {
	return &ContactPointCategoryPurposesGetSummariesOK{}
}

/*
ContactPointCategoryPurposesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type ContactPointCategoryPurposesGetSummariesOK struct {
	Payload []*models.ContactPointCategoryPurposeSummary
}

// IsSuccess returns true when this contact point category purposes get summaries o k response has a 2xx status code
func (o *ContactPointCategoryPurposesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact point category purposes get summaries o k response has a 3xx status code
func (o *ContactPointCategoryPurposesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact point category purposes get summaries o k response has a 4xx status code
func (o *ContactPointCategoryPurposesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact point category purposes get summaries o k response has a 5xx status code
func (o *ContactPointCategoryPurposesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact point category purposes get summaries o k response a status code equal to that given
func (o *ContactPointCategoryPurposesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact point category purposes get summaries o k response
func (o *ContactPointCategoryPurposesGetSummariesOK) Code() int {
	return 200
}

func (o *ContactPointCategoryPurposesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPointCategoryPurposes/Summary][%d] contactPointCategoryPurposesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ContactPointCategoryPurposesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPointCategoryPurposes/Summary][%d] contactPointCategoryPurposesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ContactPointCategoryPurposesGetSummariesOK) GetPayload() []*models.ContactPointCategoryPurposeSummary {
	return o.Payload
}

func (o *ContactPointCategoryPurposesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
