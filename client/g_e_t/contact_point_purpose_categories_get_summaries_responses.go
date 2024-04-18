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

// ContactPointPurposeCategoriesGetSummariesReader is a Reader for the ContactPointPurposeCategoriesGetSummaries structure.
type ContactPointPurposeCategoriesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactPointPurposeCategoriesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContactPointPurposeCategoriesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/ContactPointPurposeCategories/Summary] ContactPointPurposeCategories_GetSummaries", response, response.Code())
	}
}

// NewContactPointPurposeCategoriesGetSummariesOK creates a ContactPointPurposeCategoriesGetSummariesOK with default headers values
func NewContactPointPurposeCategoriesGetSummariesOK() *ContactPointPurposeCategoriesGetSummariesOK {
	return &ContactPointPurposeCategoriesGetSummariesOK{}
}

/*
ContactPointPurposeCategoriesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type ContactPointPurposeCategoriesGetSummariesOK struct {
	Payload []*models.PurposeCategorySummary
}

// IsSuccess returns true when this contact point purpose categories get summaries o k response has a 2xx status code
func (o *ContactPointPurposeCategoriesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contact point purpose categories get summaries o k response has a 3xx status code
func (o *ContactPointPurposeCategoriesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contact point purpose categories get summaries o k response has a 4xx status code
func (o *ContactPointPurposeCategoriesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contact point purpose categories get summaries o k response has a 5xx status code
func (o *ContactPointPurposeCategoriesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contact point purpose categories get summaries o k response a status code equal to that given
func (o *ContactPointPurposeCategoriesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contact point purpose categories get summaries o k response
func (o *ContactPointPurposeCategoriesGetSummariesOK) Code() int {
	return 200
}

func (o *ContactPointPurposeCategoriesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPointPurposeCategories/Summary][%d] contactPointPurposeCategoriesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ContactPointPurposeCategoriesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/ContactPointPurposeCategories/Summary][%d] contactPointPurposeCategoriesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *ContactPointPurposeCategoriesGetSummariesOK) GetPayload() []*models.PurposeCategorySummary {
	return o.Payload
}

func (o *ContactPointPurposeCategoriesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}