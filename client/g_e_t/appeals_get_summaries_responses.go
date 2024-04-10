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

// AppealsGetSummariesReader is a Reader for the AppealsGetSummaries structure.
type AppealsGetSummariesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AppealsGetSummariesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAppealsGetSummariesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /Finance/Appeals/Summary] Appeals_GetSummaries", response, response.Code())
	}
}

// NewAppealsGetSummariesOK creates a AppealsGetSummariesOK with default headers values
func NewAppealsGetSummariesOK() *AppealsGetSummariesOK {
	return &AppealsGetSummariesOK{}
}

/*
AppealsGetSummariesOK describes a response with status code 200, with default header values.

OK
*/
type AppealsGetSummariesOK struct {
	Payload []*models.AppealSummary
}

// IsSuccess returns true when this appeals get summaries o k response has a 2xx status code
func (o *AppealsGetSummariesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this appeals get summaries o k response has a 3xx status code
func (o *AppealsGetSummariesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this appeals get summaries o k response has a 4xx status code
func (o *AppealsGetSummariesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this appeals get summaries o k response has a 5xx status code
func (o *AppealsGetSummariesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this appeals get summaries o k response a status code equal to that given
func (o *AppealsGetSummariesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the appeals get summaries o k response
func (o *AppealsGetSummariesOK) Code() int {
	return 200
}

func (o *AppealsGetSummariesOK) Error() string {
	return fmt.Sprintf("[GET /Finance/Appeals/Summary][%d] appealsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *AppealsGetSummariesOK) String() string {
	return fmt.Sprintf("[GET /Finance/Appeals/Summary][%d] appealsGetSummariesOK  %+v", 200, o.Payload)
}

func (o *AppealsGetSummariesOK) GetPayload() []*models.AppealSummary {
	return o.Payload
}

func (o *AppealsGetSummariesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
