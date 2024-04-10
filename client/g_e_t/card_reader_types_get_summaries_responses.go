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

// CardReaderTypesGetSummariesReader is a Reader for the CardReaderTypesGetSummaries structure.
type CardReaderTypesGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CardReaderTypesGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCardReaderTypesGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /ReferenceData/CardReaderTypes/Summary] CardReaderTypes_GetSummaries", response, response.Code())
	}
}

// NewCardReaderTypesGetSummariesOK creates a CardReaderTypesGetSummariesOK with default headers values
func NewCardReaderTypesGetSummariesOK() *CardReaderTypesGetSummariesOK {
	return &CardReaderTypesGetSummariesOK{}
}

/*
CardReaderTypesGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type CardReaderTypesGetSummariesOK struct {
	Payload []*models.CardReaderTypeSummary
}

// IsSuccess returns true when this card reader types get summaries o k response has a 2xx status code
func (o *CardReaderTypesGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this card reader types get summaries o k response has a 3xx status code
func (o *CardReaderTypesGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this card reader types get summaries o k response has a 4xx status code
func (o *CardReaderTypesGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this card reader types get summaries o k response has a 5xx status code
func (o *CardReaderTypesGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this card reader types get summaries o k response a status code equal to that given
func (o *CardReaderTypesGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the card reader types get summaries o k response
func (o *CardReaderTypesGetSummariesOK) Code() int {
	return 200
}

func (o *CardReaderTypesGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /ReferenceData/CardReaderTypes/Summary][%d] cardReaderTypesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *CardReaderTypesGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /ReferenceData/CardReaderTypes/Summary][%d] cardReaderTypesGetSummariesOK  %+v", 200, o.Payload)
}

func (o *CardReaderTypesGetSummariesOK) GetPayload() []*models.CardReaderTypeSummary {
	return o.Payload
}

func (o *CardReaderTypesGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
