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

// ConstituentDocumentsGetAllSummaryReader is a Reader for the ConstituentDocumentsGetAllSummary structure.
type ConstituentDocumentsGetAllSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ConstituentDocumentsGetAllSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewConstituentDocumentsGetAllSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Documents/Summary] ConstituentDocuments_GetAllSummary", response, response.Code())
	}
}

// NewConstituentDocumentsGetAllSummaryOK creates a ConstituentDocumentsGetAllSummaryOK with default headers values
func NewConstituentDocumentsGetAllSummaryOK() *ConstituentDocumentsGetAllSummaryOK {
	return &ConstituentDocumentsGetAllSummaryOK{}
}

/*
ConstituentDocumentsGetAllSummaryOK describes a response with status code 200, with default header values.

OK
*/
type ConstituentDocumentsGetAllSummaryOK struct {
	Payload []*models.DocumentSummary
}

// IsSuccess returns true when this constituent documents get all summary o k response has a 2xx status code
func (o *ConstituentDocumentsGetAllSummaryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this constituent documents get all summary o k response has a 3xx status code
func (o *ConstituentDocumentsGetAllSummaryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this constituent documents get all summary o k response has a 4xx status code
func (o *ConstituentDocumentsGetAllSummaryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this constituent documents get all summary o k response has a 5xx status code
func (o *ConstituentDocumentsGetAllSummaryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this constituent documents get all summary o k response a status code equal to that given
func (o *ConstituentDocumentsGetAllSummaryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the constituent documents get all summary o k response
func (o *ConstituentDocumentsGetAllSummaryOK) Code() int {
	return 200
}

func (o *ConstituentDocumentsGetAllSummaryOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Documents/Summary][%d] constituentDocumentsGetAllSummaryOK  %+v", 200, o.Payload)
}

func (o *ConstituentDocumentsGetAllSummaryOK) String() string {
	return fmt.Sprintf("[GET /CRM/Documents/Summary][%d] constituentDocumentsGetAllSummaryOK  %+v", 200, o.Payload)
}

func (o *ConstituentDocumentsGetAllSummaryOK) GetPayload() []*models.DocumentSummary {
	return o.Payload
}

func (o *ConstituentDocumentsGetAllSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
