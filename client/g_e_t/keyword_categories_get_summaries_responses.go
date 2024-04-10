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

// KeywordCategoriesGetSummariesReader is a Reader for the KeywordCategoriesGetSummaries structure.
type KeywordCategoriesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *KeywordCategoriesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewKeywordCategoriesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/KeywordCategories/Summary] KeywordCategories_GetSummaries", response, response.Code())
	}
}

// NewKeywordCategoriesGetSummariesOK creates a KeywordCategoriesGetSummariesOK with default headers values
func NewKeywordCategoriesGetSummariesOK() *KeywordCategoriesGetSummariesOK {
	return &KeywordCategoriesGetSummariesOK{}
}

/*
KeywordCategoriesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type KeywordCategoriesGetSummariesOK struct {
	Payload []*models.KeywordCategorySummary
}

// IsSuccess returns true when this keyword categories get summaries o k response has a 2xx status code
func (o *KeywordCategoriesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this keyword categories get summaries o k response has a 3xx status code
func (o *KeywordCategoriesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this keyword categories get summaries o k response has a 4xx status code
func (o *KeywordCategoriesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this keyword categories get summaries o k response has a 5xx status code
func (o *KeywordCategoriesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this keyword categories get summaries o k response a status code equal to that given
func (o *KeywordCategoriesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the keyword categories get summaries o k response
func (o *KeywordCategoriesGetSummariesOK) Code() int {
	return 200
}

func (o *KeywordCategoriesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/KeywordCategories/Summary][%d] keywordCategoriesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *KeywordCategoriesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/KeywordCategories/Summary][%d] keywordCategoriesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *KeywordCategoriesGetSummariesOK) GetPayload() []*models.KeywordCategorySummary {
	return o.Payload
}

func (o *KeywordCategoriesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
