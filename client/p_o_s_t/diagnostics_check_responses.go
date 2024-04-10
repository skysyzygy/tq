// Code generated by go-swagger; DO NOT EDIT.

package p_o_s_t

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/skysyzygy/tq/models"
)

// DiagnosticsCheckReader is a Reader for the DiagnosticsCheck structure.
type DiagnosticsCheckReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DiagnosticsCheckReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDiagnosticsCheckOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /Diagnostics/Check] Diagnostics_Check", response, response.Code())
	}
}

// NewDiagnosticsCheckOK creates a DiagnosticsCheckOK with default headers values
func NewDiagnosticsCheckOK() *DiagnosticsCheckOK {
	return &DiagnosticsCheckOK{}
}

/*
DiagnosticsCheckOK describes a response with status code 200, with default header values.

OK
*/
type DiagnosticsCheckOK struct {
	Payload *models.DatabaseCheckResponse
}

// IsSuccess returns true when this diagnostics check o k response has a 2xx status code
func (o *DiagnosticsCheckOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this diagnostics check o k response has a 3xx status code
func (o *DiagnosticsCheckOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this diagnostics check o k response has a 4xx status code
func (o *DiagnosticsCheckOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this diagnostics check o k response has a 5xx status code
func (o *DiagnosticsCheckOK) IsServerError() bool {
	return false
}

// IsCode returns true when this diagnostics check o k response a status code equal to that given
func (o *DiagnosticsCheckOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the diagnostics check o k response
func (o *DiagnosticsCheckOK) Code() int {
	return 200
}

func (o *DiagnosticsCheckOK) Error() string {
	return fmt.Sprintf("[POST /Diagnostics/Check][%d] diagnosticsCheckOK  %+v", 200, o.Payload)
}

func (o *DiagnosticsCheckOK) String() string {
	return fmt.Sprintf("[POST /Diagnostics/Check][%d] diagnosticsCheckOK  %+v", 200, o.Payload)
}

func (o *DiagnosticsCheckOK) GetPayload() *models.DatabaseCheckResponse {
	return o.Payload
}

func (o *DiagnosticsCheckOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DatabaseCheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
