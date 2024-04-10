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

// ResearchGetReader is a Reader for the ResearchGet structure.
type ResearchGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ResearchGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewResearchGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /CRM/Research/{researchEntryId}] Research_Get", response, response.Code())
	}
}

// NewResearchGetOK creates a ResearchGetOK with default headers values
func NewResearchGetOK() *ResearchGetOK {
	return &ResearchGetOK{}
}

/*
ResearchGetOK describes a response with status code 200, with default header values.

OK
*/
type ResearchGetOK struct {
	Payload *models.ResearchEntry
}

// IsSuccess returns true when this research get o k response has a 2xx status code
func (o *ResearchGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this research get o k response has a 3xx status code
func (o *ResearchGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this research get o k response has a 4xx status code
func (o *ResearchGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this research get o k response has a 5xx status code
func (o *ResearchGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this research get o k response a status code equal to that given
func (o *ResearchGetOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the research get o k response
func (o *ResearchGetOK) Code() int {
	return 200
}

func (o *ResearchGetOK) Error() string {
	return fmt.Sprintf("[GET /CRM/Research/{researchEntryId}][%d] researchGetOK  %+v", 200, o.Payload)
}

func (o *ResearchGetOK) String() string {
	return fmt.Sprintf("[GET /CRM/Research/{researchEntryId}][%d] researchGetOK  %+v", 200, o.Payload)
}

func (o *ResearchGetOK) GetPayload() *models.ResearchEntry {
	return o.Payload
}

func (o *ResearchGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResearchEntry)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
